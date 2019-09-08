package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/GolangWebSocket/key"
)

// Machine struct
type Machine struct {
	key
	name        string
	machineType byte

	// Owner
	company *Company

	// Owning objects
	tasks []Task

	// Relation
	firstTask Task
	lastTask  Task
}

// CreateTask method
func (machine *Machine) CreateTask(duration int) Task {
	taskBase := &taskBase{
		key:           keyConfiguration.NewKey(),
		taskType:      machine.machineType,
		duration:      duration,
		machine:       machine,
		previousTask:  nil,
		nextTask:      nil,
		startDateTime: -1, // Hack, need a method to initialize all functions after instance created
	}

	// Create a specific task that wraps the created base task
	specificTask := machine.newSpecificTask(taskBase)

	// Set first task
	if len(machine.tasks) == 0 {
		machine.firstTask = specificTask
	}

	if machine.lastTask != nil {
		// Set previous task
		specificTask.SetPreviousTask(machine.lastTask)

		// Set previous next task
		machine.lastTask.SetNextTask(specificTask)
	}

	// Set last task
	machine.lastTask = specificTask

	// Add task to this Machine list
	machine.tasks = append(machine.tasks, specificTask)

	// Run declarative functions here
	// Remove the hack above, and call an init() method using Once.Do to initialize/calculate functions value, then remove this call
	specificTask.setStartDateTime()

	// Return the intended specific task
	return specificTask
}

// newSpecificTask creates a new specific task and wrap the created base task in it
//
// Specific tasks are: Rolling, Cutting, Folding, and Packing task
func (machine *Machine) newSpecificTask(base *taskBase) Task {
	var specificTask Task

	switch machine.machineType {
	case Rolling:
		specificTask = &TaskRolling{
			taskBase: base,
		}
	case Cutting:
		specificTask = &TaskCutting{
			taskBase: base,
		}
	case Folding:
		specificTask = &TaskFolding{
			taskBase: base,
		}
	case Packing:
		specificTask = &TaskPacking{
			taskBase: base,
		}
	default:
		panic(errors.New("machine has invalid type:" + string(machine.machineType)).Error())
	}

	return specificTask
}

// Tasks returns all tasks owned by this machine
func (machine *Machine) Tasks() []Task {
	if machine == nil {
		return nil
	}

	return machine.tasks
}

// MachineName returns the name of the machine
func (machine *Machine) MachineName() *string {
	if machine == nil {
		return nil
	}

	return &machine.name
}

// MachineType returns the type of the machine
func (machine *Machine) MachineType() *byte {
	if machine == nil {
		return nil
	}

	return &machine.machineType
}

// FirstTask returns the first task of the machine
func (machine *Machine) FirstTask() Task {
	if machine == nil {
		return nil
	}

	return machine.firstTask
}

// LastTask returns the last task of the machine
func (machine *Machine) LastTask() Task {
	if machine == nil {
		return nil
	}

	return machine.lastTask
}

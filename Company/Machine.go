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
	tasks     []Task
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

	// Set first task
	if len(machine.tasks) == 0 {
		machine.firstTask = taskBase
	}

	if machine.lastTask != nil {
		// Set previous task
		taskBase.previousTask = machine.lastTask

		// Set previous next task
		machine.lastTask.SetNextTask(taskBase)
	}

	// Set last task
	machine.lastTask = taskBase

	// Add task to this Machine list
	machine.tasks = append(machine.tasks, taskBase)

	// Run declarative functions here
	// Remove the hack above, and call an init() method using Once.Do to initialize/calculate functions value, then remove this call
	taskBase.setStartDateTime()

	// Store the real task type needed
	specificTask := machine.newSpecificTask(taskBase)

	return specificTask
}

// Create a new specific task and wrap the created base task in it
//
// Specific tasks are: Rolling, Cutting, Folding, and Packing task
func (machine *Machine) newSpecificTask(base *taskBase) Task {
	var specificTask Task

	switch machine.machineType {
	case 'R':
		specificTask = &TaskRolling{
			taskBase: base,
		}
	case 'C':
		specificTask = &TaskCutting{
			taskBase: base,
		}
	case 'F':
		specificTask = &TaskFolding{
			taskBase: base,
		}
	case 'P':
		specificTask = &TaskPacking{
			taskBase: base,
		}
	default:
		panic(errors.New("Machine has invalid type:" + string(machine.machineType)).Error())
	}

	return specificTask
}

// GetAllTasks return all tasks owned by this machine
func (machine *Machine) GetAllTasks() []Task {
	if machine == nil {
		return nil
	}

	return machine.tasks
}

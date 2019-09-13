package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/QuiLite/key"
)

// Machine struct
type Machine struct {
	key
	name        string
	machineType byte

	// Owner
	company *Company

	// Owning objects
	tasks []*Task

	// Relation
	firstTask *Task
	lastTask  *Task
}

// Set machine type in constant
const (
	rolling = 'R'
	cutting = 'C'
	folding = 'F'
	packing = 'P'
)

// CreateTask method
func (machine *Machine) CreateTask(duration int) *Task {
	task := &Task{
		key:           keyConfiguration.NewKey(),
		taskType:      machine.machineType,
		duration:      duration,
		machine:       machine,
		previousTask:  nil,
		nextTask:      nil,
		startDateTime: -1, // Hack, need a method to initialize all functions after instance created
	}

	// Create a specific task and add it to the new task instance
	machine.newSpecificTask(task)

	// Set first task
	if len(machine.tasks) == 0 {
		machine.firstTask = task
	}

	if machine.lastTask != nil {
		// Set previous task
		task.SetPreviousTask(machine.lastTask)

		// Set previous next task
		machine.lastTask.SetNextTask(task)
	}

	// Set last task
	machine.lastTask = task

	// Add task to this Machine list
	machine.tasks = append(machine.tasks, task)

	// Run declarative functions here
	// Remove the hack above, and call an init() method using Once.Do to initialize/calculate functions value, then remove this call
	task.setStartDateTime()

	// Return the intended specific task
	return task
}

// newSpecificTask creates a new specific task and assign it to the specific task field of the input *Task
//
// Specific tasks are: rolling, cutting, folding, and packing task
func (machine *Machine) newSpecificTask(task *Task) {
	switch machine.machineType {
	case rolling:
		task.specificTask = &taskRolling{
			Task: task,
		}

	case cutting:
		task.specificTask = &taskCutting{
			Task: task,
		}

	case folding:
		task.specificTask = &taskFolding{
			Task: task,
		}

	case packing:
		task.specificTask = &taskPacking{
			Task: task,
		}

	default:
		panic(errors.New("machine has invalid type:" + string(machine.machineType)).Error())
	}
}

// Tasks returns all tasks owned by this machine
func (machine *Machine) Tasks() []*Task {
	if machine == nil {
		return nil
	}

	return machine.tasks
}

// MachineName returns the name of the machine
func (machine *Machine) MachineName() string {
	if machine == nil {
		panic(errors.New("machine is nil").Error())
	}

	return machine.name
}

// MachineType returns the type of the machine
func (machine *Machine) MachineType() byte {
	if machine == nil {
		panic(errors.New("machine is nil").Error())
	}

	return machine.machineType
}

// FirstTask returns the first task of the machine
func (machine *Machine) FirstTask() *Task {
	if machine == nil {
		return nil
	}

	return machine.firstTask
}

// LastTask returns the last task of the machine
func (machine *Machine) LastTask() *Task {
	if machine == nil {
		return nil
	}

	return machine.lastTask
}

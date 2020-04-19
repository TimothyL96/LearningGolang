package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/LearningGolang/key"
)

// Machine struct
type Machine struct {
	key
	name        string
	machineType byte

	// Owner
	company *Company

	// Relation
	// Owning objects
	tasks     []Task
	firstTask Task
	lastTask  Task
}

// Set machine type in constant
const (
	Rolling = 'R'
	Cutting = 'C'
	Folding = 'F'
	Packing = 'P'
)

// CreateTask method
func (machine *Machine) CreateTask(duration int) Task {
	bTask := &BaseTask{
		key:           keyConfiguration.NewKey(),
		taskType:      machine.machineType,
		duration:      duration,
		machine:       machine,
		previousTask:  nil,
		nextTask:      nil,
		startDateTime: -1, // Hack, need a method to initialize all functions after instance created
	}

	// Create a specific task and add it to the new task instance
	task := machine.newTask(bTask)

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

// newTask creates a new specific task.
// Specific tasks are: rolling, cutting, folding, and packing task
func (machine *Machine) newTask(task *BaseTask) Task {
	var sTask Task

	switch machine.machineType {
	case Rolling:
		sTask = &taskRolling{
			task,
		}

	case Cutting:
		sTask = &taskCutting{
			task,
		}

	case Folding:
		sTask = &taskFolding{
			task,
		}

	case Packing:
		sTask = &taskPacking{
			task,
		}

	default:
		panic(errors.New("machine has invalid type:" + string(machine.machineType)).Error())
	}

	return sTask
}

// Name returns the name of the machine
func (machine *Machine) Name() string {
	if machine == nil {
		panic(errors.New("machine is nil").Error())
	}

	return machine.name
}

// Type returns the type of the machine
func (machine *Machine) Type() byte {
	if machine == nil {
		panic(errors.New("machine is nil").Error())
	}

	return machine.machineType
}

// Company returns the machine owner Company
func (machine *Machine) Company() *Company {
	if machine == nil {
		return nil
	}

	return machine.company
}

// Tasks returns all tasks owned by this machine
func (machine *Machine) Tasks() []Task {
	if machine == nil {
		return nil
	}

	return machine.tasks
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

func (machine *Machine) Plan() {

}

func (machine *Machine) UnPlan() {

}

func (machine *Machine) MoveBeforeTask() {

}

func (machine *Machine) MoveAfterTask() {

}

package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/LearningGolang/key"
	. "github.com/ttimt/LearningGolang/stdlib"
)

// Task interface registers all methods of task
type Task interface {
	// Declarative functions
	setStartDateTime()
	setEndDateTime()

	// Set values
	SetPreviousTask(Task)
	SetNextTask(Task)
	SetDuration(duration int)

	// Get values
	Key() *keyConfiguration.BaseKey
	Super() *BaseTask
	EndDateTime() int
	StartDateTime() int
	TaskType() byte
	Duration() int
	PreviousTask() Task
	NextTask() Task
	Machine() *Machine
}

// BaseTask is the base and main struct for task.
// This is used instead of interface to allow nil pointer struct to call methods
type BaseTask struct {
	key
	taskType      byte
	duration      int
	startDateTime int // function
	endDateTime   int // function

	// Super()
	*BaseTask

	// Owner
	machine *Machine

	// Relation
	previousTask Task
	nextTask     Task
}

// taskRolling is the struct for rolling task
type taskRolling struct {
	*BaseTask
}

// taskCutting is the struct for cutting task
type taskCutting struct {
	*BaseTask
}

// taskFolding is the struct for folding task
type taskFolding struct {
	*BaseTask
}

// taskPacking is the struct for packing task
type taskPacking struct {
	*BaseTask
}

// StartDateTime returns the start date time of the task
func (task *BaseTask) StartDateTime() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.startDateTime
}

// EndDateTime returns the end date time of the task
func (task *BaseTask) EndDateTime() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.endDateTime
}

// TaskType returns the task type of the task
func (task *BaseTask) TaskType() byte {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.taskType
}

// Duration returns the duration of the task
func (task *BaseTask) Duration() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.duration
}

// PreviousTask returns the previous task of the task
func (task *BaseTask) PreviousTask() Task {
	if task == nil {
		return nil
	}

	return task.previousTask
}

// NextTask returns the next task of the task
func (task *BaseTask) NextTask() Task {
	if task == nil {
		return nil
	}

	return task.nextTask
}

// Machine returns the owner of the task
func (task *BaseTask) Machine() *Machine {
	if task == nil {
		return nil
	}

	return task.machine
}

// setEndDateTime is a declarative function that gets called when any of its binding is changed.
// Set the task end date time based on the summation of the task start date time and duration
func (task *BaseTask) setEndDateTime() {
	if task == nil {
		return
	}

	value := task.StartDateTime() + task.Duration()

	var funcToPass []func()
	if task.NextTask() != nil {
		funcToPass = append(funcToPass, task.NextTask().setStartDateTime)
	}
	CalcDeclarative(&task.endDateTime, &value, funcToPass...)
}

// SetDuration sets the task duration from the parameter
func (task *BaseTask) SetDuration(duration int) {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	value := duration

	CalcDeclarative(&task.duration, &value, task.setEndDateTime)
}

// SetNextTask set the next task of the task to the parameter input of task
func (task *BaseTask) SetNextTask(newTask Task) {
	task.nextTask = newTask
}

// SetPreviousTask set the previous task of the task to the parameter input of task
func (task *BaseTask) SetPreviousTask(newTask Task) {
	task.previousTask = newTask
}

// Get base task
func (task *BaseTask) Super() *BaseTask {
	return task.BaseTask
}

func (task *BaseTask) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

// setStartDateTime for rolling task
func (task *taskRolling) setStartDateTime() {
	// Modify logic in here as needed for rolling task
	value := task.machine.company.dateTime + 123

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

// setStartDateTime for cutting task
func (task *taskCutting) setStartDateTime() {
	value := task.machine.company.dateTime + 456

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

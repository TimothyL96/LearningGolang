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
	SetDuration(duration int)

	// Set relations
	SetPreviousTask(Task)
	SetNextTask(Task)

	// Key
	Key() *keyConfiguration.BaseKey

	// Get values
	StartDateTime() int
	EndDateTime() int
	TaskType() byte
	Duration() int

	// Get relations
	Super() *BaseTask
	PreviousTask() Task
	NextTask() Task
	Machine() *Machine

	// Conversion
	AsRolling() *taskRolling
	AsCutting() *taskCutting
	AsFolding() *taskFolding
	AsPacking() *taskPacking
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

// Base set start date time
func (task *BaseTask) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

// Conversion Base for Interface
func (task *BaseTask) AsRolling() *taskRolling {
	return nil
}

func (task *BaseTask) AsCutting() *taskCutting {
	return nil
}

func (task *BaseTask) AsFolding() *taskFolding {
	return nil
}

func (task *BaseTask) AsPacking() *taskPacking {
	return nil
}

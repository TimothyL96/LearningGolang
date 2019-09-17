package company

import (
	"errors"

	. "github.com/ttimt/QuiLite/stdlib"
)

// specificTask interface registers all methods of task
type specificTask interface {
	setStartDateTime()
}

// Task is the base and main struct for task.
// This is used instead of interface to allow nil pointer struct to call methods
type Task struct {
	key
	taskType      byte
	duration      int
	startDateTime int // function
	endDateTime   int // function

	// Subclass: Rolling, Cutting, Folding or Packing
	specificTask

	// Owner
	machine *Machine

	// Relation
	previousTask *Task
	nextTask     *Task
}

// taskRolling is the struct for rolling task
type taskRolling struct {
	*Task
}

// taskCutting is the struct for cutting task
type taskCutting struct {
	*Task
}

// taskFolding is the struct for folding task
type taskFolding struct {
	*Task
}

// taskPacking is the struct for packing task
type taskPacking struct {
	*Task
}

// StartDateTime returns the start date time of the task
func (task *Task) StartDateTime() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.startDateTime
}

// EndDateTime returns the end date time of the task
func (task *Task) EndDateTime() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.endDateTime
}

// TaskType returns the task type of the task
func (task *Task) TaskType() byte {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.taskType
}

// Duration returns the duration of the task
func (task *Task) Duration() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.duration
}

// PreviousTask returns the previous task of the task
func (task *Task) PreviousTask() *Task {
	if task == nil {
		return nil
	}

	return task.previousTask
}

// NextTask returns the next task of the task
func (task *Task) NextTask() *Task {
	if task == nil {
		return nil
	}

	return task.nextTask
}

// Machine returns the owner of the task
func (task *Task) Machine() *Machine {
	if task == nil {
		return nil
	}

	return task.machine
}

// setEndDateTime is a declarative function that gets called when any of its binding is changed.
// Set the task end date time based on the summation of the task start date time and duration
func (task *Task) setEndDateTime() {
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
func (task *Task) SetDuration(duration int) {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	value := duration

	CalcDeclarative(&task.duration, &value, task.setEndDateTime)
}

// SetNextTask set the next task of the task to the parameter input of task
func (task *Task) SetNextTask(newTask *Task) {
	task.nextTask = newTask
}

// SetPreviousTask set the previous task of the task to the parameter input of task
func (task *Task) SetPreviousTask(newTask *Task) {
	task.previousTask = newTask
}

// setStartDateTime for main task
//
// Check interface for nil before getting the derived method.
// Also check for recursive call and panic in case derived struct does not implement the method
func (task *Task) setStartDateTime() {
	if task == nil {
		return
	} else if task.specificTask == nil {
		panic(errors.New("fatal error: task does not have specific task").Error())
	}

	// Check for recursive call and panic.
	if isInfinite, err := IsInfiniteRecursiveCall(); isInfinite {
		panic(errors.New(err).Error())
	}

	// Call the overridden method
	task.specificTask.setStartDateTime()
}

// setStartDateTime for rolling task
func (task *taskRolling) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

// setStartDateTime for cutting task
func (task *taskCutting) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

// setStartDateTime for folding task
func (task *taskFolding) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

// setStartDateTime for packing task
func (task *taskPacking) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

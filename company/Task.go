package company

import (
	"errors"
)

// specificTask interface wraps all kinds of tasks for polymorphism
type specificTask interface {
	StartDateTime() int
	startDateTimePtr() *int
	EndDateTime() int
	endDateTimePtr() *int
	TaskType() byte
	Duration() int
	durationPtr() *int
	PreviousTask() *Task
	NextTask() *Task

	setEndDateTime()
	SetDuration(int)
	SetNextTask(*Task)
	SetPreviousTask(*Task)
	setStartDateTime()
}

// Task is the base struct for task
type Task struct {
	key
	taskType      byte
	duration      int
	startDateTime int // function
	endDateTime   int // function

	// Subclass: Rolling, Cutting, Folding or Packing
	specificTask

	// Machine
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

func (task *Task) setSpecificTask(specificTask specificTask) {
	task.specificTask = specificTask
}

// StartDateTime returns the start date time of the task
func (task *Task) StartDateTime() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.startDateTime
}

// startDateTimePtr returns the pointer of start date time of the task
func (task *Task) startDateTimePtr() *int {
	if task == nil {
		return nil
	}

	return &task.startDateTime
}

// EndDateTime returns the end date time of the task
func (task *Task) EndDateTime() int {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	return task.endDateTime
}

// endDateTimePtr returns the pointer of end date time of the task
func (task *Task) endDateTimePtr() *int {
	if task == nil {
		return nil
	}

	return &task.endDateTime
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

// durationPtr returns the pointer of duration of the task
func (task *Task) durationPtr() *int {
	if task == nil {
		return nil
	}

	return &task.duration
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
	CalcDeclarative(task.endDateTimePtr(), &value, funcToPass...)
}

// SetDuration sets the task duration from the parameter
func (task *Task) SetDuration(duration int) {
	if task == nil {
		panic(errors.New("task is nil").Error())
	}

	value := duration

	CalcDeclarative(task.durationPtr(), &value, task.setEndDateTime)
}

// SetNextTask set the next task of the task to the parameter input of task
func (task *Task) SetNextTask(newTask *Task) {
	task.nextTask = newTask
}

// SetPreviousTask set the next task of the task to the parameter input of task
func (task *Task) SetPreviousTask(newTask *Task) {
	task.previousTask = newTask
}

// setStartDateTime for taskBase so that the struct implements specificTask interface
func (task *Task) setStartDateTime() {
	if task == nil || task.specificTask == nil {
		return
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

	CalcDeclarative(task.startDateTimePtr(), &value, task.setEndDateTime)
}

// setStartDateTime for cutting task
func (task *taskCutting) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(task.startDateTimePtr(), &value, task.setEndDateTime)
}

// setStartDateTime for folding task
func (task *taskFolding) setStartDateTime() {
	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(task.startDateTimePtr(), &value, task.setEndDateTime)
}

// setStartDateTime for packing task
func (task *taskPacking) setStartDateTime() {

	value := task.machine.company.dateTime

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(task.startDateTimePtr(), &value, task.setEndDateTime)
}

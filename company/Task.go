package company

// Task interface wraps all kinds of tasks for polymorphism
type Task interface {
	SetDuration(int)
	EndDateTime() *int
	setStartDateTime()
	SetNextTask(newTask Task)
	SetPreviousTask(newTask Task)
}

// taskBase is the base struct for task
type taskBase struct {
	key
	taskType      byte
	duration      int
	startDateTime int // function
	endDateTime   int // function

	// Machine
	machine *Machine

	// Relation
	previousTask Task
	nextTask     Task
}

// TaskRolling is the struct for rolling task
type TaskRolling struct {
	*taskBase
}

// TaskCutting is the struct for cutting task
type TaskCutting struct {
	*taskBase
}

// TaskFolding is the struct for folding task
type TaskFolding struct {
	*taskBase
}

// TaskPacking is the struct for packing task
type TaskPacking struct {
	*taskBase
}

// SetDuration sets the task duration from the parameter
func (task *taskBase) SetDuration(duration int) {
	value := duration

	CalcFunc(&(task.duration), &value, task.setEndDateTime)
}

// setEndDateTime is a declarative function that gets called when any of its binding is changed.
// Set the task end date time based on the summation of the task start date time and duration
func (task *taskBase) setEndDateTime() {
	if task == nil {
		return
	}

	value := task.startDateTime + task.duration

	CalcFunc(&(task.endDateTime), &value, task.nextTask.setStartDateTime)
}

// EndDateTime returns the end date time of the task
func (task *taskBase) EndDateTime() *int {
	if task == nil {
		return nil
	}

	return &task.endDateTime
}

// SetNextTask set the next task of the task to the parameter input of task
func (task *taskBase) SetNextTask(newTask Task) {
	task.nextTask = newTask
}

// SetPreviousTask set the next task of the task to the parameter input of task
func (task *taskBase) SetPreviousTask(newTask Task) {
	task.previousTask = newTask
}

// setStartDateTime for rolling task
func (task *TaskRolling) setStartDateTime() {
	if task == nil {
		return
	}

	value := Guard(task.previousTask.EndDateTime(), task.machine.company.dateTime).(int)

	CalcFunc(&(task.startDateTime), &value, task.setEndDateTime)
}

// setStartDateTime for cutting task
func (task *TaskCutting) setStartDateTime() {
	if task == nil {
		return
	}

	value := Guard(task.previousTask.EndDateTime(), task.machine.company.dateTime).(int)

	CalcFunc(&(task.startDateTime), &value, task.setEndDateTime)
}

// setStartDateTime for folding task
func (task *TaskFolding) setStartDateTime() {
	if task == nil {
		return
	}

	value := Guard(task.previousTask.EndDateTime(), task.machine.company.dateTime).(int)

	CalcFunc(&(task.startDateTime), &value, task.setEndDateTime)
}

// setStartDateTime for packing task
func (task *TaskPacking) setStartDateTime() {
	if task == nil {
		return
	}

	value := Guard(task.previousTask.EndDateTime(), task.machine.company.dateTime).(int)

	CalcFunc(&(task.startDateTime), &value, task.setEndDateTime)
}

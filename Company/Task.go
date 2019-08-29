package company

// Task struct
type Task struct {
	TaskType      byte
	Duration      int
	StartDateTime int
	EndDateTime   int // function

	// Machine
	Machine *Machine

	// Previous task
	PreviousTask *Task
	NextTask     *Task
}

// *** Procedurally set attributes have SET method with parameters, while declarative attributes have no parameters

// SetStartDateTime xaxa
func (task *Task) SetStartDateTime() {
	if task == nil {
		return
	}

	value := Guard(task.PreviousTask.EndDateTime, task.Machine.Company.DateTime).(int)

	CalcFunc(&(task.StartDateTime), value, task.SetEndDateTime, task.Machine.UpdateTasksSorting)
}

// SetDuration xaxa
func (task *Task) SetDuration(duration int) {
	value := duration

	CalcFunc(&(task.Duration), value, task.SetEndDateTime)
}

// SetEndDateTime xaxa
func (task *Task) SetEndDateTime() {
	if task == nil {
		return
	}

	value := task.StartDateTime + task.Duration

	CalcFunc(&(task.EndDateTime), value, task.NextTask.SetStartDateTime)
}

// GetEndDateTime xaxa
func (task *Task) GetEndDateTime() *int {
	if task == nil {
		return nil
	}

	return &task.EndDateTime
}

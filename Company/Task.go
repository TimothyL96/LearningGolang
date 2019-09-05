package company

// Task struct
type Task struct {
	Key           Key
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

// SetStartDateTime xaxa
func (task *Task) SetStartDateTime() {
	if task == nil {
		return
	}

	value := Guard(task.PreviousTask.GetEndDateTime(), task.Machine.Company.DateTime).(int)

	CalcFunc(&(task.StartDateTime), &value, task.SetEndDateTime) //, task.Machine.RelationTaskUpdateSorting)
}

// SetDuration xaxa
func (task *Task) SetDuration(duration int) {
	value := duration

	CalcFunc(&(task.Duration), &value, task.SetEndDateTime)
}

// SetEndDateTime xaxa
func (task *Task) SetEndDateTime() {
	if task == nil {
		return
	}

	value := task.StartDateTime + task.Duration

	CalcFunc(&(task.EndDateTime), &value, task.NextTask.SetStartDateTime)
}

// GetEndDateTime xaxa
func (task *Task) GetEndDateTime() *int {
	if task == nil {
		return nil
	}

	return &task.EndDateTime
}

// GetKey xaxa
func (task *Task) GetKey() *Key {
	if task == nil {
		return nil
	}

	return &task.Key
}

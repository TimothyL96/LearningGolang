package company

// Task interface xaxa
type Task interface {
	SetDuration(int)
	GetEndDateTime() *int
	GetKey() *Key
}

// Task base struct
type taskBase struct {
	key           Key
	taskType      byte
	duration      int
	startDateTime int
	endDateTime   int // function

	// Machine
	Machine *Machine

	// Previous task
	PreviousTask *taskBase
	NextTask     *taskBase
}

type taskCutting struct {
	*taskBase
}

type taskRolling struct {
	*taskBase
}

type taskFolding struct {
	*taskBase
}

type taskPackaging struct {
	*taskBase
}

func (task *taskBase) setStartDateTime() {
	if task == nil {
		return
	}

	value := Guard(task.PreviousTask.GetEndDateTime(), task.Machine.Company.DateTime).(int)

	CalcFunc(&(task.startDateTime), &value, task.setEndDateTime) //, task.Machine.RelationTaskUpdateSorting)
}

// SetDuration xaxa
func (task *taskBase) SetDuration(duration int) {
	value := duration

	CalcFunc(&(task.duration), &value, task.setEndDateTime)
}

func (task *taskBase) setEndDateTime() {
	if task == nil {
		return
	}

	value := task.startDateTime + task.duration

	CalcFunc(&(task.endDateTime), &value, task.NextTask.setStartDateTime)
}

// GetEndDateTime xaxa
func (task *taskBase) GetEndDateTime() *int {
	if task == nil {
		return nil
	}

	return &task.endDateTime
}

// GetKey xaxa
func (task *taskBase) GetKey() *Key {
	if task == nil {
		return nil
	}

	return &task.key
}

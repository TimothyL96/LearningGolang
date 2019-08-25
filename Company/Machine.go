package company

// Machine struct
type Machine struct {
	MachineName string
	MachineType byte

	// Owner
	Owner *Company

	// Owning objects
	Tasks []*Task
}

// CreateTask method
func (machine *Machine) CreateTask(TaskType byte, Duration int, StartDateTime int) *Task {
	task := &Task{
		TaskType:      TaskType,
		Duration:      Duration,
		StartDateTime: StartDateTime,
		Owner:         machine,
	}

	// Add task to this Machine list
	machine.Tasks = append(machine.Tasks, task)

	return task
}

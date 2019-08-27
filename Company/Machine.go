package company

// Machine struct
type Machine struct {
	MachineName string
	MachineType byte

	// Owner
	Company *Company

	// Owning objects
	Tasks     []*Task
	FirstTask *Task
	LastTask  *Task
}

// CreateTask method
func (machine *Machine) CreateTask(TaskType byte, Duration int) *Task {
	task := &Task{
		TaskType: TaskType,
		Duration: Duration,
		Machine:  machine,
	}

	// Run declarative functions here
	task.SetStartDateTime() // omit SetEndDateTime
	task.SetEndDateTime()   // testing

	// Add task to this Machine list
	machine.Tasks = append(machine.Tasks, task)

	// Re-sort the tasks for machine

	return task
}

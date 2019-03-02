package Company

type machine struct {
	MachineName string
	MachineType byte
	
	// Owner
	Owner *Company
	
	// Owning objects
	Tasks[] *task
}

func (machine *machine) CreateTask(TaskType byte) *task {
	task := &task{
		TaskType: TaskType,
		Owner: machine,
	}
	
	// Add task to this machine list
	machine.Tasks = append(machine.Tasks, task)
	
	return task
}
package company

// Task struct
type Task struct {
	TaskType      byte
	Duration      int
	StartDateTime int
	EndDateTime   int

	// Owner
	Owner *Machine
}

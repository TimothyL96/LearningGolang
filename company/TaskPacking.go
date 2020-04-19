package company

// taskPacking is the struct for packing task
type taskPacking struct {
	*BaseTask
}

func (task *taskPacking) AsTaskPacking() *taskPacking {
	return task
}

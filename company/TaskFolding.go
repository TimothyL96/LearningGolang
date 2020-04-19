package company

// taskFolding is the struct for folding task
type taskFolding struct {
	*BaseTask
}

func (task *taskFolding) AsFolding() *taskFolding {
	return task
}

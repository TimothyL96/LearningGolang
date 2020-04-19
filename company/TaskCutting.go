package company

import (
	. "github.com/ttimt/LearningGolang/stdlib"
)

// taskCutting is the struct for cutting task
type taskCutting struct {
	*BaseTask
}

func (task *taskCutting) AsCutting() *taskCutting {
	return task
}

// setStartDateTime for cutting task
func (task *taskCutting) setStartDateTime() {
	value := task.machine.company.dateTime + 456

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

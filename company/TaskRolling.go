package company

import (
	. "github.com/ttimt/LearningGolang/stdlib"
)

// taskRolling is the struct for rolling task
type taskRolling struct {
	*BaseTask
}

func (task *taskRolling) AsRolling() *taskRolling {
	return task
}

// setStartDateTime for rolling task
func (task *taskRolling) setStartDateTime() {
	// Modify logic in here as needed for rolling task
	value := task.machine.company.dateTime + 123

	if task.PreviousTask() != nil {
		value = task.PreviousTask().EndDateTime()
	}

	CalcDeclarative(&task.startDateTime, &value, task.setEndDateTime)
}

// Unique methods
func (task *taskRolling) UniqueToRolling() string {
	return "Im unique to task rolling only!"
}

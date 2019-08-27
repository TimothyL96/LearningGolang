package company

import (
	"sort"
)

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

// UpdateTasksSorting xaxa
func (machine *Machine) UpdateTasksSorting() {
	sort.SliceStable(machine.Tasks, func(i, j int) bool {
		return machine.Tasks[i].StartDateTime < machine.Tasks[j].StartDateTime
	})

	for k, t := range machine.Tasks {
		if k == 0 {
			machine.FirstTask = t
		} else {
			value := machine.Tasks[k-1]
			CalcFuncRelation(t.PreviousTask, value, t.SetStartDateTime)
		}

		if k == len(machine.Tasks)-1 {
			machine.FirstTask = t
			continue
		} else {
			value := machine.Tasks[k+1]
			CalcFuncRelation(t.NextTask, value)
		}
	}
}

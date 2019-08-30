package company

import (
	"sort"
)

// Machine struct
type Machine struct {
	Key         Key
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
		Key:          machine.Company.GetNewKey(),
		TaskType:     TaskType,
		Duration:     Duration,
		Machine:      machine,
		PreviousTask: nil,
		NextTask:     nil,
		StartDateTime: -1, // Hack, need a method to initialize all functions after instance created
	}
	
	if machine.LastTask != nil {
		// Set previous task
		task.PreviousTask = machine.LastTask
		
		// Set previous next task
		machine.LastTask.NextTask = task
	}

	// Set first task
	if len(machine.Tasks) == 0 {
		machine.FirstTask = task
	}

	// Set last task
	machine.LastTask = task

	// Add task to this Machine list
	machine.Tasks = append(machine.Tasks, task)

	// Run declarative functions here
	task.SetStartDateTime() // omit SetEndDateTime

	return task
}

// RelationTaskUpdateSorting xaxa
func (machine *Machine) RelationTaskUpdateSorting() {
	// Sort tasks based on StartDateTime
	sort.SliceStable(machine.Tasks, func(i, j int) bool {
		return machine.Tasks[i].StartDateTime < machine.Tasks[j].StartDateTime
	})

	// Set machine first and last task, and every task's previous and next task
	for k, t := range machine.Tasks {
		if k == 0 {
			machine.FirstTask = t
		} else {
			value := machine.Tasks[k-1]
			if t.PreviousTask == nil {
				t.PreviousTask = &Task{}
			}
			CalcFunc(t.PreviousTask, value, t.SetStartDateTime)
		}

		if k == len(machine.Tasks)-1 {
			machine.LastTask = t
		} else {
			value := machine.Tasks[k+1]
			if t.NextTask == nil {
				t.NextTask = &Task{}
			}
			CalcFunc(t.NextTask, value)
		}
	}
}

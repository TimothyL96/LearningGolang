package company

import (
	"errors"
	"sort"
)

// Machine struct
type Machine struct {
	key         Key
	machineName string
	machineType byte

	// Owner
	Company *Company

	// Owning objects
	Tasks     []*taskBase
	FirstTask *taskBase
	LastTask  *taskBase
}

// CreateTask method
func (machine *Machine) CreateTask(duration int) Task {
	taskBase := &taskBase{
		key:           machine.Company.GetNewKey(),
		taskType:      machine.machineType,
		duration:      duration,
		Machine:       machine,
		PreviousTask:  nil,
		NextTask:      nil,
		startDateTime: -1, // Hack, need a method to initialize all functions after instance created
	}

	// Set first task
	if len(machine.Tasks) == 0 {
		machine.FirstTask = taskBase
	}

	if machine.LastTask != nil {
		// Set previous task
		taskBase.PreviousTask = machine.LastTask

		// Set previous next task
		machine.LastTask.NextTask = taskBase
	}

	// Set last task
	machine.LastTask = taskBase

	// Add task to this Machine list
	machine.Tasks = append(machine.Tasks, taskBase)

	// Run declarative functions here
	taskBase.setStartDateTime() // omit SetEndDateTime

	// Store interface taskBase
	var task Task

	switch machine.machineType {
	case 'R':
		task = &taskRolling{
			taskBase: taskBase,
		}
	case 'C':
		task = &taskCutting{
			taskBase: taskBase,
		}
	case 'F':
		task = &taskFolding{
			taskBase: taskBase,
		}
	case 'P':
		task = &taskPackaging{
			taskBase: taskBase,
		}
	default:
		panic(errors.New("Machine has invalid type:" + string(machine.machineType)).Error())
	}

	return task
}

// RelationTaskUpdateSorting xaxa
func (machine *Machine) RelationTaskUpdateSorting() {
	// Sort tasks based on StartDateTime
	sort.SliceStable(machine.Tasks, func(i, j int) bool {
		return machine.Tasks[i].startDateTime < machine.Tasks[j].startDateTime
	})

	// Set machine first and last task, and every task's previous and next task
	for k, t := range machine.Tasks {
		if k == 0 {
			machine.FirstTask = t
		} else {
			value := machine.Tasks[k-1]
			if t.PreviousTask == nil {
				t.PreviousTask = &taskBase{}
			}
			CalcFunc(t.PreviousTask, value, t.setStartDateTime)
		}

		if k == len(machine.Tasks)-1 {
			machine.LastTask = t
		} else {
			value := machine.Tasks[k+1]
			if t.NextTask == nil {
				t.NextTask = &taskBase{}
			}
			CalcFunc(t.NextTask, value)
		}
	}
}

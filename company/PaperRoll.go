package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/QuiLite/key"
)

// PaperRoll is the struct for paper roll
type PaperRoll struct {
	key
	color  int
	length int // function

	// Owner
	knifeSetting *KnifeSetting

	// Relation
	operations     []*Operation
	firstOperation *Operation
	lastOperation  *Operation
}

// Knife setting returns the owner of the paper roll
func (pr *PaperRoll) KnifeSetting() *KnifeSetting {
	if pr == nil {
		return nil
	}

	return pr.knifeSetting
}

// Length returns the length of the paper roll
func (pr *PaperRoll) Length() int {
	if pr == nil {
		panic(errors.New("paper roll is nil").Error())
	}

	return pr.length
}

// Color returns the color of the paper roll
func (pr *PaperRoll) Color() int {
	if pr == nil {
		panic(errors.New("paper roll is nil").Error())
	}

	return pr.color
}

// Operations returns all the operations of the paper roll
func (pr *PaperRoll) Operations() []*Operation {
	if pr == nil {
		return nil
	}

	return pr.operations
}

// FirstOperation returns the rolling operation of the paper roll
func (pr *PaperRoll) FirstOperation() *Operation {
	if pr == nil {
		return nil
	}

	return pr.firstOperation
}

// LastOperation returns the cutting operation of the paper roll
func (pr *PaperRoll) LastOperation() *Operation {
	if pr == nil {
		return nil
	}

	return pr.lastOperation
}

// createOperation create rolling and cutting operations for this paper roll
func (pr *PaperRoll) createOperation() []*Operation {
	operationRolling := pr.createOperationRolling()
	operationCutting := pr.createOperationCutting()

	operationRolling.setNextOperation(operationCutting)
	operationCutting.setPreviousOperation(operationRolling)

	operations := []*Operation{operationRolling, operationCutting}

	pr.operations = operations
	pr.firstOperation = operationRolling
	pr.lastOperation = operationCutting

	return operations
}

// createOperationRolling create rolling operations for this paper roll
func (pr *PaperRoll) createOperationRolling() *Operation {
	operation := &Operation{
		key:               keyConfiguration.NewKey(),
		isPlanned:         false,
		operationType:     'R',
		specificOperation: nil,
		previousOperation: nil,
		nextOperation:     nil,
		task:              nil,
	}

	operationRolling := &operationRolling{
		operation,
	}

	operation.specificOperation = &operationRollingCutting{
		specificOperation: operationRolling,
		paperRoll:         pr,
	}

	return operation
}

// createOperationCutting create cutting operations for this paper roll
func (pr *PaperRoll) createOperationCutting() *Operation {
	operation := &Operation{
		key:               keyConfiguration.NewKey(),
		isPlanned:         false,
		operationType:     'C',
		specificOperation: nil,
		previousOperation: nil,
		nextOperation:     nil,
		task:              nil,
	}

	operationRolling := &operationCutting{
		operation,
	}

	operation.specificOperation = &operationRollingCutting{
		specificOperation: operationRolling,
		paperRoll:         pr,
	}

	return operation
}

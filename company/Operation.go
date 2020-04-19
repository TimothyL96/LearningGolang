package company

import (
	"errors"
)

// Operation interface is the interface for all operations type.
// This allows polymorphism for 1 operations type to other operations type
type Operation interface {
	// Relations
	PreviousOperation() Operation
	NextOperation() Operation
	Task() Task

	// Get values
	IsPlanned() bool
	OperationType() byte

	// Set
	SetTask(task Task)
	setPreviousOperation(operation Operation)
	setNextOperation(operation Operation)
	setIsPlanned(isPlanned bool)

	// Conversion
	AsOperationRollingCutting() *operationRollingCutting
	AsOperationFoldingPacking() *operationFoldingPacking
	AsOperationRolling() *operationRolling
	AsOperationCutting() *operationCutting
	AsOperationFolding() *operationFolding
	AsOperationPacking() *operationPacking
}

// BaseOperation struct is the struct for operations
type BaseOperation struct {
	key
	isPlanned     bool
	operationType byte

	// Relation
	previousOperation Operation
	nextOperation     Operation

	task Task
}

// IsPlanned gets the isPlanned value of the operations
func (op *BaseOperation) IsPlanned() bool {
	if op == nil {
		panic(errors.New("operation is null").Error())
	}

	return op.isPlanned
}

// OperationType returns the operation type of the operation
func (op *BaseOperation) OperationType() byte {
	if op == nil {
		panic(errors.New("operation is null").Error())
	}

	return op.operationType
}

// PreviousOperation returns the first operations of the operations like rolling or folding
func (op *BaseOperation) PreviousOperation() Operation {
	if op == nil {
		return nil
	}

	return op.previousOperation
}

// NextOperation returns the last operations of the operations like cutting or packing
func (op *BaseOperation) NextOperation() Operation {
	if op == nil {
		return nil
	}

	return op.nextOperation
}

// BaseTask returns the task of the operations
func (op *BaseOperation) Task() Task {
	if op == nil {
		return nil
	}

	return op.task
}

// SetTask sets the task for the current operation
func (op *BaseOperation) SetTask(task Task) {

}

// setPreviousOperation sets the first operations for the operations
func (op *BaseOperation) setPreviousOperation(operation Operation) {
	op.previousOperation = operation
}

// setNextOperation sets the first operations for the operations
func (op *BaseOperation) setNextOperation(operation Operation) {
	op.nextOperation = operation
}

// setIsPlanned accepts a bool parameter and sets it to the operations isPlanned field
func (op *BaseOperation) setIsPlanned(isPlanned bool) {
	op.isPlanned = isPlanned
}

// Base conversion
func (op *BaseOperation) AsOperationRollingCutting() *operationRollingCutting {
	return nil
}

func (op *BaseOperation) AsOperationFoldingPacking() *operationFoldingPacking {
	return nil
}

func (op *BaseOperation) AsOperationRolling() *operationRolling {
	return nil
}

func (op *BaseOperation) AsOperationCutting() *operationCutting {
	return nil
}

func (op *BaseOperation) AsOperationFolding() *operationFolding {
	return nil
}

func (op *BaseOperation) AsOperationPacking() *operationPacking {
	return nil
}

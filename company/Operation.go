package company

import (
	"errors"

	. "github.com/ttimt/QuiLite/stdlib"
)

// specificOperation interface is the interface for all operations type.
// This allows polymorphism for 1 operations type to other operations type
type specificOperation interface {
	PreviousOperation() *Operation
	NextOperation() *Operation
	PaperRoll() *PaperRoll
	Order() *Order
}

// Operation struct is the struct for operations
type Operation struct {
	key
	isPlanned     bool
	operationType byte

	// Subclass
	specificOperation

	// Relation
	previousOperation *Operation
	nextOperation     *Operation

	task *Task
}

// operationRollingCutting is the base struct for rolling and cutting operations
type operationRollingCutting struct {
	specificOperation

	// Owner
	paperRoll *PaperRoll
}

// operationFoldingPacking is the base struct for folding and packing operations
type operationFoldingPacking struct {
	specificOperation

	// Owner
	order *Order
}

// operationRolling is the struct for rolling operations
type operationRolling struct {
	*Operation
}

// operationCutting is the struct for cutting operations
type operationCutting struct {
	*Operation
}

// operationFolding is the struct for folding operations
type operationFolding struct {
	*Operation
}

// operationPacking is the struct for packing operations
type operationPacking struct {
	*Operation
}

// setIsPlanned accepts a bool parameter and sets it to the operations isPlanned field
func (op *Operation) setIsPlanned(isPlanned bool) {
	op.isPlanned = isPlanned
}

// IsPlanned gets the isPlanned value of the operations
func (op *Operation) IsPlanned() bool {
	if op == nil {
		panic(errors.New("operation is null").Error())
	}

	return op.isPlanned
}

// OperationType returns the operation type of the operation
func (op *Operation) OperationType() byte {
	if op == nil {
		panic(errors.New("operation is null").Error())
	}

	return op.operationType
}

// PreviousOperation returns the first operations of the operations like rolling or folding
func (op *Operation) PreviousOperation() *Operation {
	if op == nil {
		return nil
	}

	return op.previousOperation
}

// NextOperation returns the last operations of the operations like cutting or packing
func (op *Operation) NextOperation() *Operation {
	if op == nil {
		return nil
	}

	return op.nextOperation
}

// PaperRoll is the base method and returns the paper roll of the specific operations
func (op *Operation) PaperRoll() *PaperRoll {
	if op == nil {
		return nil
	}

	// Check for recursive call and panic.
	if isInfinite, err := IsInfiniteRecursiveCall(); isInfinite {
		panic(errors.New(err).Error())
	}

	return op.specificOperation.PaperRoll()
}

// PaperRoll returns the paper roll of OperationRollingCutting
func (op *operationRollingCutting) PaperRoll() *PaperRoll {
	return op.paperRoll
}

// Order is the base method and returns the order of the specific operations
func (op *Operation) Order() *Order {
	if op == nil {
		return nil
	}

	// Check for recursive call and panic.
	if isInfinite, err := IsInfiniteRecursiveCall(); isInfinite {
		panic(errors.New(err).Error())
	}

	return op.specificOperation.Order()
}

// Order returns the order of OperationFoldingCutting
func (op *operationFoldingPacking) Order() *Order {
	if op == nil {
		return nil
	}

	return op.order
}

// Task returns the task of the operations
func (op *Operation) Task() *Task {
	if op == nil {
		return nil
	}

	return op.task
}

// setPreviousOperation sets the first operations for the operations
func (op *Operation) setPreviousOperation(operation *Operation) {
	op.previousOperation = operation
}

// setNextOperation sets the first operations for the operations
func (op *Operation) setNextOperation(operation *Operation) {
	op.nextOperation = operation
}

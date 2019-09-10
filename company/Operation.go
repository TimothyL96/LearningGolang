package company

import (
	"errors"
)

// specificOperation interface is the interface for all operations type.
// This allows polymorphism for 1 operation type to other operation type
type specificOperation interface {
	FirstOperation() *Operation
	LastOperation() *Operation
	PaperRoll() *PaperRoll
	Order() *Order
}

// Operation struct is the struct for operation
type Operation struct {
	key
	isPlanned     bool
	operationType byte

	// Subclass
	specificOperation

	// Relation
	firstOperation *Operation
	lastOperation  *Operation
}

// operationRollingCutting is the base struct for rolling and cutting operation
type operationRollingCutting struct {
	specificOperation

	// Owner
	paperRoll *PaperRoll
}

// operationFoldingPacking is the base struct for folding and packing operation
type operationFoldingPacking struct {
	specificOperation

	// Owner
	order *Order
}

// operationRolling is the struct for rolling operation
type operationRolling struct {
	*Operation
}

// operationCutting is the struct for cutting operation
type operationCutting struct {
	*Operation
}

// operationFolding is the struct for folding operation
type operationFolding struct {
	*Operation
}

// operationPacking is the struct for packing operation
type operationPacking struct {
	*Operation
}

// SetIsPlanned accepts a bool parameter and sets it to the operation isPlanned field
func (op *Operation) SetIsPlanned(isPlanned bool) {
	op.isPlanned = isPlanned
}

// IsPlanned gets the isPlanned value of the operation
func (op *Operation) IsPlanned() *bool {
	if op == nil {
		return nil
	}

	return &op.isPlanned
}

// FirstOperation returns the first operation of the operation like rolling or folding
func (op *Operation) FirstOperation() *Operation {
	if op == nil {
		return nil
	}

	return op.firstOperation
}

// LastOperation returns the last operation of the operation like cutting or packing
func (op *Operation) LastOperation() *Operation {
	if op == nil {
		return nil
	}

	return op.lastOperation
}

// PaperRoll is the base method and returns the paper roll of the specific operation
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

// Order is the base method and returns the order of the specific operation
func (op *Operation) Order() *Order {
	if op == nil {
		return nil
	}

	return op.specificOperation.Order()
}

// Order returns the order of OperationFoldingCutting
func (op *operationFoldingPacking) Order() *Order {
	return op.order
}

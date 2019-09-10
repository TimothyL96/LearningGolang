package company

// specificOperation interface is the interface for all operations type.
// This allows polymorphism for 1 operation type to other operation type
type specificOperation interface {
	FirstOperation() *Operation
	LastOperation() *Operation
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
	*operationRollingCutting
}

// operationCutting is the struct for cutting operation
type operationCutting struct {
	*operationRollingCutting
}

// operationFolding is the struct for folding operation
type operationFolding struct {
	*operationFoldingPacking
}

// operationPacking is the struct for packing operation
type operationPacking struct {
	*operationFoldingPacking
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

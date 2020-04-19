package company

// operationFoldingPacking is the base struct for folding and packing operations
type operationFoldingPacking struct {
	*BaseOperation

	// Owner
	order *Order
}

// Order returns the order of OperationFoldingCutting
func (op *operationFoldingPacking) Order() *Order {
	if op == nil {
		return nil
	}

	return op.order
}

// Return operation folding packing
func (op *operationFoldingPacking) AsOperationFoldingPacking() *operationFoldingPacking {
	return op
}

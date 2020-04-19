package company

// operationPacking is the struct for packing operations
type operationPacking struct {
	*operationFoldingPacking
}

func (op *operationPacking) AsOperationPacking() *operationPacking {
	return nil
}

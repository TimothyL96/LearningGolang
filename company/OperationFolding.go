package company

// operationFolding is the struct for folding operations
type operationFolding struct {
	*operationFoldingPacking
}

func (op *operationFolding) AsOperationFolding() *operationFolding {
	return nil
}

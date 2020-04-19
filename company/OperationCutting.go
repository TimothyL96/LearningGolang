package company

// operationCutting is the struct for cutting operations
type operationCutting struct {
	*operationRollingCutting
}

func (op *operationCutting) AsOperationCutting() *operationCutting {
	return nil
}

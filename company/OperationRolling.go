package company

// operationRolling is the struct for rolling operations
type operationRolling struct {
	*operationRollingCutting
}

// Return operation rolling
func (op *operationRolling) AsOperationRolling() *operationRolling {
	return op
}

package company

// operationRollingCutting is the base struct for rolling and cutting operations
type operationRollingCutting struct {
	*BaseOperation

	// Owner
	paperRoll *PaperRoll
}

// PaperRoll returns the paper roll of OperationRollingCutting
func (op *operationRollingCutting) PaperRoll() *PaperRoll {
	return op.paperRoll
}

// Return operation rolling cutting
func (op *operationRollingCutting) AsOperationRollingCutting() *operationRollingCutting {
	return op
}

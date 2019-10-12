package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/LearningGolang/key"
)

// Order struct
type Order struct {
	key
	id                int
	color             int
	quantity          int
	dueDate           int
	fulfilledQuantity int // function

	// Owner
	company *Company

	// Relation
	knifeSetting   *KnifeSetting
	operations     []*Operation
	firstOperation *Operation
	lastOperation  *Operation
}

// Company returns the company of order
func (order *Order) Company() *Company {
	if order == nil {
		return nil
	}

	return order.company
}

// KnifeSetting returns the knife setting of order
func (order *Order) KnifeSetting() *KnifeSetting {
	if order == nil {
		return nil
	}

	return order.knifeSetting
}

// ID returns the id of order
func (order *Order) ID() int {
	if order == nil {
		panic(errors.New("order is nil").Error())
	}

	return order.id
}

// Operations returns all the operations
func (order *Order) Operations() []*Operation {
	if order == nil {
		return nil
	}

	return order.operations
}

// FirstOperation returns the folding operation
func (order *Order) FirstOperation() *Operation {
	if order == nil {
		return nil
	}

	return order.firstOperation
}

// LastOperation returns the packing operation
func (order *Order) LastOperation() *Operation {
	if order == nil {
		return nil
	}

	return order.lastOperation
}

// setKnifeSetting will set a knife setting to this order
func (order *Order) setKnifeSetting(ks *KnifeSetting) {
	order.knifeSetting = ks
}

// createOperation create rolling and cutting operations for this paper roll
func (order *Order) createOperation() []*Operation {
	operationFolding := order.createOperationFolding()
	operationPacking := order.createOperationPacking()

	operationFolding.setNextOperation(operationPacking)
	operationPacking.setPreviousOperation(operationFolding)

	operations := []*Operation{operationFolding, operationPacking}

	order.operations = operations
	order.firstOperation = operationFolding
	order.lastOperation = operationPacking

	return operations
}

// createOperationFolding create folding operation
func (order *Order) createOperationFolding() *Operation {
	operation := &Operation{
		key:               keyConfiguration.NewKey(),
		isPlanned:         false,
		operationType:     'F',
		specificOperation: nil,
		previousOperation: nil,
		nextOperation:     nil,
		task:              nil,
	}

	operationFolding := &operationFolding{
		operation,
	}

	operation.specificOperation = &operationFoldingPacking{
		specificOperation: operationFolding,
		order:             order,
	}

	return operation
}

// createOperationPacking create packing operation
func (order *Order) createOperationPacking() *Operation {
	operation := &Operation{
		key:               keyConfiguration.NewKey(),
		isPlanned:         false,
		operationType:     'P',
		specificOperation: nil,
		previousOperation: nil,
		nextOperation:     nil,
		task:              nil,
	}

	operationPacking := &operationPacking{
		operation,
	}

	operation.specificOperation = &operationFoldingPacking{
		specificOperation: operationPacking,
		order:             order,
	}

	return operation
}

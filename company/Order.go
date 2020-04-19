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
	operations     []Operation
	firstOperation Operation
	lastOperation  Operation
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

// Operations returns all the operations
func (order *Order) Operations() []Operation {
	if order == nil {
		return nil
	}

	return order.operations
}

// FirstOperation returns the folding operation
func (order *Order) FirstOperation() Operation {
	if order == nil {
		return nil
	}

	return order.firstOperation
}

// LastOperation returns the packing operation
func (order *Order) LastOperation() Operation {
	if order == nil {
		return nil
	}

	return order.lastOperation
}

// ID returns the id of order
func (order *Order) ID() int {
	if order == nil {
		panic(errors.New("order is nil").Error())
	}

	return order.id
}

// Color returns the color of order
func (order *Order) Color() int {
	if order == nil {
		panic(errors.New("order is nil").Error())
	}

	return order.color
}

// Quantity returns the quantity of order
func (order *Order) Quantity() int {
	if order == nil {
		panic(errors.New("order is nil").Error())
	}

	return order.quantity
}

// DueDate returns the due date of order
func (order *Order) DueDate() int {
	if order == nil {
		panic(errors.New("order is nil").Error())
	}

	return order.dueDate
}

// FulfilledQuantity returns the fulfilled quantity of order
func (order *Order) FulfilledQuantity() int {
	if order == nil {
		panic(errors.New("order is nil").Error())
	}

	return order.fulfilledQuantity
}

// SetKnifeSetting will set a knife setting to this order
func (order *Order) SetKnifeSetting(ks *KnifeSetting) {
	order.knifeSetting = ks
}

// createOperations create rolling and cutting operations for this paper roll
func (order *Order) createOperations() []Operation {
	operationFolding := order.createOperationFolding()
	operationPacking := order.createOperationPacking()

	operationFolding.setNextOperation(operationPacking)
	operationPacking.setPreviousOperation(operationFolding)

	operations := []Operation{operationFolding, operationPacking}

	order.operations = operations
	order.firstOperation = operationFolding
	order.lastOperation = operationPacking

	return operations
}

// createOperationFolding create folding operation
func (order *Order) createOperationFolding() Operation {
	operation := &operationFolding{
		order.createOperation('F'),
	}

	return operation
}

// createOperationPacking create packing operation
func (order *Order) createOperationPacking() Operation {
	operation := &operationPacking{
		order.createOperation('P'),
	}

	return operation
}

func (order *Order) createOperation(oType byte) *operationFoldingPacking {
	baseOperation := &BaseOperation{
		key:               keyConfiguration.NewKey(),
		isPlanned:         false,
		operationType:     oType,
		previousOperation: nil,
		nextOperation:     nil,
		task:              nil,
	}

	rcOperation := &operationFoldingPacking{
		BaseOperation: baseOperation,
		order:         order,
	}

	return rcOperation
}

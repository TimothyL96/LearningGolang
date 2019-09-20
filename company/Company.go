package company

import (
	"errors"

	keyConfiguration "github.com/ttimt/QuiLite/key"
)

// Map the Key struct from key package to the type key so it will be unexported
// (instead of using keyword interface 'Key' where it will become exported if used)
type key = *keyConfiguration.Key

// Company struct represents the root instance of the dataset
type Company struct {
	key
	version  float32
	dateTime int

	// Owning objects
	machines      []*Machine
	knifeSettings []*KnifeSetting
	orders        []*Order
}

// CreateCompany creates the root company instance and returns a pointer of it
func CreateCompany(version float32, dateTime int) *Company {
	company := &Company{
		key:           keyConfiguration.NewKey(),
		version:       version,
		dateTime:      dateTime,
		machines:      nil,
		knifeSettings: nil,
	}

	return company
}

// CreateMachine creates a single machine that is owned by the company.
//
// Machine type can only be 'R', 'C', 'F', or 'P'
func (company *Company) CreateMachine(name string, machineType byte) *Machine {
	if isValid := company.IsValidMachineType(machineType); !isValid {
		panic(errors.New("machine is being created with invalid type " + string(machineType)).Error())
	}

	machine := &Machine{
		key:         keyConfiguration.NewKey(),
		name:        name,
		machineType: machineType,
		company:     company,
		tasks:       nil,
		firstTask:   nil,
		lastTask:    nil,
	}

	// Add it to company unsorted
	company.machines = append(company.machines, machine)

	return machine
}

// CreateKnifeSetting creates a knife setting owned by the company
func (company *Company) CreateKnifeSetting(numberOfCut, color, repetition int) *KnifeSetting {
	knifeSetting := &KnifeSetting{
		key:         keyConfiguration.NewKey(),
		numberOfCut: numberOfCut,
		color:       color,
		repetition:  repetition,
		company:     company,
		paperRoll:   nil,
		orders:      nil,
	}

	// Add it to company unsorted
	company.knifeSettings = append(company.knifeSettings, knifeSetting)

	return knifeSetting
}

// CreateOrder creates an order owned by the company
func (company *Company) CreateOrder(id, color, quantity, dueDate, fulfilledQuantity int) *Order {
	order := &Order{
		key:               keyConfiguration.NewKey(),
		id:                id,
		color:             color,
		quantity:          quantity,
		dueDate:           dueDate,
		fulfilledQuantity: fulfilledQuantity,
		company:           company,
		knifeSetting:      nil,
	}

	// Add it to company unsorted
	company.orders = append(company.orders, order)

	// Create operations for order
	order.createOperation()

	return order
}

// IsValidMachineType checks if the input machine type is valid (rolling, cutting, folding, or packing) and returns a bool
func (company *Company) IsValidMachineType(machineType byte) bool {
	return machineType == rolling || machineType == cutting || machineType == folding || machineType == packing
}

// SetDateTime sets the date time for the company
func (company *Company) SetDateTime(dateTime int) {
	company.dateTime = dateTime

	for _, x := range company.machines {
		x.firstTask.setStartDateTime()
	}
}

// Machines will return all machines owned by this company
func (company *Company) Machines() []*Machine {
	return company.machines
}

// Orders will return all orders owned by this company
func (company *Company) Orders() []*Order {
	return company.orders
}

// knifeSettings will return all knife settings owned by this company
func (company *Company) KnifeSettings() []*KnifeSetting {
	return company.knifeSettings
}

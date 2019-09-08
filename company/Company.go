package company

import (
	"errors"
	"reflect"

	keyConfiguration "github.com/ttimt/GolangWebSocket/key"
)

// Map the interface Key from key package to the type key so it will be unexported (instead of using interface Key)
type key = keyConfiguration.Key

// Company struct represents the root instance of the dataset
type Company struct {
	key
	version  float32
	dateTime int

	// Owning objects
	machines      []*Machine
	KnifeSettings []*KnifeSetting
}

// CreateCompany creates the root company instance and returns a pointer of it
func CreateCompany(version float32, dateTime int) *Company {
	company := &Company{
		key:           keyConfiguration.NewKey(),
		version:       version,
		dateTime:      dateTime,
		machines:      nil,
		KnifeSettings: nil,
	}

	return company
}

// CreateMachine creates a single machine that is owned by the company.
//
// Machine type can only be 'R', 'C', 'F', or 'P'
func (company *Company) CreateMachine(name string, machineType byte) *Machine {
	// Machine type can only be rolling, cutting, folding, or packing
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
func (company *Company) CreateKnifeSetting(numberOfCut, color int) *KnifeSetting {
	knifeSetting := &KnifeSetting{
		key:         keyConfiguration.NewKey(),
		numberOfCut: numberOfCut,
		color:       color,
		company:     company,
		paperRoll:   nil,
		orders:      nil,
	}

	return knifeSetting
}

// CreateOrder creates an order owned by the company
func (company *Company) CreateOrder(ID, color, quantity, dueDate int) *Order {
	order := &Order{
		key:          keyConfiguration.NewKey(),
		ID:           ID,
		color:        color,
		quantity:     quantity,
		dueDate:      dueDate,
		company:      company,
		knifeSetting: nil,
	}

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

// Guard will return value of the 2nd parameter if first is invalid.
//
// Simplify the checking especially for a chain of pointer fields such as A.B.C.D
//
// First parameter must be a pointer
func Guard(currentValue interface{}, defaultValue interface{}) interface{} {
	if reflect.TypeOf(currentValue).Kind() != reflect.Ptr {
		panic(errors.New("non pointer passed to Guard function").Error())
	}

	// Get pointer of current value
	currentValuePtr := reflect.ValueOf(currentValue)

	// Return currentValue if it's not nil
	if !currentValuePtr.IsNil() {
		return currentValuePtr.Elem().Interface()
	}

	return defaultValue
}

// CalcDeclarative accepts 3 parameters where the first 2 are compared to see if their value are different.
//
// If the value is different, the value of the 2nd parameter will be set to the first parameter, then the slice of functions in parameter 3 will be executed 1 by 1.
//
// The first and second parameter must always be a pointer
func CalcDeclarative(currentValue interface{}, delta interface{}, funcToRuns ...func()) {
	if reflect.TypeOf(currentValue).Kind() != reflect.Ptr || reflect.TypeOf(delta).Kind() != reflect.Ptr {
		panic(errors.New("non pointer value received when calculating function").Error())
	}

	// Get the pointer of the two values
	currentValuePtr := reflect.ValueOf(currentValue)
	deltaPtr := reflect.ValueOf(delta)

	if currentValuePtr.Elem().Interface() != deltaPtr.Elem().Interface() {
		currentValuePtr.Elem().Set(deltaPtr.Elem())

		// Run all the functions to propagate
		for _, funcToRun := range funcToRuns {
			funcToRun()
		}
	}
}

// CalcFunction is generally used to calculate functions of variables (non pointer type)
// due to Get method for normal fields returning non-pointer value compared to Get method of relation/pointer fields.
//
// This can be considered as a wrapper to CalcDeclarative for propagating variables' declarative functions
//
// Compared to CalcDeclarative, this method accepts non-pointer type in first parameter, and a third parameter that accepts the pointer of the field to be mutated
// if the first parameter value and the value of the pointer in the second parameter is different.
//
// Panic if first parameter is a pointer or if third parameter is a non-pointer type
func CalcFunction(currentValue interface{}, delta interface{}, addrCurrentValue interface{}, funcToRuns ...func()) {
	if reflect.TypeOf(currentValue).Kind() == reflect.Ptr || reflect.TypeOf(addrCurrentValue).Kind() != reflect.Ptr {
		panic(errors.New("use CalcDeclarative if calculating relation. This is for variables where first parameter is non pointer and third parameter is pointer of field to mutate").Error())
	}

	// Store the existing value of current value
	existingValue := reflect.ValueOf(currentValue)

	// Call the main declarative logic method with address of currentValue
	CalcDeclarative(&currentValue, delta, funcToRuns...)

	// If currentValue is mutated, mutate the value of pointer of the third parameter
	if existingValue != currentValue {
		reflect.ValueOf(addrCurrentValue).Elem().Set(existingValue)
	}
}

// Machines will return all machines owned by this company
func (company *Company) Machines() []*Machine {
	return company.machines
}

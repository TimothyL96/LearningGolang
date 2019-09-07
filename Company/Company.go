package company

import (
	"errors"
	"reflect"

	keyConfiguration "github.com/ttimt/GolangWebSocket/key"
)

// Map the interface Key from key package to the type key so it will be unexported ( instead of using interface Key)
type key = keyConfiguration.Key

// Company struct represent the root instance of the dataset
type Company struct {
	key
	version  float32
	dateTime int

	// Owning objects
	machines      []*Machine
	KnifeSettings []*KnifeSetting
}

// Set machine type in constant
const (
	Rolling = 'R'
	Cutting = 'C'
	Folding = 'F'
	Packing = 'P'
)

// CreateMachine creates a single machine that is owned by the company
func (company *Company) CreateMachine(name string, machineType byte) *Machine {
	// Machine type can only be Rolling,C,F,P
	if machineType != Rolling && machineType != Cutting && machineType != Folding && machineType != Packing {
		panic(errors.New("machine is being created with invalid type" + string(machineType)).Error())
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
func Guard(check interface{}, defaultValue interface{}) interface{} {
	// Get value of check
	checkValue := reflect.ValueOf(check)

	// Return first value if it's not nil
	if !checkValue.IsNil() {
		return checkValue.Elem().Interface()
	}

	return defaultValue
}

// CalcFunc accepts 3 parameters where the first 2 are compared to see if their value are different.
//
// If the value is different, the value of the 2nd parameter will be set to the first parameter, then the slice of functions in parameter 3 will be executed 1 by 1.
//
// The first and second parameter must always be a pointer
func CalcFunc(currentValue interface{}, newValue interface{}, funcToRuns ...func()) {
	if reflect.TypeOf(currentValue).Kind() != reflect.Ptr || reflect.TypeOf(newValue).Kind() != reflect.Ptr {
		panic(errors.New("non pointer value received when calculating function").Error())
	}

	// Get the current pointer
	currentValuePtr := reflect.ValueOf(currentValue)
	newValuePtr := reflect.ValueOf(newValue)

	if currentValuePtr.Elem().Interface() != newValuePtr.Elem().Interface() {
		currentValuePtr.Elem().Set(newValuePtr.Elem())

		// Run all the functions to propagate
		for _, funcToRun := range funcToRuns {
			funcToRun()
		}
	}
}

// GetAllMachines will return all machines owned by this company
func (company *Company) GetAllMachines() []*Machine {
	return company.machines
}

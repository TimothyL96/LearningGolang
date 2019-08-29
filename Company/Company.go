package company

import (
	"errors"
	"reflect"
)

// Company struct
type Company struct {
	Version  int
	DateTime int

	// Owning objects
	Machines []*Machine
}

// CreateMachine method
func (company *Company) CreateMachine(MachineName string, MachineType byte) *Machine {
	machine := &Machine{
		MachineName: MachineName,
		MachineType: MachineType,
		Company:     company,
		Tasks:       nil,
		FirstTask:   nil,
		LastTask:    nil,
	}

	// Add it to company unsorted
	company.Machines = append(company.Machines, machine)

	return machine
}

// SetDateTime xaxa
func (company *Company) SetDateTime(dateTime int) {
	company.DateTime = dateTime

	for _, x := range company.Machines {
		x.FirstTask.SetStartDateTime()
	}
}

// Guard Set value to 2nd parameter if first is invalid
func Guard(check interface{}, defaultValue interface{}) interface{} {
	// Get value of check
	checkValue := reflect.ValueOf(check)

	if !checkValue.IsNil() {
		return checkValue.Convert(
			reflect.TypeOf(checkValue.Elem()),
		)
	}

	return defaultValue
}

// CalcFunc xaxa
func CalcFunc(currentValue interface{}, newValue interface{}, funcToRuns ...func()) {
	// If currentValue is null, panic
	if currentValue == nil {
		panic(errors.New("currentValue is null").Error())
	}

	// Get the current pointer
	currentValuePtr := reflect.ValueOf(currentValue)

	// TODO check the Kind for a struct and compare keys instead
	// reflect.TypeOf(currentValue).Kind() == reflect.Struct

	if currentValuePtr.Elem() != newValue {
		currentValuePtr.Elem().Set(reflect.ValueOf(newValue))

		// Run all the functions to recalculate
		for _, funcToRun := range funcToRuns {
			funcToRun()
		}
	}
}

// CalcFuncRelation xaxa
func CalcFuncRelation(currentValue interface{}, newValue interface{}, funcToRuns ...func()) interface{} {

	if reflect.ValueOf(currentValue).IsNil() {
		currentValue = reflect.New(reflect.TypeOf(currentValue))
	}

	CalcFunc(currentValue, newValue, funcToRuns...)

	return currentValue
}

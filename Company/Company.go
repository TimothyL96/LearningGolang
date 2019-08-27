package company

import (
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

	if !reflect.ValueOf(check).IsNil() {
		return reflect.ValueOf(check).Elem().Convert(
			reflect.TypeOf(reflect.ValueOf(check).Elem()),
		)
	}

	return defaultValue
}

// CalcFunc xaxa
func CalcFunc(currentValue interface{}, newValue interface{}, funcToRun func()) {
	if currentValue != newValue {
		reflect.ValueOf(currentValue).Elem().Set(reflect.ValueOf(newValue))

		funcToRun()
	}
}

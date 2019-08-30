package company

import (
	"reflect"
	"strconv"
)

// Company struct
type Company struct {
	Version         int
	DateTime        int
	SiteKey         int // Should be changed to key and store current major and minor key globally
	MajorKeyCurrent int
	MinorKeyCurrent int

	// Owning objects
	Machines []*Machine
}

// Key struct
type Key struct {
	SiteKey  int
	MajorKey int
	MinorKey int
}

// CreateMachine method
func (company *Company) CreateMachine(MachineName string, MachineType byte) *Machine {
	machine := &Machine{
		Key:         company.GetNewKey(),
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
	// Check if parameter 1 and 2 have the same type

	// Get value of check
	checkValue := reflect.ValueOf(check)

	if !checkValue.IsNil() {
		return checkValue.Elem().Interface()
	}

	return defaultValue
}

// Check if loop in function cause less calculation

// CalcFunc xaxa
func CalcFunc(currentValue interface{}, newValue interface{}, funcToRuns ...func()) {
	// Get the current pointer
	currentValuePtr := reflect.ValueOf(currentValue)
	newValuePtr := reflect.ValueOf(newValue)

	if currentValuePtr.Elem().Interface() != newValuePtr.Elem().Interface() {
		currentValuePtr.Elem().Set(newValuePtr.Elem())

		// Run all the functions to recalculate
		for _, funcToRun := range funcToRuns {
			funcToRun()
		}
	}
}

// ToString for key
func (key Key) ToString() string {
	return strconv.Itoa(key.SiteKey) + "." + strconv.Itoa(key.MajorKey) + "." + strconv.Itoa(key.MinorKey)
}

// GetNewKey xaxa
func (company *Company) GetNewKey() Key {
	key := Key{
		SiteKey:  company.SiteKey,
		MajorKey: company.MajorKeyCurrent,
		MinorKey: company.MinorKeyCurrent,
	}

	company.MinorKeyCurrent++

	return key
}

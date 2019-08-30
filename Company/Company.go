package company

import (
	"errors"
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

	if currentValuePtr.Elem() != newValue {
		if reflect.TypeOf(newValue).Kind() != reflect.Ptr {
			currentValuePtr.Elem().Set(reflect.ValueOf(newValue))
		} else {
			currentValuePtr.Elem().Set(reflect.ValueOf(newValue).Elem())
		}
	}

	// Run all the functions to recalculate
	for _, funcToRun := range funcToRuns {
		funcToRun()
	}
}

// CalcRelation xaxa
func CalcRelation(currentValue interface{}, newValue interface{}, funcToRuns ...func()) interface{} {
	if reflect.ValueOf(currentValue).IsNil() {
		currentValue = reflect.New(reflect.TypeOf(currentValue))
	}

	CalcFunc(&currentValue, newValue, funcToRuns...)

	return &currentValue
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

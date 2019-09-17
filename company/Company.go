package company

import (
	"errors"
	"reflect"
	"runtime"
	"strings"

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

// Machines will return all machines owned by this company
func (company *Company) Machines() []*Machine {
	return company.machines
}

// CalcDeclarative accepts 3 parameters where the first 2 are compared to see if their value are different.
//
// If the value is different, the value of the 2nd parameter will be set to the first parameter, then the slice of functions in parameter 3 will all be executed.
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

// IsInfiniteRecursiveCall method will check for recursive call and panic
func IsInfiniteRecursiveCall() (isInfinite bool, err string) {
	// "Current method" has the ID excluding current method which is 0
	const currentMethodID, callerMethodID = 1, 2

	// Get the program counter of current and previous method
	// pc = Program counter
	currentPC, _, _, currentIsValid := runtime.Caller(currentMethodID)
	callerPC, _, _, callerIsValid := runtime.Caller(callerMethodID)

	if !currentIsValid || !callerIsValid {
		err = "can't retrieve program counter of caller or current method"
	}

	// Get the details of current and previous method from the program counter
	currentDetail := runtime.FuncForPC(currentPC)
	previousDetail := runtime.FuncForPC(callerPC)

	// Split the details separated by dot
	currentNameType := strings.Split(currentDetail.Name(), ".")
	previousNameType := strings.Split(previousDetail.Name(), ".")

	// Get the last 2 elements which contain the current type and method
	currentNameType = currentNameType[len(currentNameType)-2:]
	previousNameType = previousNameType[len(previousNameType)-2:]

	// Compare the previous and current's name and type
	// Panic if they are the same (infinite recursive call)
	if reflect.DeepEqual(currentNameType, previousNameType) {
		isInfinite = true
		err = "fatal error: infinite recursive call on " + strings.Join(currentNameType, ", ")
	}

	return
}

// Traverse
func Traverse(object interface{}, pathOfRelation string, logicToExecute interface{}) {
	if reflect.ValueOf(object).Kind() != reflect.Ptr {
		panic("non relation passed to first parameter for Traverse")
	}
	// Panic if first parameter is not unary

	instances := make([]interface{}, 0)
	paths := strings.Split(pathOfRelation, ".")

	for _, p := range paths {
		// Check if current value in object is slice
		if reflect.ValueOf(object).Kind() == reflect.Slice {
			instances = nil
			for i := 0; i < reflect.ValueOf(object).Len(); i++ {
				j := reflect.ValueOf(object).Index(i).MethodByName(p).Call(nil)[0].Interface()

				for z := 0; z < reflect.ValueOf(j).Len(); z++ {
					instances = append(instances, reflect.ValueOf(j).Index(z).Interface())
				}
			}

			object = reflect.ValueOf(instances)
		} else {
			// Set the object to the current unary relation
			object = reflect.ValueOf(object).MethodByName(p).Call(nil)[0].Interface()
		}
	}

	for _, v := range instances {
		reflect.ValueOf(logicToExecute).Call([]reflect.Value{reflect.ValueOf(v)})
	}
}

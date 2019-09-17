package stdlib

import (
	"errors"
	"reflect"
)

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

package stdlib

import (
	"errors"
	"reflect"
	"runtime"
	"strings"
)

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
func Traverse(instance interface{}, relationPath string, logicToExecute interface{}) {
	if isValid, err := traverseCheckError(instance, logicToExecute); !isValid {
		panic(errors.New(err).Error())
	}

	// Get all elements from path
	retrievedInstances := traverseRetrievePath(instance, relationPath)

	// Execute the logic of traverse for every instance
	for _, instance := range retrievedInstances {
		instanceValue := reflect.ValueOf(instance)

		// Avoid nil instance
		if instanceValue.IsNil() {
			continue
		}

		reflect.ValueOf(logicToExecute).Call([]reflect.Value{instanceValue}) // see if possible to check for type assertion error
	}
}

// Counter counts all the elements with the filter applied
func Counter(instance interface{}, relationPath string, logicToExecute interface{}) (count int) {
	if isValid, err := traverseCheckError(instance, logicToExecute); !isValid {
		panic(errors.New(err).Error())
	}

	// Get all elements from path
	retrievedInstances := traverseRetrievePath(instance, relationPath)

	// Execute the logic of traverse for every instance
	for _, instance := range retrievedInstances {
		instanceValue := reflect.ValueOf(instance)

		// Avoid nil instance
		if instanceValue.IsNil() || !reflect.ValueOf(logicToExecute).Call([]reflect.Value{instanceValue})[0].Interface().(bool) {
			continue
		}

		count++
	}

	return
}

// Select will select the first instance that matches the filter
func Select(instance interface{}, relationPath string, logicToExecute interface{}) (returnInstance interface{}) {
	if isValid, err := traverseCheckError(instance, logicToExecute); !isValid {
		panic(errors.New(err).Error())
	}

	// Get all elements from path
	retrievedInstances := traverseRetrievePath(instance, relationPath)

	// Execute the logic of traverse for every instance
	for _, instance := range retrievedInstances {
		instanceValue := reflect.ValueOf(instance)

		// Avoid nil instance
		if instanceValue.IsNil() || !reflect.ValueOf(logicToExecute).Call([]reflect.Value{instanceValue})[0].Interface().(bool) {
			continue
		}

		return instance
	}

	return
}

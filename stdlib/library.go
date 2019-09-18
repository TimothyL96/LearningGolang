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

// Retrieve all the elements from the given relation path
func traverseRetrievePath(instance interface{}, relationPath string) []interface{} {
	var retrievedInstances []interface{}
	paths := strings.Split(relationPath, ".")

	// Go through all the paths in parameter two
	for _, path := range paths {
		instanceValue := reflect.ValueOf(instance)
		retrievedInstances = nil

		// Check if current value in instance is slice
		if instanceValue.Kind() == reflect.Slice {
			// Call the method for each instance in instance
			for i := 0; i < instanceValue.Len(); i++ {
				retrievedInstances = traverseInsertToSlice(reflect.ValueOf(instanceValue.Index(i).Interface()).MethodByName(path).Call(nil)[0].Interface(), retrievedInstances)
			}
		} else {
			// Set the instance to the current unary relation
			retrievedInstances = traverseInsertToSlice(instanceValue.MethodByName(path).Call(nil)[0].Interface(), retrievedInstances)
		}

		// Update instance with the next set of retrievedInstances to be traversed
		instance = retrievedInstances
	}

	return retrievedInstances
}

// If retrieve value is slice, traverse them 1 by 1 to store them
func traverseInsertToSlice(retrievedInstance interface{}, storeSlice []interface{}) []interface{} {
	retrievedInstanceValue := reflect.ValueOf(retrievedInstance)

	if retrievedInstanceValue.Kind() == reflect.Slice {
		for i := 0; i < retrievedInstanceValue.Len(); i++ {
			storeSlice = append(storeSlice, retrievedInstanceValue.Index(i).Interface())
		}
	} else {
		storeSlice = append(storeSlice, retrievedInstanceValue.Interface())
	}

	return storeSlice
}

// Check for correctness of input of traverse parameter
func traverseCheckError(instance interface{}, logicToExecute interface{}) (isValid bool, err string) {
	instanceValue := reflect.ValueOf(instance)

	if instanceValue.Kind() == reflect.Slice {
		// Panic if first parameter is not unary
		err = "non unary relation in first parameter"
	} else if instanceValue.Kind() != reflect.Ptr {
		err = "non relation passed to first parameter for Traverse"
	} else if reflect.ValueOf(logicToExecute).Kind() != reflect.Func {
		err = "non function passed to the third parameter"
	}

	isValid = err == ""
	return
}

// Counter counts all the elements with the filter applied
func Counter(instance interface{}, relationPath string, logicToExecute interface{}) int {
	return 0
}

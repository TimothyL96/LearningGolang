package stdlib

import (
	"errors"
	"reflect"
	"strings"
)

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
				methodToCall := reflect.ValueOf(instanceValue.Index(i).Interface()).MethodByName(path)

				if !methodToCall.IsValid() {
					panic(errors.New("relation error: " + path).Error())
				}
				retrievedInstances = traverseInsertToSlice(methodToCall.Call(nil)[0].Interface(), retrievedInstances)
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

package main

import (
	"fmt"
	"log"

	"./Company"
)

// DataSetsStorage struct
type DataSetsStorage struct {
	Companies []*Company.Company
}

// DataSets abc
var DataSets DataSetsStorage

// PrintAllDataSets prints
func (DataSetsStorage DataSetsStorage) printAllDataSets() {
	// Print all company data sets
	for _, ds := range DataSetsStorage.Companies {
		fmt.Println("Data set printing: ", ds)

		for _, m := range ds.Machines {
			fmt.Println("Machine printing: ", m)
		}
	}
}

// CreateDataSet creates
func createDataSet(DataSetKind string) interface{} {
	var DataSet interface{}

	switch DataSetKind {
	case "Company":
		DataSet = newCompany()
	default:
		log.Fatal("No such data set kind found!")
	}

	return DataSet
}

// NewCompany company
func newCompany() Company.Company {
	company := Company.Company{
		Version: 1,
	}

	return company
}

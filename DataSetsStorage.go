package main

import (
	"fmt"
	"log"
	"strings"

	DScompany "github.com/ttimt/LearningGolang/company"
)

// DataSetsStorage struct
type DataSetsStorage struct {
	Companies []*DScompany.Company
}

// DataSets abc
var DataSets DataSetsStorage

// PrintAllDataSets prints
func (DataSetsStorage DataSetsStorage) printAllDataSets() {
	// Print all company data sets
	for _, ds := range DataSetsStorage.Companies {
		fmt.Println("Data set (company) printing: ", ds)

		for _, m := range ds.Machines() {
			fmt.Println("Machine printing: ", m)

			for _, t := range m.Tasks() {
				fmt.Println("specificTask printing: ", t)
			}
		}
	}
}

// CreateDataSet creates
func createDataSet(DataSetKind string) interface{} {
	var DataSet interface{}

	switch strings.ToLower(DataSetKind) {
	case "company":
		DataSet = DScompany.CreateCompany(1.0, 0)
	default:
		log.Fatal("No such data set kind found!")
	}

	return DataSet
}

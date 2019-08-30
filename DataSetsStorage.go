package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/ttimt/GolangWebSocket/company"
)

// DataSetsStorage struct
type DataSetsStorage struct {
	Companies []*company.Company
}

// DataSets abc
var DataSets DataSetsStorage

// PrintAllDataSets prints
func (DataSetsStorage DataSetsStorage) printAllDataSets() {
	// Print all company data sets
	for _, ds := range DataSetsStorage.Companies {
		fmt.Println("Data set (company) printing: ", ds)

		for _, m := range ds.Machines {
			fmt.Println("Machine printing: ", m)

			for _, t := range m.Tasks {
				fmt.Println("Task printing: ", t)
			}
		}
	}
}

// CreateDataSet creates
func createDataSet(DataSetKind string) interface{} {
	var DataSet interface{}
	
	switch strings.ToLower(DataSetKind) {
	case "company":
		DataSet = newCompany()
	default:
		log.Fatal("No such data set kind found!")
	}

	return DataSet
}

// NewCompany company
func newCompany() company.Company {
	company := company.Company{
		Version: 1,
		DateTime: 0,
		SiteKey: 12345,
	}

	return company
}

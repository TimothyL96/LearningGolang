package main

import (
	"./Company"
	. "fmt"
	"log"
)

type DataSetsStorage struct {
	Companies[] *Company.Company
}

var DataSets DataSetsStorage

func (DataSetsStorage DataSetsStorage) PrintAllDataSets() {
	// Print all company data sets
	for _, ds := range DataSetsStorage.Companies{
		Println("Data set printing: ", ds)
		
		for _, m := range ds.Machines {
			Println("Machine printing: ", m)
		}
	}
}

func CreateDataSet(DataSetKind string) interface{} {
	var DataSet interface{}
	
	switch DataSetKind {
		case "Company":
			DataSet = NewCompany()
		default:
			log.Fatal("No such data set kind found!")
	}
	
	return DataSet
}

func NewCompany() Company.Company {
	company := Company.Company{
		Version:1,
	}
	
	return company
}
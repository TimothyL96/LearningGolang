package main

import (
	"errors"
	"fmt"
	"runtime/debug"

	DSCompany "github.com/ttimt/LearningGolang/company"
	_ "github.com/ttimt/LearningGolang/stdlib"
)

func main() {
	// *********** Test QT Golang GUI
	// app := widgets.NewQApplication(len(os.Args), os.Args)

	// // NewTestForm().Show()
	// ui.NewMainWindow().Show()

	// app.Exec()
	// End Test QT Golang GUI

	// Get input to create new data set
	var datasetKind string
	fmt.Println("Enter a data set kind to be created (company):")
	_, _ = fmt.Scanln(&datasetKind)

	if datasetKind == "" {
		datasetKind = "company"
	}

	convertDataSet(datasetKind)

	fmt.Println("Back to main")
}

func convertDataSet(DataSetKind string) {
	defer recoverPanic()

	// Server startup - Create the data set from the input using type assertion
	// From here, we are assuming the dataset or string entered is "company"
	DataSetInstance, succeed := (createDataSet(DataSetKind)).(*DSCompany.Company)

	if !succeed {
		panic(errors.New("type assertion failed").Error())
	} else {
		DataSets.Companies = append(DataSets.Companies, DataSetInstance)
		testCompany(DataSetInstance)
	}
}

func recoverPanic() {
	r := recover()
	if r != nil {
		fmt.Println("Recovered! from: ", r)
		fmt.Println(string(debug.Stack()))
	}
}

func testCompany(company *DSCompany.Company) {
	// Create a machine
	company.CreateMachine("1st machine", 'R')
	company.CreateMachine("2nd machine", 'C')
	company.CreateMachine("3rd machine", 'F')
	company.CreateMachine("4th machine", 'P')

	// // Create orders
	// for i := 1; i <= 30; i++ {
	// 	company.CreateOrder(i, rand.Int()%5, 5, 10, 2)
	// }
	//
	// // Create knife settings
	// for i := 1; i <= 15; i++ {
	// 	ks := company.CreateKnifeSetting(rand.Int()%10+1, rand.Int()%5, rand.Int()%20+1)
	//
	// 	ks.AssignOrder(Select(company, "Orders", func(order *DSCompany.Order) bool {
	// 		return order.ID() == i
	// 	}).(*DSCompany.Order))
	//
	// 	ks.AssignOrder(Select(company, "Orders", func(order *DSCompany.Order) bool {
	// 		return order.ID() == i+15
	// 	}).(*DSCompany.Order))
	// }

	// fmt.Println("Print all dataset:")
	// DataSets.printAllDataSets()

	// Print machines and tasks
	for _, m := range company.Machines() {
		// Create tasks
		m.CreateTask(123)
		m.CreateTask(123)

		fmt.Println("\n******************************************************************************")
		fmt.Println("Machine key: ", m.Key().String())
		fmt.Println("Machine Name: ", m.Name())
		fmt.Println("Machine Type: ", string(m.Type()))
		fmt.Println("Machine First BaseTask: ", m.FirstTask().Key())
		fmt.Println("Machine Last BaseTask: ", m.LastTask().Key())

		fmt.Println("\n***************************************")
		for _, t := range m.Tasks() {
			fmt.Println("Task key: ", t.Key())
			fmt.Println("Task type: ", string(t.TaskType()))
			fmt.Printf("Task type with %%T: %T\n", t)
			fmt.Println("Task start date time: ", t.StartDateTime())
			fmt.Println("Task end date time: ", t.EndDateTime())
			fmt.Println("Task duration: ", t.Duration())
			if t.TaskType() == 'R' {
				fmt.Println(t.AsTaskRolling().UniqueToRolling())
			}
			fmt.Println("\n***************************************")
		}
	}

	fmt.Println()

	// // Print orders
	// Traverse(company, "Orders", func(order *DSCompany.Order) {
	// 	fmt.Println("Order:", order.ID())
	// 	fmt.Println("Knife setting of order:", order.KnifeSetting().Key())
	// 	fmt.Println("BaseOperation folding:")
	// 	fmt.Println(order.FirstOperation().Key())
	// 	fmt.Println(order.FirstOperation().IsPlanned())
	// 	fmt.Println("BaseOperation packing:")
	// 	fmt.Println(order.LastOperation().Key())
	// 	fmt.Println(order.LastOperation().IsPlanned())
	// 	fmt.Println("Order.BaseOperation.OrderID", order.LastOperation().Order().ID())
	// })
	//
	// // Print knife settings and paper roll
	// Traverse(company, "KnifeSettings", func(ks *DSCompany.KnifeSetting) {
	// 	ks.CreatePaperRoll(ks.Color(), ks.NumberOfCut()*ks.Repetition())
	// 	fmt.Println("Knife setting:", ks.Key().String(), ks.Color())
	// 	fmt.Println("Paper roll:", ks.PaperRoll().Key().String(), "Length:", ks.PaperRoll().Length())
	// 	fmt.Println("BaseOperation rolling:")
	// 	fmt.Println(ks.PaperRoll().FirstOperation().Key())
	// 	fmt.Println(ks.PaperRoll().FirstOperation().IsPlanned())
	// 	fmt.Println(string(ks.PaperRoll().FirstOperation().OperationType()))
	// 	fmt.Println("BaseOperation cutting:")
	// 	fmt.Println(ks.PaperRoll().LastOperation().Key())
	// 	fmt.Println(ks.PaperRoll().LastOperation().IsPlanned())
	// 	fmt.Println(string(ks.PaperRoll().LastOperation().OperationType()))
	// })
}

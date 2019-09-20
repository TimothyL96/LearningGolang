package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime/debug"

	DSCompany "github.com/ttimt/QuiLite/company"
	. "github.com/ttimt/QuiLite/stdlib"
)

// var counter int
// var c client

// type client struct {
// 	m       sync.Mutex
// 	clients []*websocket.Conn
// }

// // Upgraded: Upgrade normal HTTP connection to WebSocket
// var upgrader = websocket.Upgrader{}

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// Upgrade initial GET request to a web socket
// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Make sure we close the connection when the function returns
// 	defer ws.Close()

// 	c.clients = append(c.clients, ws)

// 	// Infinite loop
// 	for {
// 		// Read in a new message as JSON and map it to a Message object
// 		_, m, err := ws.ReadMessage()

// 		if err != nil {
// 			if !websocket.IsCloseError(err, websocket.CloseNormalClosure,
// 				websocket.CloseGoingAway,
// 				websocket.CloseNoStatusReceived) {
// 				log.Println("ReadConnection:", err)
// 			} else {
// 				log.Println("Connection closed")
// 			}
// 			break
// 		} else {
// 			// println("ReadMessage succeeded:", string(b), time.Now().Unix())

// 			// c.m.Lock()
// 			for _, client := range c.clients {
// 				err := client.WriteMessage(websocket.TextMessage, m)
// 				if err != nil {
// 					log.Println("Write Connection:", err)
// 					break
// 				} else {
// 					println("WriteMessage succeeded:", time.Now().Unix())
// 				}
// 			}
// 			// c.m.Unlock()
// 		}
// 	}
// }

func main() {
	// *********** Test QT Golang GUI
	// app := widgets.NewQApplication(len(os.Args), os.Args)

	// // NewTestForm().Show()
	// ui.NewMainWindow().Show()

	// app.Exec()
	// End Test QT Golang GUI

	// *********** Golang Gorilla WebSocket
	// Create a simple file server
	// fs := http.FileServer(http.Dir("./"))
	// http.Handle("/", fs)

	// http.HandleFunc("/ws", handleConnections)

	// log.Println("http server stated on :8081")
	// err := http.ListenAndServe(":8081", nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
	// End Golang Gorilla WebSocket

	// *********** QTQ declarative calculation
	// Get input to create new data set
	var datasetKind string
	fmt.Println("Enter a data set kind to be created (company):")
	_, _ = fmt.Scanln(&datasetKind)

	if datasetKind == "" {
		datasetKind = "company"
	}

	convertDataSet(datasetKind)

	fmt.Println("Back to main")
	// End QTQ declarative calculation
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
	m1 := company.CreateMachine("1st machine", 'R')
	m2 := company.CreateMachine("2nd machine", 'C')

	for i := 1; i < 10; i++ {
		m1.CreateTask(rand.Int() % 10000)
	}
	m2.CreateTask(2)
	m2.CreateTask(7)
	m2.CreateTask(3)

	// Create orders
	for i := 1; i <= 30; i++ {
		company.CreateOrder(i, rand.Int()%5, 5, 10, 2)
	}

	// Create knife settings
	for i := 1; i <= 15; i++ {
		ks := company.CreateKnifeSetting(rand.Int()%10+1, rand.Int()%5, rand.Int()%20+1)

		ks.AssignOrder(Select(company, "Orders", func(order *DSCompany.Order) bool {
			return order.ID() == i
		}).(*DSCompany.Order))

		ks.AssignOrder(Select(company, "Orders", func(order *DSCompany.Order) bool {
			return order.ID() == i+15
		}).(*DSCompany.Order))
	}

	// fmt.Println("Print all dataset:")
	// DataSets.printAllDataSets()
	for _, m := range company.Machines() {
		fmt.Println("\n***************************************")
		fmt.Println("Machine key: ", m.Key().String())
		fmt.Println("Machine Name: ", m.Name())
		fmt.Println("Machine Type: ", string(m.Type()))
		fmt.Println("Machine First Task: ", m.FirstTask().Key())
		fmt.Println("Machine Last Task: ", m.LastTask().Key())

		for _, t := range m.Tasks() {
			fmt.Println("******")
			fmt.Println("Task key: ", t.Key().String())
			fmt.Println("Task type: ", string(t.TaskType()))
			fmt.Println("Duration: ", t.Duration())
			fmt.Println("Start date time: ", t.StartDateTime())
			fmt.Println("End date time: ", t.EndDateTime())
			if t.PreviousTask() != nil {
				fmt.Println("Previous task: ", t.PreviousTask().Key())
			}
			if t.NextTask() != nil {
				fmt.Println("Next task: ", t.NextTask().Key())
			}
		}
	}

	fmt.Println()

	// Print orders
	Traverse(company, "Orders", func(order *DSCompany.Order) {
		fmt.Println("Order:", order.ID())
		fmt.Println("Knife setting of order:", order.KnifeSetting().Key())
		fmt.Println("Operation folding:")
		fmt.Println(order.FirstOperation().Key())
		fmt.Println(order.FirstOperation().IsPlanned())
		fmt.Println("Operation packing:")
		fmt.Println(order.LastOperation().Key())
		fmt.Println(order.LastOperation().IsPlanned())
		fmt.Println("Order.Operation.OrderID", order.LastOperation().Order().ID())
	})

	// Print knife settings and paper roll
	Traverse(company, "KnifeSettings", func(ks *DSCompany.KnifeSetting) {
		ks.CreatePaperRoll(ks.Color(), ks.NumberOfCut()*ks.Repetition())
		fmt.Println("Knife setting:", ks.Key().String(), ks.Color())
		fmt.Println("Paper roll:", ks.PaperRoll().Key().String(), "Length:", ks.PaperRoll().Length())
		fmt.Println("Operation rolling:")
		fmt.Println(ks.PaperRoll().FirstOperation().Key())
		fmt.Println(ks.PaperRoll().FirstOperation().IsPlanned())
		fmt.Println(string(ks.PaperRoll().FirstOperation().OperationType()))
		fmt.Println("Operation cutting:")
		fmt.Println(ks.PaperRoll().LastOperation().Key())
		fmt.Println(ks.PaperRoll().LastOperation().IsPlanned())
		fmt.Println(string(ks.PaperRoll().LastOperation().OperationType()))
	})

	// Counter
	x := Counter(company, "Machines.Tasks.PreviousTask.Machine.Tasks", func(t *DSCompany.Task) bool {
		return t.TaskType() == 'R'
	})

	fmt.Println()
	fmt.Println("Counter quantor")
	fmt.Println("Result of counter:", x)

	// Select
	s := Select(company, "Machines", func(m *DSCompany.Machine) bool {
		return m.Name() == "1st machine"
	}).(*DSCompany.Machine)

	fmt.Println()
	fmt.Println("Select quantor")
	fmt.Println("Selected machine:")
	fmt.Println("Name -", s.Name())
	fmt.Println("Type -", string(s.Type()))
}

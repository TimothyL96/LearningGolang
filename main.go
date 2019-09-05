package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"runtime/debug"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/therecipe/qt/widgets"
	"github.com/ttimt/GolangWebSocket/company"
	"github.com/ttimt/GolangWebSocket/ui"
)

var counter int
var c client

type client struct {
	m       sync.Mutex
	clients []*websocket.Conn
}

// Upgraded: Upgrade normal HTTP connection to WebSocket
var upgrader = websocket.Upgrader{}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a web socket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	c.clients = append(c.clients, ws)

	// Infinite loop
	for {
		// Read in a new message as JSON and map it to a Message object
		_, m, err := ws.ReadMessage()

		if err != nil {
			if !websocket.IsCloseError(err, websocket.CloseNormalClosure,
				websocket.CloseGoingAway,
				websocket.CloseNoStatusReceived) {
				log.Println("ReadConnection:", err)
			} else {
				log.Println("Connection closed")
			}
			break
		} else {
			// println("ReadMessage succeeded:", string(b), time.Now().Unix())

			// c.m.Lock()
			for _, client := range c.clients {
				err := client.WriteMessage(websocket.TextMessage, m)
				if err != nil {
					log.Println("Write Connection:", err)
					break
				} else {
					println("WriteMessage succeeded:", time.Now().Unix())
				}
			}
			// c.m.Unlock()
		}
	}
}

func main() {
	// *********** Test QT Golang GUI
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// NewTestForm().Show()
	ui.NewMainWindow().Show()

	app.Exec()
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
	// var DataSetKind string
	// fmt.Println("Enter a data set kind to be created (company):")
	// fmt.Scanln(&DataSetKind)

	// convertDataSet(DataSetKind)

	// fmt.Println("Back to main")
	// End QTQ declarative calculation
}

func convertDataSet(DataSetKind string) {
	defer recoverPanic()

	// Server startup - Create the data set from the input using type assertion
	// From here, we are assuming the dataset or string entered is "company"
	DataSetInstance, succeed := (createDataSet(DataSetKind)).(company.Company)

	if !succeed {
		panic(errors.New("type assertion failed").Error())
	} else {
		DataSets.Companies = append(DataSets.Companies, &DataSetInstance)
		testCompany(&DataSetInstance)
	}
}

func recoverPanic() {
	r := recover()
	if r != nil {
		fmt.Println("Recovered! from: ", r)
		fmt.Println(string(debug.Stack()))
	}
}

func testCompany(company *company.Company) {
	// Create a machine
	m1 := company.CreateMachine("Golang first machine", 'T')
	m2 := company.CreateMachine("Golang second machine", 'C')

	firsttask := m1.CreateTask('A', 2)
	t := firsttask
	for i := 1; i < 10; i++ {
		t = m1.CreateTask('G', rand.Int()%10000)
	}
	m2.CreateTask('D', 2)
	m2.CreateTask('E', 7)
	m2.CreateTask('E', 3)

	// Print out last task of machine1
	// Change the duration
	// Then print out the task again
	fmt.Println("FIRST ###########")
	fmt.Println("Task key: ", t.Key.ToString())
	fmt.Println("Task type: ", string(t.TaskType))
	fmt.Println("Duration: ", t.Duration)
	fmt.Println("Start date time: ", t.StartDateTime)
	fmt.Println("End date time: ", t.EndDateTime)
	fmt.Println("Previous task: ", t.PreviousTask.GetKey())
	fmt.Println("Next task: ", t.NextTask.GetKey())
	t.SetDuration(10)
	fmt.Println("SECOND ############")
	fmt.Println("Task key: ", t.Key.ToString())
	fmt.Println("Task type: ", string(t.TaskType))
	fmt.Println("Duration: ", t.Duration)
	fmt.Println("Start date time: ", t.StartDateTime)
	fmt.Println("End date time: ", t.EndDateTime)
	fmt.Println("Previous task: ", t.PreviousTask.GetKey())
	fmt.Println("Next task: ", t.NextTask.GetKey())

	// fmt.Printf("First: \n%p\n\n", company)
	// fmt.Printf("Second: (TaskType) \n%+v\n\n", string(company.Machines[0].Tasks[0].TaskType))
	// fmt.Println("Print all dataset:")
	// DataSets.printAllDataSets()
	for _, m := range company.Machines {
		fmt.Println("***************************************")
		fmt.Println("Machine key: ", m.Key.ToString())
		fmt.Println("Machine Name: ", m.MachineName)
		fmt.Println("Machine Type: ", string(m.MachineType))
		fmt.Println("Machine First Task: ", m.FirstTask.GetKey())
		fmt.Println("Machine Last Task: ", m.LastTask.GetKey())

		for _, t := range m.Tasks {
			fmt.Println("******")
			fmt.Println("Task key: ", t.Key.ToString())
			fmt.Println("Task type: ", string(t.TaskType))
			fmt.Println("Duration: ", t.Duration)
			fmt.Println("Start date time: ", t.StartDateTime)
			fmt.Println("End date time: ", t.EndDateTime)
			fmt.Println("Previous task: ", t.PreviousTask.GetKey())
			fmt.Println("Next task: ", t.NextTask.GetKey())
		}
	}
}

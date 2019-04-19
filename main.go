package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"./Company"
	"github.com/gorilla/websocket"
)

var counter int
var mutex = &sync.Mutex{}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Upgraded: Upgrade normal HTTP connection to WebSocket
var upgrader = websocket.Upgrader{}

// Message this is a sample struct
// The text surrounded by backticks is just metadata (tag) which helps
// Go serialize and unserialize the Message object to and from JSON.
type Message struct {
	Email string `json:"email"`
	Name  string `json:"username"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a web socket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	// Infinite loop
	for {
		var msg Message

		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error json: %v", err)
			delete(clients, ws)
			break
		}

		// Send the newly received message to the channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the channel
		msg := <-broadcast

		// Send it out to every client this is currently connected
		for client := range clients {
			error := client.WriteJSON(msg)

			if error != nil {
				log.Printf("error1, %v", error)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {

	//Create a simple file server
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("http server stated on :8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// Get input to create new data set
	// var DataSetKind string
	// Println("Enter a data set kind to be created:")
	// Scanln(&DataSetKind)

	// ConvertDataSet(DataSetKind)

	// Println("Back to main")
}

func convertDataSet(DataSetKind string) {
	defer recoverPanic()

	// Server startup - Create the data set from the input
	DataSetInstance, succeed := (createDataSet(DataSetKind)).(Company.Company)

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
	}
}

func testCompany(company *Company.Company) {
	// Create a machine
	m1 := company.CreateMachine("Golang first machine", 'T')
	m2 := company.CreateMachine("Golang second machine", 'C')

	m1.CreateTask('A')
	m1.CreateTask('B')
	m1.CreateTask('C')
	m2.CreateTask('D')
	m2.CreateTask('E')

	fmt.Printf("%p\n", company)
	fmt.Printf("%+v\n", company.Machines[0].Tasks)

	DataSets.printAllDataSets()
}

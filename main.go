package main

import (
	// . "fmt"
	"log"
	"net/http"
	"sync"

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

	// // Get input to create new data set
	// // var DataSetKind string
	// // Println("Enter a data set kind to be created:")
	// // Scanln(&DataSetKind)

	// // ConvertDataSet(DataSetKind)

	// // Println("Back to main")

	// // http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
	// // 	Fprintf(writer, "Hello, %q", html.EscapeString(r.URL.Path))
	// // })

	// http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
	// 	Fprintf(w, "Hi")
	// })

	// http.HandleFunc("/increment", incrementCounter)

	// // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// // 	http.ServeFile(w, r, r.URL.Path[1:])
	// // })

	// http.Handle("/", http.FileServer(http.Dir("./Company")))
	// http.ListenAndServe(":8081", nil)
}

// func incrementCounter(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/increment" {
// 		panic("Not found 404!!")
// 	}

// 	mutex.Lock()
// 	counter++
// 	Fprint(w, strconv.Itoa(counter))
// 	mutex.Unlock()
// }
// func ConvertDataSet(DataSetKind string) {
// 	defer Recover()

// 	// Server startup - Create the data set from the input
// 	DataSetInstance, succeed := (CreateDataSet(DataSetKind)).(Company.Company)

// 	if !succeed {
// 		panic(errors.New("type assertion failed").Error())
// 	} else {
// 		DataSets.Companies = append(DataSets.Companies, &DataSetInstance)
// 		TestCompany(&DataSetInstance)
// 	}
// }

// func Recover() {
// 	r := recover()
// 	if r != nil {
// 		Println("Recovered! from: ", r)
// 	}
// }

// func TestCompany(company *Company.Company) {
// 	// Create a machine
// 	m1 := company.CreateMachine("Golang first machine", 'T')
// 	m2 := company.CreateMachine("Golang second machine", 'C')

// 	m1.CreateTask('A')
// 	m1.CreateTask('B')
// 	m1.CreateTask('C')
// 	m2.CreateTask('D')
// 	m2.CreateTask('E')

// 	Printf("%p\n", company)
// 	Printf("%+v\n", company.Machines[0].Tasks)

// 	DataSets.PrintAllDataSets()
// }

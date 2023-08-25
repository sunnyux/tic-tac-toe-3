package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Message struct {
	Command string `json:"command"`
	ID      string `json:"id"`
}

var connectedIds = make(map[int]bool)

func generateId() int {
	for {
		id := rand.Int()
		if connectedIds[id] == false {
			connectedIds[id] = true
			return id
		}
	}
}

func reader(connection *websocket.Conn) {
	for {
		_, messageBytes, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error reading websocket message", err)
			return
		}
		// print out that message for clarity

		fmt.Println(string(messageBytes))

		var v Message
		err = json.Unmarshal(messageBytes, &v)
		if err != nil {
			log.Println("Error parsing JSON", err)
			return
		}

		var response Message

		switch v.Command {
		case "requestId": // send request id
			response = Message{Command: "respondId", ID: strconv.Itoa(generateId())}

		default:
			log.Println("Unknown command:", v.Command)
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Println("Error marshaling response:", err)
			continue
		}
		err = connection.WriteMessage(websocket.TextMessage, responseBytes)
		if err != nil {
			log.Println("Error sending response:", err)
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("client connected")

	reader(ws)
}

func main() {
	fs := http.FileServer(http.Dir("./views"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", wsEndpoint)
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

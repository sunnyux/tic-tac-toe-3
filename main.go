package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Message struct {
	Command     string `json:"command"`
	ID          int    `json:"id,omitempty"`
	P2ID        int    `json:"p2id,omitempty"`
	Coordinates []struct {
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"coordinates,omitempty"`
	Player string `json:"player,omitempty"`
}

var connectedIds = make(map[int]*websocket.Conn)

var connectionList = make(map[int]int)

func generateId(connection *websocket.Conn) int {
	for {
		min := 100
		max := 10000
		id := rand.Intn(max-min) + min
		if connectedIds[id] == nil {
			connectedIds[id] = connection
			return id
		}
	}
}

func joinGame(p1id int, p2id int) bool {
	if connectedIds[p1id] != nil && connectedIds[p2id] != nil {
		connectionList[p1id] = p2id
		connectionList[p2id] = p1id
		return true
	}
	return false
}

func reader(connection *websocket.Conn) {
	for {
		_, messageBytes, err := connection.ReadMessage()
		if err != nil {
			log.Println("Error reading websocket message", err)
			return
		}

		log.Println("Request:", string(messageBytes))

		var v Message
		if err = json.Unmarshal(messageBytes, &v); err != nil {
			log.Println("Error parsing JSON", err)
			return
		}

		var response Message

		switch v.Command {
		case "requestId": // send request id
			response = Message{Command: "respondId", ID: generateId(connection)}

		case "joinGame":
			if p1id, p2id := v.ID, v.P2ID; joinGame(p1id, p2id) {
				response = Message{Command: "successIdFound"}
			} else {
				response = Message{Command: "errorIdNotFound"}
			}

		case "makeMove":
			response = Message{Command: "updateBoardFromMove", Coordinates: v.Coordinates, Player: v.Player}

		default:
			log.Println("Unknown command:", v.Command)
		}

		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Println("Error marshaling response:", err)
			return
		}

		log.Println("Response:", string(responseBytes))

		if err = connection.WriteMessage(websocket.TextMessage, responseBytes); err != nil {
			log.Println("Error sending response:", err)
		}
		if p2id := connectionList[v.ID]; p2id != 0 {
			if err = connectedIds[p2id].WriteMessage(websocket.TextMessage, responseBytes); err != nil {
				log.Println("Error sending response:", err)
			}
			log.Println("Sent response to both", v.ID, p2id)
		} else {
			log.Println("Sent response to only", v.ID)
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
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

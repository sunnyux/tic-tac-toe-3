package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()

		if err != nil {
			break
		}

		fmt.Println("The message is", string(msg), "with a comparison value of", bytes.Compare(msg, []byte("test")))

		if bytes.Compare(msg, []byte("id")) == 0 {
			conn.WriteMessage(t, []byte(fmt.Sprintf("%04d", rand.Intn(10000))))
		} else {
			conn.WriteMessage(t, msg)
		}
	}
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("/home/mattfr/Documents/Programming/Go/src/tic-tac-toe-3/build/react/", true)))

	// Test code for websockets
	router.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	// Start and run the server
	router.Run(":5000")
}

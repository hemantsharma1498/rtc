package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	// max message size allowed
	maxMessageSize = 512
	// I/O read buffer size
	readBufferSize = 1024
	// I/O write buffer size
	writeBufferSize = 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  readBufferSize,
	WriteBufferSize: writeBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		//allow for all origins
		return true
	},
}

func (c *CommunicationServer) UpgradeConn(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(33)
		writeResponse(w, err, "error encountered while upgrading connection", http.StatusInternalServerError)
	}
	fmt.Println(35)
	defer conn.Close()
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received: %s\n", message)

		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func (c *CommunicationServer) Sample(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, nil, "hello", http.StatusInternalServerError)
}

func writeResponse(w http.ResponseWriter, err error, msg any, httpStatus int) error {
	if err != nil {
		log.Printf("Error occured while decoding req json body: %s\n", err)
	}
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(msg)
}

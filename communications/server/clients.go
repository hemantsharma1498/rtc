package server

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Clients struct {
	CMap  map[*websocket.Conn]Client
	mutex *sync.Mutex
}

type Client struct {
	Email        string //email id for the time being
	Name         string
	Organisation string   //org name
	Channels     []string //Channel ids (senderEmail+receiverEmail for now) that a user is involved with
	mutex        *sync.Mutex
	cond         *sync.Cond
}

func (c *Clients) Add(conn *websocket.Conn, client Client) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	client.cond = sync.NewCond(client.mutex)
	if _, exists := c.CMap[conn]; !exists {
		c.CMap[conn] = client
	}
}

func (c *Clients) Remove(conn *websocket.Conn) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if _, exists := c.CMap[conn]; !exists {
		return
	}
	delete(c.CMap, conn)
}

func (c *Clients) UpdateChannels(conn *websocket.Conn, newChannels []string) {
	c.mutex.Lock()
	client, ok := c.CMap[conn]
	c.mutex.Unlock()
	if !ok {
		log.Println("Client not found in map")
		return
	}

	client.mutex.Lock()
	client.Channels = append(client.Channels, newChannels...)
	fmt.Printf("client channels %v for: %s ", client.Email, client.Channels)
	c.CMap[conn] = client
	client.cond.Signal()
	client.mutex.Unlock()
}

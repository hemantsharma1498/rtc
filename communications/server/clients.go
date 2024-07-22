package server

import (
	"github.com/gorilla/websocket"
	"sync"
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
}

func (c *Clients) Add(conn *websocket.Conn, client Client) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
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

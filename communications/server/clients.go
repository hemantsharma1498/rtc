package server

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Clients struct {
	CMap  map[*websocket.Conn]Client
	mutex *sync.Mutex
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

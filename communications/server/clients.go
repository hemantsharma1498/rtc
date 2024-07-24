package server

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Clients struct {
	CMap           map[*websocket.Conn]Client
	mutex          *sync.Mutex
	MessageChannel chan *Message
	NewChannelChan chan *NewChannel
}

type Client struct {
	Email        string //email id for the time being
	Name         string
	Organisation string   //org name
	Channels     []string //Channel ids (senderEmail+receiverEmail hash for now) that a user is involved with
}

type NewChannel struct {
	Conn    *websocket.Conn
	Channel string
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

/*
func (c *Clients) UpdateChannels(conn *websocket.Conn) {
	for {
		select {
		case newChannel := <-c.NewChannelChan:
			c.mutex.Lock()
			client, ok := c.CMap[conn]
			c.mutex.Unlock()
			if !ok {
				log.Println("Client not found in map")
				continue
			}
			client.Channels = append(client.Channels, newChannel.Channel)
			c.mutex.Lock()
			c.CMap[conn] = client
			c.mutex.Unlock()
		}
	}
}
*/

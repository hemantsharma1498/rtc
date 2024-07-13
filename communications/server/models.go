package server

import "github.com/gorilla/websocket"

type Message struct {
	Payload  string `json:"payload"`
	Org      string `json:"org"`
	Sender   string `json:"sender"`   //email of the person who sent the message
	Receiver string `json:"receiver"` //email of the person who received the message
}

type Client struct {
	UserId   int
	Conn     *websocket.Conn
	Channels []int //Channel ids that a user is involved with
}

type ChannelsServed struct{}

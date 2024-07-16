package server

type Message struct {
	Payload  string `json:"payload"`
	Org      string `json:"org"`
	Channel  string `json:"channel"`  //id of the channel where the message is sent
	Name     string `json:"name"`     //name of the channel where the message is sent (null for dms)
	Sender   string `json:"sender"`   //id of the person who sent the message
	Receiver string `json:"receiver"` //id of the person who received the message
}

type Client struct {
	UserId   int
	Channels []int //Channel ids that a user is involved with
}

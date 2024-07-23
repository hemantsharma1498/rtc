package server

type Message struct {
	Payload  string `json:"payload"`
	Org      string `json:"org"`
	Channel  string `json:"channel"`  //id of the channel where the message is sent
	Name     string `json:"name"`     //name of the channel where the message is sent (null for dms)
	Sender   string `json:"sender"`   //id of the person who sent the message
	Receiver string `json:"receiver"` //id of the person who received the message
}

type CreateChannelReq struct {
	Organisation string `json:"Organisation"`
	Sender       string `json:"Sender"`
	Receiver     string `json:"Receiver"`
}

type CreateChannelRes struct {
	ChannelId string `json:"channelId"`
}

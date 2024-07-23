package types

type Message struct {
	Payload    string `json:"payload"`
	OrgId      string `json:"org"`       //id of the org where the message is sent
	ChannelId  string `json:"channel"`   //id of the channel where the message is sent
	Name       string `json:"name"`      //name of the channel where the message is sent (null for dms)
	SenderId   string `json:"sender"`    //id of the person who sent the message
	ReceiverId string `json:"receiver"`  //id of the person who received the message
	CreatedAt  int    `json:"createdAt"` //unix timestamp when the message was created
}

type Client struct {
	Email    string
	Channels []string //Channel ids that a user is involved with
}

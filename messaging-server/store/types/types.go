package types

type Message struct {
	Payload    string `json:"payload"`
	OrgId      string `json:"org"`
	ChannelId  int    `json:"channel"`   //id of the channel where the message is sent
	Name       string `json:"name"`      //name of the channel where the message is sent (null for dms)
	SenderId   int    `json:"sender"`    //id of the person who sent the message
	ReceiverId int    `json:"receiver"`  //id of the person who received the message
	CreatedAt  int    `json:"createdAt"` //unix timestamp when the message was created
}

type Client struct {
	UserId   int
	Channels []int //Channel ids that a user is involved with
}

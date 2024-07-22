package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"strconv"

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

func (c *CommunicationServer) ActiveConnections(w http.ResponseWriter, r *http.Request) {
	orgName := r.PathValue("orgName")
	res := make([]Client, 0)
	for _, v := range c.Clients.CMap {
		if orgName == v.Organisation {
			res = append(res, v)
		}
	}
	writeResponse(w, nil, res, http.StatusOK)
}

func (c *CommunicationServer) ReceiveMsg(conn *websocket.Conn) {
	defer func() {
		c.Clients.mutex.Lock()
		delete(c.Clients.CMap, conn)
		c.Clients.mutex.Unlock()
	}()

	for {
		client := c.Clients.CMap[conn]
		pbsb, err := c.Redis.Receive(client.Channels)
		if err != nil {
			log.Printf("encountered an error while receiving for")
		}

		msg := <-pbsb
		fmt.Println("message: ", msg)
		conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	}

}
func (c *CommunicationServer) UpgradeConn(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			//allow for all origins
			return true
		},
	}
	userEmail := r.URL.Query().Get("email")
	organisation := r.PathValue("org")
	name := r.URL.Query().Get("name")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		writeResponse(w, err, "error encountered while upgrading connection", http.StatusInternalServerError)
	}
	client := Client{Email: userEmail, Name: name, Organisation: organisation, Channels: make([]string, 0)}
	log.Printf("created client: %v\n", client)
	c.Clients.Add(conn, client)
	go c.ReceiveMsg(conn)
}

func (c *CommunicationServer) CreateChannel(w http.ResponseWriter, r *http.Request) {
	fmt.Println(78)
	data := &CreateChannelReq{}
	if err := decodeReqBody(r, data); err != nil {
		log.Printf("encountered an error: %s\n", err)
		writeResponse(w, err, "error encountered while parsing body", http.StatusInternalServerError)
		return
	}
	channel := strconv.Itoa(int(hash(data.Sender + data.Receiver)))
	c.Clients.mutex.Lock()
	defer c.Clients.mutex.Unlock()
	conns := getClientsFromEmail(c.Clients.CMap, []string{data.Receiver, data.Sender})
	if len(conns) == 0 {
		log.Printf("no users found for the given emails: %s, %s\n", data.Sender, data.Receiver)
		writeResponse(w, errors.New("encountered an error"), "error encountered", http.StatusInternalServerError)
		return
	}
	for _, conn := range conns {
		channels := append(c.Clients.CMap[conn].Channels, channel)
		client := c.Clients.CMap[conn]
		fmt.Println("client: ", client)
		client.Channels = channels
		c.Clients.CMap[conn] = client
	}
	var res struct {
		ChannelId string `json:"channelId"`
	}
	res.ChannelId = channel
	writeResponse(w, nil, res, http.StatusCreated)
	return
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
func getClientsFromEmail(cMap map[*websocket.Conn]Client, emails []string) []*websocket.Conn {
	res := make([]*websocket.Conn, 0)
	for k, v := range cMap {
		for _, email := range emails {
			if email == v.Email {
				res = append(res, k)
			}
		}
	}
	return res

}

func decodeReqBody(r *http.Request, d any) error {
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return err
	}
	return nil
}
func writeResponse(w http.ResponseWriter, err error, msg any, httpStatus int) error {
	if err != nil {
		log.Printf("Error occured while decoding req json body: %s\n", err)
	}
	w.WriteHeader(httpStatus)
	return json.NewEncoder(w).Encode(msg)
}

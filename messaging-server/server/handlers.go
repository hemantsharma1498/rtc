package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hemantsharma1498/rtc/messaging/server/types"
)

type CreateChannelReq struct {
	Channel      string `json:"channel"`
	Organisation string `json:"organisationn"`
	Sender       string `json:"senderEmail"`
	Receiver     string `json:"receiverEmail"`
}

const (
	saltSize int    = 16
	time     uint32 = 6
	memory   uint32 = 32
	keyLen   uint32 = 32
)

func (m *Members) SaveMessage(w http.ResponseWriter, r *http.Request) {
	//@TODO add bit to save mesage to db
	data := types.Message{}
	err := decodeReqBody(r, &data)
	fmt.Printf("message received: %v\n", data)
	if err != nil {
		log.Printf("encountered an error while decoding message body: %s\n", err)
		writeResponse(w, nil, nil, http.StatusOK)
		return
	}
	m.Cache.Publish(data)
	writeResponse(w, nil, nil, http.StatusOK)
}

func (m *Members) GetMessages(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, nil, "account created successfully", http.StatusOK)
}

func decodeReqBody(r *http.Request, d any) error {
	// Read the body
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("could not read request body: %v", err)
	}
	defer r.Body.Close()

	// Log the raw body content

	// Reset the request body so it can be read again
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Decode the JSON body
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return fmt.Errorf("could not decode request body: %v", err)
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

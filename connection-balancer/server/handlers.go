package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/hemantsharma1498/rtc/pkg/proto"
)

type ServerAddressResponse struct {
	Address string `json:"Address"`
}

func (c *ConnectionBalancer) GetCommServerAddr(ctx context.Context, in *proto.GetCommServerAddrReq) (*proto.GetCommServerAddrReply, error) {
	if len(c.ServerAddresses) == 0 {
		c.LoadingStatus = false
		c.LoadCommunicationServers()
	}
	var address string = ""
	address = c.ServerAddresses[in.Org]
	if address == "" {
		log.Printf("No address found for %s\n", in.Org)
		return nil, errors.New("error occured")
	}
	return &proto.GetCommServerAddrReply{Address: address}, nil
}

func (c *ConnectionBalancer) GetCommServerAddress(w http.ResponseWriter, r *http.Request) {
	//  org := r.PathValue("org")
	orgArr := strings.Split(r.URL.Path, "/")
	org := orgArr[len(orgArr)-1]
	if len(c.ServerAddresses) == 0 {
		c.LoadingStatus = false
		c.LoadCommunicationServers()
	}
	var address string = ""
	address = c.ServerAddresses[org]
	if address == "" {
		log.Printf("No address found for %s\n", org)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No address found for the given org name"))
		return
	}

	resp := ServerAddressResponse{Address: address}
	log.Printf("Sent address for org: %s\n", resp.Address)
	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Encountered an error while marshalling address for %s\n", org)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Server encountered an error. Please try again."))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

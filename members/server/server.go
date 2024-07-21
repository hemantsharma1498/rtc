package server

import (
	"log"
	"net/http"

	"github.com/hemantsharma1498/rtc/members/pkg/proto"
	"github.com/hemantsharma1498/rtc/members/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Members struct {
	Router           *http.ServeMux
	ConnectionClient proto.ConnectionClient
	store            store.Storage
}

func InitServer(store store.Storage) *Members {
	s := &Members{Router: http.NewServeMux(), store: store}
	s.Routes()
	return s
}

func (m *Members) Start(httpAddr string, grpcAddr string) error {
	log.Printf("Starting members server at address: %s\n", httpAddr)

	// Set up a connection to the server.
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("Instantiated grpc client")
	//@TODO Add graceful connection closingk

	m.ConnectionClient = proto.NewConnectionClient(conn)
	if err := http.ListenAndServe(httpAddr, m.Router); err != nil {
		return err
	}

	return nil
}

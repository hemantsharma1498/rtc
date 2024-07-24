package server

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/hemantsharma1498/rtc/communications/pkg/cache"
	"github.com/hemantsharma1498/rtc/communications/pkg/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CommunicationServer struct {
	Router     *http.ServeMux
	GrpcClient proto.MessagesClient
	Redis      *cache.Cache
	Clients    *Clients
}

func InitServer(cache *cache.Cache) *CommunicationServer {
	return &CommunicationServer{
		Router:  http.NewServeMux(),
		Redis:   cache,
		Clients: &Clients{make(map[*websocket.Conn]Client, 0), &sync.Mutex{}, make(chan *Message), make(chan *NewChannel)},
	}
}

func (c *CommunicationServer) Start(httpAddr string, grpcAddr string) error {
	c.Routes()
	log.Printf("Starting http server at: %s\n", httpAddr)
	// Set up a connection to the server.
	conn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return err
	}
	c.GrpcClient = proto.NewMessagesClient(conn)
	log.Printf("Instantiated grpc client")

	if err := http.ListenAndServe(httpAddr, c.Router); err != nil {
		log.Fatalf("Could not instantiate server: %s\n", err)
	}
	return nil
}

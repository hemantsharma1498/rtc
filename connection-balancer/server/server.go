package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/hemantsharma1498/rtc/pkg/proto"
	"github.com/hemantsharma1498/rtc/store"
	"google.golang.org/grpc"
)

type ConnectionBalancer struct {
	Router          *http.ServeMux
	GrpcServer      *grpc.Server
	LoadingStatus   bool
	ServerAddresses map[string]string
	Store           store.Storage
	proto.ConnectionServer
}

func InitServer(store store.Storage) *ConnectionBalancer {
	s := &ConnectionBalancer{Router: http.NewServeMux(), LoadingStatus: false, ServerAddresses: make(map[string]string, 1), Store: store}
	s.LoadCommunicationServers()
	s.Routes()
	return s
}

func (c *ConnectionBalancer) Start(httpAddr string, grpcAddr int) error {
	log.Printf("Starting http server at: %s\n", httpAddr)

	//Start Grpc Server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcAddr))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		c.GrpcServer = grpc.NewServer()
		proto.RegisterConnectionServer(c.GrpcServer, c)
		log.Printf("Grpc server listening at %v", lis.Addr())
		if err := c.GrpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc: %v", err)
		}
	}()

	//Start Http server
	if err := http.ListenAndServe(httpAddr, c.Router); err != nil {
		log.Fatalf("failed to serve http: %v", err)
		return err
	}
	return nil
}

func (c *ConnectionBalancer) LoadCommunicationServers() {
	servers, err := c.Store.GetAllCommunicationServers()
	if err != nil {
		log.Panicf("error while loading servers: %s\n", err)
	}
	for _, s := range servers {
		c.ServerAddresses[s.Organisation] = s.Address
	}
}

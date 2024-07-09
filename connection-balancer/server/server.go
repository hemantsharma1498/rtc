package server

import (
	"connection-balancer/pkg/proto"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

type ConnectionBalancer struct {
	Router          *http.ServeMux
	GrpcServer      *grpc.Server
	LoadingStatus   bool
	ServerAddresses map[string]string
	Db              *sql.DB
	proto.ConnectionServer
}

func InitServer(d *sql.DB) *ConnectionBalancer {
	s := &ConnectionBalancer{Router: http.NewServeMux(), LoadingStatus: false, ServerAddresses: make(map[string]string, 1), Db: d}
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
	if err := c.getServerAddresses(100, 0); err != nil {
		log.Printf("Encountered an error while fetching server addresses: %s\n", err)
	}
}

func (c *ConnectionBalancer) getServerAddresses(limit int, offset int) error {
	rows, err := c.Db.Query("SELECT org, address FROM communication_servers LIMIT ? OFFSET ?", limit, offset)
	defer rows.Close()
	if err != nil {
		if err == sql.ErrNoRows {
			c.LoadingStatus = true
			return nil
		}
		return err
	}
	foundRows := false
	for rows.Next() {
		foundRows = true
		var (
			org     string
			address string
		)
		if err := rows.Scan(&org, &address); err != nil {
			return err
		}
		if c.ServerAddresses[org] == "" {
			c.ServerAddresses[org] = address
		}
	}

	if !foundRows {
		c.LoadingStatus = true
		return nil
	}

	c.getServerAddresses(limit+100, limit)

	return nil
}

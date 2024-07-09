package main

import (
	"connection-balancer/server"
	"connection-balancer/store"
	"context"
	"log"
)

const httpAddress = ":8080"
const grpcAddress = 8081

func main() {
	log.Printf("Initialising connection balancer")
	log.Printf("Connecting to database...")
	ctx := context.Background()
	store, err := store.NewConnBalConnector().Connect(ctx)
	if err != nil {
		log.Panicf("Unable to connect to db, error: %s\n", err)
	}
	log.Printf("Db connection established")
	s := server.InitServer(store.Db)

	if err = s.Start(httpAddress, grpcAddress); err != nil {
		log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
	}

}

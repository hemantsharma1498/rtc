package main

import (
	"context"
	"log"

	"github.com/hemantsharma1498/rtc/server"
	"github.com/hemantsharma1498/rtc/store"
)

const httpAddress = ":3010"
const grpcAddress = 3011

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

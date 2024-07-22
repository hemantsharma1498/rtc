package main

import (
	"github.com/hemantsharma1498/rtc/messaging/server"
	"github.com/hemantsharma1498/rtc/messaging/store/mysqlDb"
	"log"
)

const httpAddress = ":3030"
const grpcAddress = ":3031"

func main() {
	log.Printf("Initialising members server")

	log.Printf("Connecting to database...")
	store, err := mysqlDb.NewMessagingDbConnector().Connect()
	if err != nil {
		log.Panicf("Unable to connect to db, error: %s\n", err)
	}
	log.Printf("Db connection established")

	s := server.InitServer(store)
	if err = s.Start(httpAddress, grpcAddress); err != nil {
		log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
	}
}

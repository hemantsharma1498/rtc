package main

import (
	"log"

	"github.com/hemantsharma1498/rtc/server"
	"github.com/hemantsharma1498/rtc/store/mysqlDb"
)

const httpAddress = ":3020"
const grpcAddress = ":8081"

func main() {
	log.Printf("Initialising members server")

	log.Printf("Connecting to database...")
	store, err := mysqlDb.NewMembersDbConnector().Connect()
	if err != nil {
		log.Panicf("Unable to connect to db, error: %s\n", err)
	}
	log.Printf("Db connection established")

	s := server.InitServer(store)
	if err = s.Start(httpAddress, grpcAddress); err != nil {
		log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
	}
}

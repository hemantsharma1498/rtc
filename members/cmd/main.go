package main

import (
	"log"
	"members/server"
	"members/store/mysqlDb"
)

const httpAddress = ":3000"
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

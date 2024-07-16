package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hemantsharma1498/rtc/server"
	"github.com/hemantsharma1498/rtc/store"
	"github.com/hemantsharma1498/rtc/store/mysqlDb"
	"github.com/hemantsharma1498/rtc/store/types"
)

const httpAddress = ":3030"
const grpcAddress = ":3031"

func main() {
	log.Printf("Initialising members server")

	log.Printf("Connecting to database...")
	store, err := mysqlDb.NewMembersDbConnector().Connect()
	if err != nil {
		log.Panicf("Unable to connect to db, error: %s\n", err)
	}
	log.Printf("Db connection established")

	s := server.InitServer(store)
	go func() {
		if err = s.Start(httpAddress, grpcAddress); err != nil {
			log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
		}
	}()
	fmt.Println("testing everything")
	testEverything(store)
}

func testEverything(store store.Storage) {
	if err := store.SaveMessage(&types.Message{Payload: "Test message number 1", ChannelId: 1, SenderId: 1, ReceiverId: 1, CreatedAt: int(time.Now().Unix())}); err != nil {
		log.Printf("save msg err: %s\n", err)
	}

	msgs, err := store.GetMessages([]int{1})
	if err != nil {
		log.Printf("get msg err: %s\n", err)
	}
	log.Printf("get messages %v\n", msgs)
}

package main

import (
	"communications/pkg/cache"
	"communications/server"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	CacheAddr     = "localhost:6379"
	CachePassword = ""
)

const httpAddress = ":3010"
const grpcAddress = ":9091"

func main() {
	log.Printf("Initialising members server")

	cache, err := cache.NewCache().Start(CacheAddr, CachePassword)
	if err != nil {
		log.Fatalf("Unable to start redis %s\n", err)
	}
	go receive(cache)
	publish(cache)
	s := server.InitServer(cache)
	if err = s.Start(httpAddress, grpcAddress); err != nil {
		log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
	}

}

func publish(cache *cache.Cache) {
	log.Println("starting to publish")
	time.Sleep(time.Second * 3)
	counter := 0
	for {
		if counter == 10 {
			break
		}
		msg := "hello" + strconv.Itoa(counter)
		fmt.Println(msg)
		cache.Publish("test_channel", msg)
		counter++
	}

}
func receive(cache *cache.Cache) {
	log.Println("starting to receive")
	fmt.Println(54)
	err := cache.Subscribe([]string{"test_channel"})
	if err != nil {
		log.Fatalf("Unable to start redis %s\n", err)
	}
	fmt.Println(59)
	msgs, err := cache.Receive("test_channel")
	if err != nil {
		log.Fatalf("Unable to start redis %s\n", err)
	}
	fmt.Println(64)
	for {
		msg := <-msgs
		if msg.Payload == "hello3" {
			break
		}
		fmt.Println(msg)
	}

}

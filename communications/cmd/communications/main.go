package main

import (
	"github.com/hemantsharma1498/rtc/pkg/cache"
	"github.com/hemantsharma1498/rtc/server"
	"log"
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
	s := server.InitServer(cache)
	if err = s.Start(httpAddress, grpcAddress); err != nil {
		log.Panicf("Failed to initialise server at %s, error: %s\n", httpAddress, err)
	}

}

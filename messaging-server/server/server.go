package server

import (
	"log"
	"net/http"

	"github.com/hemantsharma1498/rtc/messaging/pkg/cache"
	"github.com/hemantsharma1498/rtc/messaging/server/types"
	"github.com/hemantsharma1498/rtc/messaging/store"
)

type Members struct {
	Router  *http.ServeMux
	store   store.Storage
	Cache   *cache.Cache
	Clients []*types.Client
}

func InitServer(store store.Storage) *Members {
	s := &Members{Router: http.NewServeMux(), store: store, Cache: cache.NewCache()}
	s.Routes()
	return s
}

func (m *Members) Start(httpAddr string, grpcAddr string) error {
	log.Printf("Starting members server at address: %s\n", httpAddr)
	if err := http.ListenAndServe(httpAddr, m.Router); err != nil {
		return err
	}
	return nil
}

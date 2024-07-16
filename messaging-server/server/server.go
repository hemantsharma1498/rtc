package server

import (
	"log"
	"net/http"

	"github.com/hemantsharma1498/rtc/store"
)

type Members struct {
	Router *http.ServeMux
	store  store.Storage
}

func InitServer(store store.Storage) *Members {
	s := &Members{Router: http.NewServeMux(), store: store}
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

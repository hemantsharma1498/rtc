package server

import (
	"net/http"
)

func (m *CommunicationServer) Routes() {
	m.Router.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		m.UpgradeConn(w, r)
	})
	m.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		m.Sample(w, r)
	})
}

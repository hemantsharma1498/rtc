package server

import (
	"net/http"
)

func (m *CommunicationServer) Routes() {
	m.Router.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		m.UpgradeConn(w, r)
	})
	m.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		m.Sample(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

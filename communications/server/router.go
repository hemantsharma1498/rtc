package server

import (
	"net/http"
)

func (m *CommunicationServer) Routes() {
	m.Router.HandleFunc("/socket/{org}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		m.UpgradeConn(w, r)
	})
	m.Router.HandleFunc("/active-connections/{orgName}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		m.ActiveConnections(w, r)
	})
	m.Router.HandleFunc("/create-channel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		}
		enableCors(&w)
		m.CreateChannel(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

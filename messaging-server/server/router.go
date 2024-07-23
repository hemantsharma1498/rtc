package server

import (
	"net/http"
)

func (m *Members) Routes() {
	m.Router.HandleFunc("/save-message", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		m.SaveMessage(w, r)
	})
	m.Router.HandleFunc("/get-chat-history", func(w http.ResponseWriter, r *http.Request) {
		m.GetMessages(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

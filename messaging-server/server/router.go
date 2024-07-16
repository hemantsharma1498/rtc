package server

import (
	"net/http"
)

func (m *Members) Routes() {
	m.Router.HandleFunc("/save-message", func(w http.ResponseWriter, r *http.Request) {
		m.SaveMessage(w, r)
	})
	m.Router.HandleFunc("/get-chat-history", func(w http.ResponseWriter, r *http.Request) {
		m.GetMessages(w, r)
	})
}

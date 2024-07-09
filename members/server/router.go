package server

import (
	"net/http"
)

func (m *Members) Routes() {
	m.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		m.Login(w, r)
	})
	m.Router.HandleFunc("/get-connection", func(w http.ResponseWriter, r *http.Request) {
		m.GetCommServerAddress(w, r)
	})
}

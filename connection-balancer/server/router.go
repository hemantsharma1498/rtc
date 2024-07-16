package server

import (
	"net/http"
)

func (c *ConnectionBalancer) Routes() {
	c.Router.HandleFunc("/get-cserver-addresses/{org}", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		c.GetCommServerAddress(w, r)
	})
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

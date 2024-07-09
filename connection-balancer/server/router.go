package server

import (
	"net/http"
)

func (c *ConnectionBalancer) Routes() {
	c.Router.HandleFunc("/get-cserver-addresses/{org}", func(w http.ResponseWriter, r *http.Request) {
		c.GetCommServerAddress(w, r)
	})
}

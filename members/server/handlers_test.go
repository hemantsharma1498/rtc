package server

import (
	"connection-balancer/store"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCommServerAddress(t *testing.T) {
	ctx := context.TODO()
	db, _ := store.NewConnBalConnector().Connect(ctx)
	cb := &ConnectionBalancer{
		Router:          http.NewServeMux(),
		LoadingStatus:   false,
		serverAddresses: make(map[string]string),
	}
	cb.Routes(db.Db)
	tests := []struct {
		name           string
		org            string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Error | Org not found in map",
			org:            "Org_399",
			expectedStatus: 400,
			expectedBody:   "No address found for the given org name",
		},
		{
			name:           "Success | Address returned for given org",
			org:            "Org_99",
			expectedStatus: 200,
			expectedBody:   `{"Address":"Address_99"}`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/get-cserver-addresses/"+tc.org, nil)
			if err != nil {
				t.Fatal(err)
			}
			ctx := context.TODO()
			req = req.WithContext(ctx)
			rr := httptest.NewRecorder()
			cb.GetCommServerAddress(rr, req, db.Db)

			if status := rr.Code; status != tc.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.expectedStatus)
			}
			if strings.TrimSpace(rr.Body.String()) != tc.expectedBody {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tc.expectedBody)
			}
		})

	}
	//req, err := http.NewRequest("GET", "/get-cserver-addresses/", nil)
}

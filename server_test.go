/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */
package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/gorilla/websocket"
	"strings"
)
func TestServer(t *testing.T) {
	t.Run("health check at GET /health returns 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		health(response, request)

		got := response.Result().StatusCode
		want :=  200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("accepts an upgraded connection to GET /", func (t *testing.T)  {
		server := httptest.NewServer(http.HandlerFunc(connectionHandler))
    	defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")

		websocket, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer websocket.Close()
	})

	t.Run("only accepts connections to the ocpp16 sub protocol", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(connectionHandler))
    	defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")
		websocket.DefaultDialer.Subprotocols = append(websocket.DefaultDialer.Subprotocols, "ocpp16")
		
		webserver, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil { t.Fatalf("%v", err) }

		got := webserver.Subprotocol()
		want := "ocpp16"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		defer webserver.Close()
	})

	t.Run("saves connection when connected", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(connectionHandler))
    	defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")
		
		webserver, _, err := websocket.DefaultDialer.Dial(url + "/charge-box-id", nil)
		if err != nil { t.Fatalf("%v", err) }

		got := getConnected("charge-box-id")
		want := true

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
		defer webserver.Close()
	})
}
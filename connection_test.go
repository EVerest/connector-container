/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

 package main

 import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
 )
 func TestWebsocket(t *testing.T) {
	 
	 t.Run("accepts an upgraded connection to GET", func (t *testing.T)  {
		server := httptest.NewServer(http.HandlerFunc(connectionHandler))
    	defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")

		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer ws.Close()
	})

	t.Run("only accepts connections to the ocpp1.6 sub protocol", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(connectionHandler))
    	defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")
		websocket.DefaultDialer.Subprotocols = append(websocket.DefaultDialer.Subprotocols, "ocpp1.6")
		
		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil { t.Fatalf("%v", err) }

		got := ws.Subprotocol()
		want := "ocpp1.6"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		defer ws.Close()
	})

	t.Run("saves connection when connected", func(t *testing.T) {
		storage := make(localStore)
		StartConnectionHandler(&storage)
		
		srv := httptest.NewServer(http.HandlerFunc(connectionHandler))
    	defer srv.Close()

		url := "ws" + strings.TrimPrefix(srv.URL, "http")
		
		client, _, err := websocket.DefaultDialer.Dial(url + "/charge-box-id", nil)
		if err != nil { t.Fatalf("%v", err) }

		got := getConnection("charge-box-id")

		if got == nil {
			t.Errorf("got no connector")
		}
		defer client.Close()
	})
}
/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package connection

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestConnection(t *testing.T) {

	t.Run ("accepts an upgraded connection to GET", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")

		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		defer ws.Close()
	})

	t.Run("enforces ocpp1.6 sub protocol", func(t *testing.T) {
		options := ConnectionOptions {
			SubProtocol: "ocpp1.6",
		}
		NewConnectionHandler(options)
		
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")
		websocket.DefaultDialer.Subprotocols = append(websocket.DefaultDialer.Subprotocols, "ocpp1.6")

		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}

		got := ws.Subprotocol()
		want := "ocpp1.6"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		defer ws.Close()
	})

	t.Run("enforces any sub protocol", func(t *testing.T) {
		options := ConnectionOptions {
			SubProtocol: "any",
		}
		NewConnectionHandler(options)
		
		server := httptest.NewServer(http.HandlerFunc(handler))
		defer server.Close()

		url := "ws" + strings.TrimPrefix(server.URL, "http")
		websocket.DefaultDialer.Subprotocols = append(websocket.DefaultDialer.Subprotocols, "any")

		ws, _, err := websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			t.Fatalf("%v", err)
		}

		got := ws.Subprotocol()
		want := "any"

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
		defer ws.Close()
	})

	t.Run("saves connection when connected", func(t *testing.T) {
		storage := &storeIt{}
		doer := doIt{}
		options := ConnectionOptions {
			SubProtocol: "ocpp1.6",
			ConnectionStore: storage,
			Doer: &doer,
		}
		NewConnectionHandler(options)

		srv := httptest.NewServer(http.HandlerFunc(handler))
		defer srv.Close()

		url := "ws" + strings.TrimPrefix(srv.URL, "http")

		client, _, err := websocket.DefaultDialer.Dial(url+"/charge-box-id", nil)
		defer client.Close()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if !storage.put {
			t.Fatal("Put not called")
		}
	})

	t.Run("a connection starts doer.ConnectionWriter", func(t *testing.T) {
		storage := &storeIt{}
		doer := doIt{}
		options := ConnectionOptions {
			SubProtocol: "ocpp1.6",
			ConnectionStore: storage,
			Doer: &doer,
		}
		NewConnectionHandler(options)

		srv := httptest.NewServer(http.HandlerFunc(handler))
		defer srv.Close()

		url := "ws" + strings.TrimPrefix(srv.URL, "http")

		client, _, err := websocket.DefaultDialer.Dial(url+"/a-charge-box-id", nil)
		defer client.Close()
		if err != nil {
			t.Fatalf("%v", err)
		}

		if doer.connectionWriter == false {
			t.Errorf("ConnectionWriter not called")
		}
	})

	t.Run("a connection starts doer.ConnectionReader", func(t *testing.T) {
		storage := &storeIt{}
		doer := &doIt{}
		options := ConnectionOptions {
			SubProtocol: "ocpp1.6",
			ConnectionStore: storage,
			Doer: doer,
		}
		NewConnectionHandler(options)

		srv := httptest.NewServer(http.HandlerFunc(handler))
		defer srv.Close()

		url := "ws" + strings.TrimPrefix(srv.URL, "http")

		client, _, err := websocket.DefaultDialer.Dial(url+"/a-charge-box-id", nil)
		
		if err != nil {
			t.Fatalf("%v", err)
		}

		client.WriteMessage(1, []byte(`a-payload`))

		if doer.connectionReader == false {
			t.Errorf("ConnectionReader not called")
		}
	})
}

type doIt struct {
	connectionWriter bool
	connectionReader bool
}

func (d *doIt) ConnectionReader(URIpath string, b []byte) {
	log.Printf("ConnectionReader: %s", b)
	d.connectionReader = true			
}

func (d *doIt) ConnectionWriter() (URIpath string, payload []byte){
	// log.Println("ConnectionWriter")
	d.connectionWriter = true
	return "", nil
}

type storeIt struct {
	put bool
	get bool
	delete bool
}

func (s *storeIt) Put(string, *websocket.Conn) bool {
	s.put = true
	return true
}

func (s *storeIt) Get(string) *websocket.Conn {
	s.get = true
	return nil
}

func (s *storeIt) Delete(string) bool {
	s.delete = true
	return true
}
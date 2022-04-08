/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: Apache-2.0
 */

package connection

import (
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
)

type connectionStore interface {
	Put(string, *websocket.Conn) bool
	Get(string) *websocket.Conn
	Delete(string) bool
}

type converter interface {
	ConnectionReader(URIpath string, b []byte)
	ConnectionWriter() (URIpath string, payload []byte)
	ConnectEvent(URIpath string)
	DisconnectEvent(URIpath string)
}

type ConnectionHandler struct {
	SubProtocol     string
	ConnectionStore connectionStore
	Converter       converter
}

const forwardSlash = "/"

func (ch *ConnectionHandler) Handler(writer http.ResponseWriter, request *http.Request) {
	conn := upgrade(writer, request, ch.SubProtocol)
	defer conn.Close()

	URIpath := getURIpath(*request)
	if URIpath == "" {
		log.Println("URI path is required as a reference for storing websocket connections.")
		conn.Close()
		return
	}
	ch.ConnectionStore.Put(URIpath, conn)
	ch.Converter.ConnectEvent(URIpath)

	var wg sync.WaitGroup

	wg.Add(2)

	go connectionWriter(&wg, ch.ConnectionStore, ch.Converter)
	go connectionReader(URIpath, conn, &wg, ch.Converter)

	wg.Wait()
}

func connectionReader(URIpath string, conn *websocket.Conn, wg *sync.WaitGroup, do converter) {
	defer wg.Done()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			do.DisconnectEvent(URIpath)
			break
		}
		do.ConnectionReader(URIpath, message)
	}
}

func connectionWriter(wg *sync.WaitGroup, cs connectionStore, do converter) {
	defer wg.Done()
	for {
		path, payload := do.ConnectionWriter()

		conn := cs.Get(path)
		if conn == nil {
			log.Printf("no connection for path %s", path)
			break
		}
		conn.WriteMessage(1, payload)
	}
}

func getURIpath(request http.Request) string {
	return strings.TrimPrefix(request.RequestURI, forwardSlash)
}

func upgrade(writer http.ResponseWriter, request *http.Request, subProtocol string) *websocket.Conn {
	upgrader := websocket.Upgrader{Subprotocols: []string{subProtocol}}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, _ := upgrader.Upgrade(writer, request, nil)

	return conn
}

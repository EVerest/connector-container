/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
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
	Connect(URIpath string)
	Disconnect(URIpath string)
}

type ConnectionOptions struct {
	SubProtocol 		string
	ConnectionStore 	connectionStore
	Converter			converter 
}

var cs connectionStore
var subProtocol string
var do converter
const forwardSlash = "/"

func NewConnectionHandler(connectionOptions ConnectionOptions) func(http.ResponseWriter, *http.Request) {
	cs = connectionOptions.ConnectionStore
	subProtocol = connectionOptions.SubProtocol
	do = connectionOptions.Converter

	return handler
}

func handler(writer http.ResponseWriter, request *http.Request) {
	conn := upgrade(writer, request)
	defer conn.Close()

	URIpath := getURIpath(*request)
	if URIpath == "" {
		log.Println("URI path is required as a reference for storing websocket connections.")
		conn.Close()
		return
	}
	saveConnection(URIpath, conn)
	do.Connect(URIpath)

	var wg sync.WaitGroup

	wg.Add(2)

	go connectionWriter(&wg)
	go connectionReader(URIpath, conn, &wg)
	
	wg.Wait()
}

func connectionReader(URIpath string, conn *websocket.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {	
		_, message, err := conn.ReadMessage()
		if err != nil { 
			do.Disconnect(URIpath)
			break
		}
		do.ConnectionReader(URIpath, message)
	}
}

func connectionWriter(wg *sync.WaitGroup) {
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

func saveConnection(URIpath string, conn *websocket.Conn) bool {
	return cs.Put(URIpath,conn)
}

func getConnection(URIpath string) *websocket.Conn {
	return cs.Get(URIpath)
}

func getURIpath(request http.Request) string {
	return strings.TrimPrefix(request.RequestURI, forwardSlash)
}

func upgrade(writer http.ResponseWriter, request *http.Request) *websocket.Conn {
	upgrader := websocket.Upgrader{Subprotocols: []string{subProtocol}}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, _ := upgrader.Upgrade(writer, request, nil)
	// if error log debug?
	return conn
}
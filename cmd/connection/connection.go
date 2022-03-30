/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package connection

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type connectionStore interface {
	Put(string, *websocket.Conn) bool
	Get(string) *websocket.Conn
	Delete(string) bool
}

type doer interface {
	FromConnection(URIpath string, b []byte) interface{}
	ToConnection(o interface{}) []byte
}

var cs connectionStore
var subProtocol string
var do doer
const forwardSlash = "/"

type ConnectionOptions struct {
	subProtocol 		string
	connectionStore 	connectionStore
	doer				doer 
}

func NewConnectionHandler(connectionOptions ConnectionOptions) func(http.ResponseWriter, *http.Request) {
	cs = connectionOptions.connectionStore
	subProtocol = connectionOptions.subProtocol
	do = connectionOptions.doer

	return handler
}

func handler(writer http.ResponseWriter, request *http.Request) {
	conn := upgrade(writer, request)
	defer conn.Close()

	URIpath := getURIpath(*request)
	cs.Put(URIpath, conn)

	go connectionReader(URIpath, conn)
	go connectionWriter()
}

func connectionWriter() {
	log.Println("Write")
	// cs.Get("cbid")
	// Call writer
	// Get connection by URIpath
	// Send payload
	// for {
	// 	conn.WriteMessage(1, do.ToConnection())
	// }
}

func connectionReader(URIpath string, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		do.FromConnection(URIpath, message)
	}
}

func getConnection(chargeBoxId string) *websocket.Conn {
	return cs.Get(chargeBoxId)
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

/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

type connectionStore interface {
	Put(string, *websocket.Conn) bool
	Get(string) *websocket.Conn
	Delete(string) bool
}

var cs connectionStore
var subProtocol string

type ConnectionOptions struct {
	subProtocol 	string
	connectionStore connectionStore
}

func NewConnectionHandler(connectionOptions ConnectionOptions) func(http.ResponseWriter, *http.Request) {
	cs = connectionOptions.connectionStore
	subProtocol = connectionOptions.subProtocol

	return handler
}

func handler(writer http.ResponseWriter, request *http.Request) {
	connection := upgrade(writer, request)
	defer connection.Close()

	chargeBoxId := getChargeBoxID(*request)
	cs.Put(chargeBoxId, connection)

	// Maybe a channel here that's consumed by convert?
	evseReader(chargeBoxId, connection)
}

func GetConnection(chargeBoxId string) *websocket.Conn {
	return cs.Get(chargeBoxId)
}

func getChargeBoxID(request http.Request) string {
	return strings.TrimPrefix(request.RequestURI, "/")
}

func upgrade(writer http.ResponseWriter, request *http.Request) *websocket.Conn {
	upgrader := websocket.Upgrader{Subprotocols: []string{subProtocol}}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, _ := upgrader.Upgrade(writer, request, nil)
	// if error log debug?
	return conn
}

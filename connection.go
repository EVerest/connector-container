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

var ls = make(localStore)

const subProtocol string = "ocpp1.6"

func connectionHandler(writer http.ResponseWriter, request *http.Request) {
	connection := upgrade(writer, request)
	defer connection.Close()

	chargeBoxId := getChargeBoxId(*request)
	connectionStore.Put(ls, chargeBoxId, connection)

	evseMessenger(chargeBoxId, connection)
}

func getConnection(chargeBoxId string) *websocket.Conn {
	return ls.Get(chargeBoxId)
}

func getChargeBoxId(request http.Request) string {
	return strings.TrimPrefix(request.RequestURI, "/")
}

func upgrade(writer http.ResponseWriter, request *http.Request) *websocket.Conn {
	upgrader := websocket.Upgrader{ Subprotocols: []string{subProtocol} }
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, _ := upgrader.Upgrade(writer, request, nil)
	// if error log debug?
	return conn
}
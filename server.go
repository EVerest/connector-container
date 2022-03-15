/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
)

var store = make(map[string]*websocket.Conn)

var upgrader = websocket.Upgrader{
	Subprotocols: []string{"ocpp16"},
}

func main() {
	http.HandleFunc("/", connectionHandler)
	http.HandleFunc("/health", health)
}

func health(writer http.ResponseWriter, request *http.Request) { /*noop*/ }

func connectionHandler(writer http.ResponseWriter, request *http.Request) {
	websocket, error := upgrader.Upgrade(writer, request, nil)
    if error != nil {
        log.Printf("Upgrading error: %#v\n", error)
        return
    }
	store["charge-box-id"] = websocket
    defer websocket.Close()
}

func getConnection(chargeBoxId string) *websocket.Conn {
	return store[chargeBoxId]
}
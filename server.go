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


var upgrader = websocket.Upgrader{
	Subprotocols: []string{"ocpp16"},
}

func main() {
	http.HandleFunc("/", connectionHandler)
	http.HandleFunc("/health", health)
}
func health(writer http.ResponseWriter, request *http.Request) { /*noop*/ }

func connectionHandler(writer http.ResponseWriter, request *http.Request) {
	headers := make(http.Header)
	c, err := upgrader.Upgrade(writer, request, headers)
    if err != nil {
        log.Printf("Upgrading error: %#v\n", err)
        return
    }
    defer c.Close()
}
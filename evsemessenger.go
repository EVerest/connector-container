/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import (
	"log"
	"github.com/gorilla/websocket"
)

func evseMessenger(chargeBoxId string, conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		ocppcc := fromByteSlice(chargeBoxId, message)
		log.Printf("ocppcc: %v", ocppcc)	
	}	
}
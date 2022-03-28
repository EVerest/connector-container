/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import "github.com/gorilla/websocket"

type localStore map[string]*websocket.Conn

func (ls localStore) Put(chargeBoxId string, conn *websocket.Conn) bool {
	ls[chargeBoxId] = conn
	return true
}

func (ls localStore) Get(chargeBoxId string) *websocket.Conn {
	return ls[chargeBoxId]
}

func (ls localStore) Delete(chargeBoxId string) bool {
	delete(ls, chargeBoxId)
	return true
}

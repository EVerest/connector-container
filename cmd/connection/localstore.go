/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package connection

import "github.com/gorilla/websocket"

type localStore map[string]*websocket.Conn

func (ls localStore) Put(key string, conn *websocket.Conn) bool {
	ls[key] = conn
	return true
}

func (ls localStore) Get(key string) *websocket.Conn {
	return ls[key]
}

func (ls localStore) Delete(key string) bool {
	delete(ls, key)
	return true
}

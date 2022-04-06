/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package store

import "github.com/gorilla/websocket"

type LocalStore map[string]*websocket.Conn

func (ls LocalStore) Put(key string, conn *websocket.Conn) bool {
	ls[key] = conn
	return true
}

func (ls LocalStore) Get(key string) *websocket.Conn {
	return ls[key]
}

func (ls LocalStore) Delete(key string) bool {
	delete(ls, key)
	return true
}
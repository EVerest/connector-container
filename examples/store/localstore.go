/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: Apache-2.0
 */

package store

import (
	"log"

	"github.com/gorilla/websocket"
)

type LocalStore map[string]*websocket.Conn

func (ls LocalStore) Put(key string, conn *websocket.Conn) bool {
	ls[key] = conn
	log.Printf("Put: localstore has %d items", len(ls))
	return true
}

func (ls LocalStore) Get(key string) *websocket.Conn {
	log.Printf("Get: localstore has %d items", len(ls))
	return ls[key]
}

func (ls LocalStore) Delete(key string) bool {
	delete(ls, key)
	log.Printf("Delete: localstore has %d items", len(ls))
	return true
}

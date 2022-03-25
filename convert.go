/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type ocppcc struct {
    timestamp       uint32    
    messageTypeId   uint16
    chargeBoxId     string
    messageId       string
    action          string
    payload         map[string]interface {}
}

func toOcppByteSlice(o ocppcc) []byte {
    var arr = make([]interface{}, 4)
    arr[0] = o.messageTypeId
    arr[1] = o.messageId
    arr[2] = o.action
    arr[3] = o.payload

    res, err := json.Marshal(arr)
    if err != nil {
        fmt.Println(err)
    }
    return res
}

func fromOcppByteSlice(chargeBoxId string, b []byte) ocppcc {
    ocppcc := ocppcc{}
    var arr []interface{}

    err := json.Unmarshal(b, &arr)
    if err != nil {
        fmt.Println(err)
    }

    ocppcc.action       = arr[2].(string)
    ocppcc.messageId    = arr[1].(string)
    ocppcc.timestamp    = uint32(time.Now().UnixMilli())
    ocppcc.payload      = arr[3].(map[string]interface {})
    ocppcc.chargeBoxId  = chargeBoxId

    return ocppcc
}

type Reader struct {
    data []byte
    readIndex int64
}

// Start function that creates a reader for evseReader to write to?

// Called by io.Read impl?
func evseReader(chargeBoxId string, conn *websocket.Conn) {
    // Writes to Reader struct
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		ocppcc := fromOcppByteSlice(chargeBoxId, message)
		log.Printf("ocppcc: %v", ocppcc)	
	}		
}

func (r *Reader) Read(p []byte) (n int, err error) {
    if r.readIndex >= int64(len(r.data)) {
        err = io.EOF
        return
    }

    n = copy(p, r.data[r.readIndex:])
    r.readIndex += int64(n)
    return
}

func (o *ocppcc) Write(data []byte) (n int, err error) {
    // for 
    // data contains JSON encoded ocppcc message
    // ocppcc = decoded JSON []byte
    // pull chargeBoxId, messageId from ocppcc
    // convert ocppcc to array of interface byte array
    // get connection from store
    // send message 
    // return len data, nil

    // feels dirty
    getConnection("charge-box-id from data []byte")
    return 
}

// Called by io.Write impl?
func evseWriter(o ocppcc) {
    
}
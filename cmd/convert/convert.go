/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package convert

import (
	"encoding/json"
	// "fmt"
	"io"
	"log"
	"time"
)

type OCPPCC struct {
	Timestamp     uint32 					`json:"timestamp"`
	MessageTypeID float64 					`json:"messageTypeId"`
	ChargeBoxID   string 					`json:"chargeBoxId"`
	MessageID     string 					`json:"messageId"`
	Action        string 					`json:"action"`
	Payload       map[string]interface{}	`json:"payload"`
}

func (r *EVSEreader) ConnectionReader(chargeBoxID string, b []byte) {
	ocppcc := OCPPCC{}
	var arr []interface{}

	err := json.Unmarshal(b, &arr)
	if err != nil {
		log.Println(err)
	}

	ocppcc.Action 		= arr[2].(string)
	ocppcc.MessageID 	= arr[1].(string)
	ocppcc.Timestamp 	= uint32(time.Now().UnixMilli())
	ocppcc.Payload 		= arr[3].(map[string]interface{})
	ocppcc.ChargeBoxID 	= chargeBoxID
	ocppcc.MessageTypeID= arr[0].(float64)

	bytes, err := json.Marshal(ocppcc)
	if err != nil {
		log.Println(err)
	}
	
	r.data = append(r.data, bytes...)
}

func NewEVSEreader() *EVSEreader {
	r = &EVSEreader{
		data: []byte{},
		readIndex: 0,
	}
	
	return r
}

type EVSEreader struct {
	data      	[]byte
	readIndex 	int64
}

var r *EVSEreader

func (r *EVSEreader) Read(p []byte) (n int, err error) {
	if r.readIndex >= int64(len(r.data)) {
		err = io.EOF
		return
	}
	n = copy(p, r.data[r.readIndex:])
	r.readIndex += int64(n)

	return
}

var w *EVSEwriter

func NewEVSEWriter() *EVSEwriter {
	w = &EVSEwriter{
		data: []byte{},
	}

	return w
}

type EVSEwriter struct {
	data []byte
}

func (w *EVSEwriter) Write(data []byte) (n int, err error) {
	for _, b := range data {
		w.data = append(w.data, b)
	}
	log.Printf("Write: %s", w.data)
	w.data = data
	return len(data), nil
	// n = 0
	// for _, b := range data {
	// 	append(w.data, b)
	// 	n++
	// }
	// return n, nil
	// for
	// data contains JSON encoded ocppcc message
	// ocppcc = decoded JSON []byte
	// pull chargeBoxId, messageId from ocppcc
	// convert ocppcc to array of interface byte array
	// get connection from store
	// send message
	// return len data, nil

	// feels dirty
}

func (w *EVSEwriter) Close() error {
	w.Close()

	return nil
}

// func ConnectionWriter(o OCPPCC) []byte {
// 	var arr = make([]interface{}, 4)
// 	arr[0] = o.messageTypeID
// 	arr[1] = o.messageID
// 	arr[2] = o.action
// 	arr[3] = o.payload

// 	res, err := json.Marshal(arr)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	return res
// }
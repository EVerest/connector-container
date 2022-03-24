/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ocppReader struct {
    src []byte
    cur int
}

type ocppcc struct {
    timestamp       uint32    
    messageTypeId   string
    chargeBoxId     string
    messageId       string
    action          string
    payload         map[string]interface {}
}

func fromByteSlice(chargeBoxId string, b []byte) ocppcc {
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

func (o *ocppReader) Read(p []byte) (n int, e error) {
    fmt.Printf("Action: %v", o.src)
    y, err := json.Marshal(o)
    if err != nil {
        fmt.Println(err)
    }
    p = y
	return	
}
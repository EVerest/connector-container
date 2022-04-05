/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package convert

import (
	"encoding/json"
	"log"
	"time"
)

type OCPPCC struct {
	Timestamp     uint32 					`json:"timestamp"`
	MessageTypeID uint8					`json:"messageTypeId"`
	ChargeBoxID   string 					`json:"chargeBoxId"`
	MessageID     string 					`json:"messageId"`
	Action        string 					`json:"action"`
	Payload       map[string]interface{}	`json:"payload"`
}

func (eData *EVSEdata) ConnectionReader(URIpath string, b []byte) {	
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
	ocppcc.ChargeBoxID 	= URIpath
	ocppcc.MessageTypeID= uint8(arr[0].(float64))

	eData.rDataCh <- ocppcc
}

type EVSEdata struct {
	rDataCh chan OCPPCC
	wData []byte
	wDataCh chan byte
}

var eData EVSEdata

func NewEVSEdata() *EVSEdata {
	eData = EVSEdata{}
	eData.rDataCh = make(chan OCPPCC, 100)
	eData.wDataCh = make(chan byte)

	eData.wData = []byte("")
	
	return &eData
}

func (eData *EVSEdata) Read(b []byte) (int, error) {
	ocppcc :=  <-eData.rDataCh

	bytes, err := json.Marshal(ocppcc)
	if err != nil {
		log.Println(err)
	}

	n := copy(b, bytes)
    return n, nil
}

func (eData *EVSEdata) ConnectionWriter() (URIpath string, payload []byte) {
	ocppcc := &OCPPCC{}
	err := json.Unmarshal(eData.wData, ocppcc)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	log.Printf("Some: %s", ocppcc.ChargeBoxID)
	
	// TESTING!!!!
	return ocppcc.ChargeBoxID, eData.wData
}

func (eData *EVSEdata) Write(data []byte) (n int, err error) {
	for _, b := range data {
		eData.wData = append(eData.wData, b)
	}
	log.Printf("Write: %s", eData.wData)
	eData.wData = data
	
	return len(data), nil
}

func (eData *EVSEdata) Close() error {
	eData.Close()

	return nil
}
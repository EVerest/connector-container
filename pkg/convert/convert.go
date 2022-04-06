/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package convert

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type OCPPCC struct {
	Timestamp     int64 					`json:"timestamp"`
	MessageTypeID string					`json:"messageTypeId"`
	ChargeBoxID   string 					`json:"chargeBoxId"`
	MessageID     string 					`json:"messageId"`
	Action        string 					`json:"action"`
	Payload       map[string]interface {}	`json:"payload"`
}

type EVSEdata struct {
	read chan OCPPCC
	write chan OCPPCC
}

var eData EVSEdata
const call 		= "2"
const callError = "4"

func NewEVSEdata() *EVSEdata {
	eData = EVSEdata{}
	eData.read = make(chan OCPPCC, 100)
	eData.write = make(chan OCPPCC, 100)
	
	return &eData
}

func (eData *EVSEdata) ConnectionReader(URIpath string, b []byte) {	
	ocppcc := OCPPCC{}
	var arr []interface{}

	err := json.Unmarshal(b, &arr)
	if err != nil {
		eData.sendErrorToServer(URIpath, err.Error())
		return
	}

	if len(arr) < 4 {
		eData.sendErrorToServer(URIpath, "length of OCPP array less than 4")
		return
	}

	q := floatToString(arr[0].(float64))

	ocppcc.MessageTypeID	= q
	ocppcc.MessageID 		= fmt.Sprintf("%s", arr[1])
	ocppcc.Action 			= fmt.Sprintf("%s", arr[2])
	ocppcc.Timestamp 		= time.Now().UnixMilli()
	ocppcc.Payload 			= arr[3].(map[string]interface {})
	ocppcc.ChargeBoxID 		= URIpath

	eData.read <- ocppcc
}

func (eData *EVSEdata) Read(b []byte) (int, error) {
	ocppcc := <-eData.read
	bytes, err := json.Marshal(ocppcc)
	if err != nil {
		eData.sendErrorToServer("unknown", err.Error())
		return 0, err
	}

	n := copy(b, bytes)
    return n, nil
}

func (eData *EVSEdata) ConnectionWriter() (URIpath string, payload []byte) {
	ocppcc := <-eData.write

	p, _ := json.Marshal(ocppcc.Payload)
	var sb strings.Builder
	sb.WriteString("[")
	sb.WriteString(ocppcc.MessageTypeID)
	sb.WriteString(`,"`)
	sb.WriteString(ocppcc.MessageID)
	sb.WriteString(`","`)
	sb.WriteString(ocppcc.Action)
	sb.WriteString(`",`)
	sb.WriteString(string(p))
	sb.WriteString(`]`)
	message := sb.String()

	return ocppcc.ChargeBoxID, []byte(message)
}

func (eData *EVSEdata) Write(data []byte) (n int, err error) {
	m := map[string]interface{}{}
    e := json.Unmarshal([]byte(data), &m)
    if e != nil {
		eData.sendErrorToServer("unknown", "malformed write message")
        return
    }

	ocppcc := &OCPPCC{}
	ocppcc.Action 			= m["action"].(string)
	ocppcc.ChargeBoxID 		= m["chargeBoxId"].(string)
	ocppcc.MessageID 		= m["messageId"].(string)
	ocppcc.MessageTypeID 	= m["messageTypeId"].(string)
	ocppcc.Timestamp 		= int64(m["timestamp"].(float64))
	ocppcc.Payload 			= m["payload"].(map[string]interface {})

	eData.write <- *ocppcc
	return len(data), nil
}

func (eData *EVSEdata) Close() error {
	eData.Close()

	return nil
}

func (eData *EVSEdata) ConnectEvent(URIpath string) {
	eData.send(URIpath, call, "Connect", "")
}

func (eData *EVSEdata) DisconnectEvent(URIpath string) {
	eData.send(URIpath, call, "Disconnect", "")
}

func (eData *EVSEdata) sendErrorToServer(URIpath string, err string) {
	eData.send(URIpath, callError, "Error", err)
}

func (eData *EVSEdata) send(URIpath string, typeID string, action string, err string) {
	ocppcc := OCPPCC{}

	ocppcc.MessageTypeID 	= typeID
	ocppcc.Action 			= action
	ocppcc.ChargeBoxID 		= URIpath
	ocppcc.MessageID 		= uuid.NewString()
	ocppcc.Timestamp 		= time.Now().UnixMilli()

	m := make(map[string]interface{})

	m["error"] = err

	ocppcc.Payload = m

	eData.read <- ocppcc
}

func floatToString(num float64) string {
    s := fmt.Sprintf("%.4f", num)
    return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}
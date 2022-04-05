/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package convert

import (
	// "bytes"
	"encoding/json"

	// "encoding/json"
	// "io"
	// "log"
	"testing"
)

/**
Refactor to use table driven tests
**/
func TestConvert(t *testing.T) {
	chargeBoxId := "a-charge-box-id"
	message := []byte(`[2,"a-message-id","BootNotification",{}]`)
	// ocppccMessage := []byte(`{"timestamp":3712349825,"messageTypeID":0,"chargeBoxID":"a-charge-box-id","messageID":"a-message-id","action":"BootNotification","payload":{}}`)

	t.Run("populates messageID from ocpp byte array", func(t *testing.T) {

		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, message)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)
		
		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)
		want := "a-message-id"

		if want != got.MessageID {
			t.Errorf("\ngot %s, \nwant %s", got.MessageID, want)
		}
	})

	t.Run("populates messageTypeID from ocpp byte array", func(t *testing.T) {
		
		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, message)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)
		
		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)
		want := uint8(2)

		if want != got.MessageTypeID {
			t.Errorf("\ngot %d, \nwant %d", got.MessageTypeID, want)
		}
	})

	t.Run("populates timestamp from ocpp byte array", func(t *testing.T) {
		
		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, message)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)
		
		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)
		want := uint32(3712349825)

		if got.Timestamp < want {
			t.Errorf("\ngot %d, \nwant %d", got.Timestamp, want)
		}
	})

	t.Run("populates chargeBoxID from ocpp byte array", func(t *testing.T) {
		
		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, message)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)
		
		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)
		want := "a-charge-box-id"

		if want != got.ChargeBoxID{
			t.Errorf("\ngot %s, \nwant %s", got.ChargeBoxID, want)
		}
	})

	//TODO: Test type assertion errors
}
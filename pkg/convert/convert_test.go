/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package convert

import (
	"encoding/json"
	"log"
	"testing"
)

/**
Refactor to use table driven tests
**/
func TestConvert(t *testing.T) {
	chargeBoxId := "a-charge-box-id"
	message := []byte(`[2,"a-message-id","an-action",{"key":"value"}]`)

	t.Run("sends error when wrong call type type", func(t *testing.T) {
		badMessageWrongCallTypeType := []byte(`["2",a-message-id","an-action",{"key":"value"}]`)

		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, badMessageWrongCallTypeType)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)

		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)

		if got.Action != "Error" {
			t.Fatalf("Want Error got %s", got.Action)
		}
		
		log.Printf("Number: %d Buffer: %s", n, buf)
	})

	t.Run("sends error when missing fields", func(t *testing.T) {
		badMessageMissingField := []byte(`[2,"a-message-id",{"key":"value"}]`)

		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, badMessageMissingField)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)

		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)

		if got.Action != "Error" {
			t.Fatalf("Want Error got %s", got.Action)
		}
		
		log.Printf("Number: %d Buffer: %s", n, buf)
	})

	t.Run("sends error to backend when message malformed", func(t *testing.T) {
		badMessageMissingQuote := []byte(`[2,a-message-id","an-action",{"key":"value"}]`)

		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, badMessageMissingQuote)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)

		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)

		if got.Action != "Error" {
			t.Fatalf("Want Error got %s", got.Action)
		}
		
		log.Printf("Number: %d Buffer: %s", n, buf)
	})

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
		want := "2"

		if want != got.MessageTypeID {
			t.Errorf("\ngot %s, \nwant %s", got.MessageTypeID, want)
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

		if got.Timestamp > want {
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

	t.Run("populates action from ocpp byte array", func(t *testing.T) {
		
		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, message)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)
		
		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)
		want := "an-action"

		if want != got.Action{
			t.Errorf("\ngot %s, \nwant %s", got.Action, want)
		}
	})

	t.Run("populates payload from ocpp byte array", func(t *testing.T) {

		type empty struct {}
		
		data := NewEVSEdata()
		data.ConnectionReader(chargeBoxId, message)
		
		buf := make([]byte, 1024)
		n, _ := data.Read(buf)
		
		sizedBuffer := buf[:n]
		got := &OCPPCC{}
		json.Unmarshal(sizedBuffer, got)

		a := got.Payload

		if a["key"] != "value" {
			t.Errorf("\ngot %s, \nwant %s", a["key"], "value")
		}
	})
}
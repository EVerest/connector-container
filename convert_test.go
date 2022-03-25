/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import (
	"testing"
	"bytes"
)

func TestConvert(t *testing.T) {
	chargeBoxId := "a-charge-box-id"
	message := []byte(`[2,"a-message-id","BootNotification",{}]`)

	t.Run("converts ocppcc to a JSON encoded ocpp byte slice", func (t *testing.T)  {
		ocppcc := ocppcc{}
		ocppcc.messageId = "a-message-id"
		ocppcc.action = "BootNotification"
		ocppcc.messageTypeId = 2
		ocppcc.payload = make(map[string]interface {})

		got := toOcppByteSlice(ocppcc)
		want := []byte(`[2,"a-message-id","BootNotification",{}]`)

		res := bytes.Compare(got, want)

		if res != 0 {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("populates ocppcc.action from ocpp byte array", func (t *testing.T)  {
		
		got := fromOcppByteSlice(chargeBoxId, message)	
		want := ocppcc {
			action: "BootNotification",
		}
				
		if got.action != want.action {
			t.Errorf("got %v, want %v", got.action, want.action)
		}
	})

	t.Run("populates messageId from ocpp byte array", func (t *testing.T)  {

		got := fromOcppByteSlice(chargeBoxId, message)	
		want := ocppcc {
			messageId: "a-message-id",
		}
				
		if got.messageId != want.messageId {
			t.Errorf("got %v, want %v", got.messageId, want.messageId)
		}
	})

	t.Run("populates unspecified payload from ocpp byte array", func (t *testing.T)  {

		got := fromOcppByteSlice(chargeBoxId, message)	

		if got.payload == nil {
			t.Error("Payload has been lost")
		}
	})

	t.Run("populates messageId from ocpp byte array", func (t *testing.T)  {

		got := fromOcppByteSlice(chargeBoxId, message)	
		want := ocppcc {
			messageId: "a-message-id",
		}
				
		if got.messageId != want.messageId {
			t.Errorf("got %v, want %v", got.messageId, want.messageId)
		}
	})

	t.Run("creates and populates a timestamp", func (t *testing.T)  {

		got := fromOcppByteSlice(chargeBoxId, message)	
				
		if got.timestamp <= 0 {
			t.Errorf("no timestamp set")
		}
	})

	t.Run("passes thru a chargeBoxId", func (t *testing.T)  {

		got := fromOcppByteSlice(chargeBoxId, message)	
				
		if got.chargeBoxId == "" {
			t.Errorf("no chargeBoxId set")
		}
	})
}
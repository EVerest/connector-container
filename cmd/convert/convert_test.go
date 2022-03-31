/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package convert

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"testing"
)

func TestConvert(t *testing.T) {
	chargeBoxId := "a-charge-box-id"
	message := []byte(`[2,"a-message-id","BootNotification",{}]`)

	t.Run("converts ocpp to a JSON encoded ocppcc byte slice with buffer size 255", func(t *testing.T) {
		
		reader := NewEVSEreader()
		reader.ConnectionReader(chargeBoxId, message)
		
		p := make([]byte, 255)
		collect := []byte{}
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				break
			}
			collect = append(collect, p[:n]...)
		}

		want := []byte(`{"timestamp":3712349825,"MessageTypeID":0,"ChargeBoxID":"a-charge-box-id","MessageID":"a-message-id","Action":"BootNotification","Payload":{}}`)
		got := collect

		res := bytes.Compare(want, got)

		if res != -1 {
			t.Errorf("\ngot %s, \nwant %s", got, want)
		}
	})

	t.Run("converts ocpp to a JSON encoded ocppcc byte slice with buffer size 16", func(t *testing.T) {
		
		reader := NewEVSEreader()
		reader.ConnectionReader(chargeBoxId, message)
		
		p := make([]byte, 16)
		collect := []byte{}
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				break
			}
			collect = append(collect, p[:n]...)
		}

		want := []byte(`{"timestamp":3712349825,"MessageTypeID":0,"ChargeBoxID":"a-charge-box-id","MessageID":"a-message-id","Action":"BootNotification","Payload":{}}`)
		got := collect

		res := bytes.Compare(want, got)

		if res != -1 {
			t.Errorf("\ngot %s, \nwant %s", got, want)
		}
	})

	t.Run("converts ocpp to a JSON encoded ocppcc byte slice with buffer size 1", func(t *testing.T) {
		
		reader := NewEVSEreader()
		reader.ConnectionReader(chargeBoxId, message)
		
		p := make([]byte, 1)
		collect := []byte{}
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				break
			}
			collect = append(collect, p[:n]...)
		}

		want := []byte(`{"timestamp":3712349825,"MessageTypeID":0,"ChargeBoxID":"a-charge-box-id","MessageID":"a-message-id","Action":"BootNotification","Payload":{}}`)
		got := collect

		res := bytes.Compare(want, got)

		if res != -1 {
			t.Errorf("\ngot %s, \nwant %s", got, want)
		}
	})

	t.Run("populates ocppcc.action from ocpp byte array", func(t *testing.T) {
		
		reader := NewEVSEreader()
		reader.ConnectionReader(chargeBoxId, message)
		
		p := make([]byte, 4)
		collect := []byte{}
		for {
			n, err := reader.Read(p)
			if err == io.EOF {
				break
			}
			collect = append(collect, p[:n]...)
		}
		got := OCPPCC{}
		err := json.Unmarshal(collect, &got)
		if err != nil {
			log.Println(err)
		}

		want := OCPPCC{
			Action: "BootNotification",
		}	
		
		if want.Action != got.Action {
			t.Errorf("\ngot %v, \nwant %v", got, want)
		}
	})

	// t.Run("populates ocppcc.action from ocpp byte array", func(t *testing.T) {

	// 	got := FromConnection(chargeBoxId, message)
	// 	want := OCPPCC{
	// 		action: "BootNotification",
	// 	}

	// 	if got.action != want.action {
	// 		t.Errorf("got %v, want %v", got.action, want.action)
	// 	}
	// })

	// t.Run("populates messageId from ocpp byte array", func(t *testing.T) {

	// 	got := FromConnection(chargeBoxId, message)
	// 	want := OCPPCC{
	// 		messageID: "a-message-id",
	// 	}

	// 	if got.messageID != want.messageID {
	// 		t.Errorf("got %v, want %v", got.messageID, want.messageID)
	// 	}
	// })

	// t.Run("populates unspecified payload from ocpp byte array", func(t *testing.T) {

	// 	got := FromConnection(chargeBoxId, message)

	// 	if got.payload == nil {
	// 		t.Error("Payload has been lost")
	// 	}
	// })

	// t.Run("populates messageId from ocpp byte array", func(t *testing.T) {

	// 	got := FromConnection(chargeBoxId, message)
	// 	want := OCPPCC{
	// 		messageID: "a-message-id",
	// 	}

	// 	if got.messageID != want.messageID {
	// 		t.Errorf("got %v, want %v", got.messageID, want.messageID)
	// 	}
	// })

	// t.Run("creates and populates a timestamp", func(t *testing.T) {

	// 	got := FromConnection(chargeBoxId, message)

	// 	if got.timestamp <= 0 {
	// 		t.Errorf("no timestamp set")
	// 	}
	// })

	// t.Run("passes thru a chargeBoxId", func(t *testing.T) {

	// 	got := FromConnection(chargeBoxId, message)

	// 	if got.chargeBoxID == "" {
	// 		t.Errorf("no chargeBoxId set")
	// 	}
	// })

	//TODO: Test type assertion errors
}

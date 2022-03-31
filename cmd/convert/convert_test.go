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
/**
Refactor to use table driven tests
**/
func TestConvert(t *testing.T) {
	chargeBoxId := "a-charge-box-id"
	message := []byte(`[2,"a-message-id","BootNotification",{}]`)
	ocppccMessage := []byte(`{"timestamp":3712349825,"messageTypeID":0,"chargeBoxID":"a-charge-box-id","messageID":"a-message-id","action":"BootNotification","payload":{}}`)

	t.Run("Writes", func(t *testing.T) {

		writer := NewEVSEWriter()
		// p := make([]byte, 4)
		n, err := writer.Write(ocppccMessage)
		if err != nil {
			t.Fatalf("Error")
		}

		log.Printf("Num: %d", n)
		
		// reader := NewEVSEreader()
		// reader.ConnectionReader(chargeBoxId, message)
		
		// p := make([]byte, 4)
		// collect := []byte{}
		// for {
		// 	n, err := reader.Read(p)
		// 	if err == io.EOF {
		// 		break
		// 	}
		// 	collect = append(collect, p[:n]...)
		// }

		// want := []byte(`{"timestamp":3712349825,"messageTypeID":0,"chargeBoxID":"a-charge-box-id","messageID":"a-message-id","action":"BootNotification","payload":{}}`)
		// got := collect

		// res := bytes.Compare(want, got)

		// if res != -1 {
		// 	t.Errorf("\ngot %s, \nwant %s", got, want)
		// }
	})

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

		want := []byte(`{"timestamp":3712349825,"messageTypeID":0,"chargeBoxID":"a-charge-box-id","messageID":"a-message-id","action":"BootNotification","payload":{}}`)
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

	t.Run("populates messageID from ocpp byte array", func(t *testing.T) {
		
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
			MessageID: "a-message-id",
		}	
		
		if want.MessageID != got.MessageID {
			t.Errorf("\ngot %v, \nwant %v", got.MessageID, want.MessageID)
		}
	})

	t.Run("populates messageTypeID from ocpp byte array", func(t *testing.T) {
		
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
			MessageTypeID: 2,
		}	
		
		if want.MessageTypeID != got.MessageTypeID {
			t.Errorf("\ngot %v, \nwant %v", got.MessageTypeID, want.MessageTypeID)
		}
	})

	t.Run("populates timestamp from ocpp byte array", func(t *testing.T) {
		
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
			Timestamp: 0,
		}	
		
		if want.Timestamp > got.Timestamp {
			t.Errorf("\ngot %v, \nwant %v", got.Timestamp, want.Timestamp)
		}
	})

	t.Run("populates chargeBoxID from ocpp byte array", func(t *testing.T) {
		
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
			ChargeBoxID: "a-charge-box-id",
		}	
		
		if want.ChargeBoxID > got.ChargeBoxID {
			t.Errorf("\ngot %v, \nwant %v", got.ChargeBoxID, want.ChargeBoxID)
		}
	})

	//TODO: Test type assertion errors
}
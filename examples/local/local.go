/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

 package main

 import (
	 "io"
	 "log"
	 "sync"
	 "time"
 
	 "github.com/ChargeNet-Stations/ocpp-cloud-connector/pkg/connection"
	 "github.com/ChargeNet-Stations/ocpp-cloud-connector/pkg/convert"
	 "github.com/ChargeNet-Stations/ocpp-cloud-connector/pkg/server"
	 "github.com/ChargeNet-Stations/ocpp-cloud-connector/examples/store"
 )
 
 func main() {
	 storage := make(store.LocalStore)
	 converter := convert.NewEVSEdata()
	 connectionOptions := connection.ConnectionOptions{
		 SubProtocol:     "ocpp1.6",
		 ConnectionStore: storage,
		 Converter:       converter,
	 }
 
	 connectionHandler := connection.NewConnectionHandler(connectionOptions)
 
	 serverOptions := server.ServerOptions{
		 Addr:            "0.0.0.0:8080",
		 Handler:         connectionHandler,
		 RootPath:        "/",
		 HealthCheckPath: "/health",
	 }
 
	 var wg sync.WaitGroup
	 wg.Add(4)
 
	 go server.StartServer(serverOptions)
 
	 go readThis(converter, &wg)
	 go writeToCat(converter, &wg)
	 go writeToDog(converter, &wg)
 
	 wg.Wait()
 }
 
 func readThis(t *convert.EVSEdata, wg *sync.WaitGroup) {
	 defer wg.Done()
	 buf := make([]byte, 1024)
	 for {
		 time.Sleep(1 * time.Second)
		 n, err := t.Read(buf)
		 if err != io.EOF {
			 log.Printf("Read:  %s", buf[:n])
		 }
	 }
 }
 
 func writeToCat(t *convert.EVSEdata, wg *sync.WaitGroup) {
	 defer wg.Done()
	 ocppccMessage := []byte(`{"timestamp:3712349825,"messageTypeId":"2","chargeBoxId":"cat","messageId":"a-message-id","action":"BootNotification","payload":{"dog":"cat"}}`)
	 for {
		 time.Sleep(1 * time.Second)
		 log.Printf("Write: %s", ocppccMessage)
		 _, err := t.Write(ocppccMessage)
		 if err != nil {
			 log.Println("Error")
		 }
	 }
 }
 
 func writeToDog(t *convert.EVSEdata, wg *sync.WaitGroup) {
	 defer wg.Done()
	 ocppccMessage := []byte(`{"timestamp":3712349825,"messageTypeId":"2","chargeBoxId":"dog","messageId":"a-message-id","action":"BootNotification","payload":{"dog":"cat"}}`)
	 for {
		 time.Sleep(1 * time.Second)
		 log.Printf("Write: %s", ocppccMessage)
		 _, err := t.Write(ocppccMessage)
		 if err != nil {
			 log.Println("Error")
		 }
	 }
 }
 
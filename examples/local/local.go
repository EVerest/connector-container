/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/ChargeNet-Stations/ocpp-cloud-connector/examples/store"
	"github.com/ChargeNet-Stations/ocpp-cloud-connector/pkg/connection"
	"github.com/ChargeNet-Stations/ocpp-cloud-connector/pkg/convert"
	"github.com/ChargeNet-Stations/ocpp-cloud-connector/pkg/server"
)

func main() {
	storage := make(store.LocalStore)
	converter := convert.NewEVSEdata()

	ch := connection.ConnectionHandler{
		SubProtocol:     "ocpp1.6",
		ConnectionStore: storage,
		Converter:       converter,
	}

	serverOptions := server.ServerOptions{
		Addr:            "0.0.0.0:8080",
		Handler:         ch.Handler,
		RootPath:        "/",
		HealthCheckPath: "/health",
	}

	var wg sync.WaitGroup
	wg.Add(5)

	go server.StartServer(serverOptions)

	go readThis(converter, &wg)
	go writeToCat(converter, &wg)
	go writeToDog(converter, &wg)
	go writeToBob(converter, &wg)
	// go writeFromConsole(converter, &wg)

	wg.Wait()
}

func writeFromConsole(t *convert.EVSEdata, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	_, err := t.Write([]byte(text))
	if err != nil {
		log.Println("Error")
	}

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

	for {
		ocppccMessage := []byte(`{"timestamp":` + getTime() + `,"messageTypeId":"2","chargeBoxId":"cat","messageId":"a-message-id","action":"BootNotification","payload":{"dog":"cat"}}`)
		time.Sleep(1 * time.Second)
		_, err := t.Write(ocppccMessage)
		if err != nil {
			log.Println("Error")
		}
		log.Printf("Wrote: %s", ocppccMessage)
	}
}

func writeToDog(t *convert.EVSEdata, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		badOCPPCCMessage := []byte(`{"timestamp":` + getTime() + `,"messageTypeId":"2","chargeBoxId":"dog","messageId":"a-message-id","action":"BootNotification","payload":{"dog":"cat"}}`)
		time.Sleep(1 * time.Second)
		_, err := t.Write(badOCPPCCMessage)
		if err != nil {
			log.Println("Error")
		}
		log.Printf("Write: %s", badOCPPCCMessage)
	}
}

func writeToBob(t *convert.EVSEdata, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		badOCPPCCMessage := []byte(`{"timestamp":` + getTime() + `,"messageTypeId":"2","chargeBoxId":"bob","messageId":"a-message-id","action":"BootNotification","payload":{"dog":"cat"}}`)
		time.Sleep(1 * time.Second)
		_, err := t.Write(badOCPPCCMessage)
		if err != nil {
			log.Println("Error")
		}
		log.Printf("Write: %s", badOCPPCCMessage)
	}
}

func getTime() string {
	return fmt.Sprintf("%d", time.Now().UnixMilli())
}

// {"timestamp":1649433271379,"messageTypeId":"2","chargeBoxId":"cat","messageId":"a-message-id","action":"BootNotification","payload":{"dog":"cat"}}
// {"timestamp":1649433271379,"messageTypeId":"2","chargeBoxId":"dog","messageId":"a-message-id","action":"BootNotification","payload":{"dog":"cat"}}

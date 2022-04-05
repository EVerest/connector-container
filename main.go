/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import (
	"io"
	"log"
	"sync"

	"github.com/ChargeNet-Stations/ocpp-cloud-connector/cmd/connection"
	"github.com/ChargeNet-Stations/ocpp-cloud-connector/cmd/convert"
	"github.com/ChargeNet-Stations/ocpp-cloud-connector/cmd/server"
)

func main() {
	storage := make(localStore)
	thing := convert.NewEVSEdata()
	connectionOptions := connection.ConnectionOptions {
		SubProtocol		: "ocpp1.6",
		ConnectionStore	: storage,
		Doer: thing,
	}

	connectionHandler := connection.NewConnectionHandler(connectionOptions)
	
	serverOptions := server.ServerOptions {
		Addr			: "0.0.0.0:8080",
		Handler			: connectionHandler,
		RootPath		: "/",
		HealthCheckPath	: "/health",
 	}
	 
	 var wg sync.WaitGroup
	 wg.Add(1)
	 
	 go server.StartServer(serverOptions)
	 go readThis(thing, &wg)
	 
	 wg.Wait()
}

func readThis(t *convert.EVSEdata, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, 1024)
	for {
		n, err := t.Read(buf)
		if err != io.EOF {
			log.Printf("Read:  %s", buf[:n])
		}
	}
}

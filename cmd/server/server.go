/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package server

import (
	"log"
	"net/http"
)

type ServerOptions struct {
	addr 			string
	handler 		func(http.ResponseWriter, *http.Request)
	rootPath 		string
	healthCheckPath string
}

func StartServer(opts ServerOptions) {
	http.HandleFunc(opts.rootPath, opts.handler)
	http.HandleFunc(opts.healthCheckPath, health)
	
	log.Fatal(http.ListenAndServe(opts.addr, nil))
}

func health(writer http.ResponseWriter, request *http.Request) { /*noop*/ }
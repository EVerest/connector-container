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
	Addr            string
	Handler         func(http.ResponseWriter, *http.Request)
	RootPath        string
	HealthCheckPath string
}

func StartServer(opts ServerOptions) {
	http.HandleFunc(opts.RootPath, opts.Handler)
	http.HandleFunc(opts.HealthCheckPath, health)

	log.Fatal(http.ListenAndServe(opts.Addr, nil))
}

func health(writer http.ResponseWriter, request *http.Request) { /*noop*/ }

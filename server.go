/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

 package main

 import (
	 "log"
	 "net/http"
	 "flag"
 )
 
 var addr = flag.String("addr", "localhost:8080", "http service address")
 
 func serve() {
	 flag.Parse()
	 http.HandleFunc("/", connectionHandler)
	 http.HandleFunc("/health", health) 
	 log.Fatal(http.ListenAndServe(*addr, nil))
 }
 
 func health(writer http.ResponseWriter, request *http.Request) { /*noop*/ }
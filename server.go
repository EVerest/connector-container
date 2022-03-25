/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

 package main

 import (
	 "log"
	 "net/http"
 )
  // Todo: externalize health endpoint path as env var
 func StartServer(addr string) {
	http.HandleFunc("/", connectionHandler)
	http.HandleFunc("/health", health) 
	log.Fatal(http.ListenAndServe(addr, nil))
 }
 
 func health(writer http.ResponseWriter, request *http.Request) { /*noop*/ }
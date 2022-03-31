/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

import(

)

func main() {
	// only approrpiate for single instance deployments like local or development.  Not thread safe.
	// storage := make(localStore)
	// connectionOptions := ConnectionOptions {
	// 	subProtocol		: "ocpp1.6",
	// 	connectionStore	: storage,
	// }
	// connectionHandler := NewConnectionHandler(connectionOptions)
	// doer := NewEVSEreader()
	// serverOptions := ServerOptions {
	// 	addr			: "0.0.0.0:8080",
	// 	handler			: connectionHandler,
	// 	rootPath		: "/",
	// 	healthCheckPath	: "/health",
	// 	doer			: doer,
	// }

	// StartServer(serverOptions)
}

// Environmet exposer here only / config
// Define and parse flags here

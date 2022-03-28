/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: CC-BY-4.0
 */

package main

func main() {
	storage := make(localStore)
	StartConnectionHandler(&storage)
	StartServer("0.0.0.0:8080")
}

// Environmet exposer here only / config
// Define and parse flags here
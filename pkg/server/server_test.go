/**
 * Copyright 2022 Charge Net Stations and Contributors.
 * SPDX-License-Identifier: Apache-2.0
 */
package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	t.Run("health check at GET health returns 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		health(response, request)

		got := response.Result().StatusCode
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

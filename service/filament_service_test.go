package service_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/panicpanicpanic/filament/lifx"
	"github.com/panicpanicpanic/filament/service"
)

func TestLIFXGet(t *testing.T) {
	var client lifx.Client

	mock, err := os.Open("../mocks/lifx_devices_payload.json")
	if err != nil {
		t.Errorf("it should not throw an error, got %d", err)
	}

	payload, err := ioutil.ReadAll(mock)
	if err != nil {
		t.Errorf("it should not throw an error, got %d", err)
	}
	defer mock.Close()

	t.Run("when AccessToken is missing from LIFXReq struct", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}))

		client.Endpoint = server.URL

		_, err = service.Get(&client)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an AccessToken, got %d", err)
		}
	})

	t.Run("when URL is missing from LIFXReq struct", func(t *testing.T) {
		_ = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}))

		client.AccessToken = "someRandomToken"
		client.Endpoint = ""

		_, err = service.Get(&client)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an AccessToken, got %d", err)
		}
	})

	t.Run("when the LIFX API returns a non-200 HTTP status code for GET request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(payload)
		}))

		client.AccessToken = "someRandomToken"
		client.Endpoint = server.URL

		_, err = service.Get(&client)
		if err == nil {
			t.Errorf("it should have thrown an error for returning a 500 HTTP status, got %d", err)
		}
	})

	t.Run("when the expected response is not []device.Device", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}))

		client.AccessToken = "someRandomToken"
		client.Endpoint = server.URL

		body, err := service.Get(&client)
		if len(body) == 0 {
			t.Errorf("it should have returned 1 empty device, got %d", err)
		}
	})
}

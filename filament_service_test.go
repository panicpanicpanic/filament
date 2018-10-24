package filament_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/panicpanicpanic/filament"
)

func TestLIFXGetRequest(t *testing.T) {
	var lifx filament.LIFXReq
	var filament filament.Filament

	mock, err := os.Open("./mocks/lifx_devices_payload.json")
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

		lifx.URL = server.URL

		_, err = filament.LIFXGetRequest(&lifx)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an AccessToken, got %d", err)
		}
	})

	t.Run("when URL is missing from LIFXReq struct", func(t *testing.T) {
		_ = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}))

		lifx.LIFXClient.AccessToken = "someRandomToken"
		lifx.URL = ""

		_, err = filament.LIFXGetRequest(&lifx)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an AccessToken, got %d", err)
		}
	})

	t.Run("when the LIFX API returns a non-200 HTTP status code for GET request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(payload)
		}))

		lifx.LIFXClient.AccessToken = "someRandomToken"
		lifx.URL = server.URL

		_, err = filament.LIFXGetRequest(&lifx)
		if err == nil {
			t.Errorf("it should have thrown an error for returning a 500 HTTP status, got %d", err)
		}
	})

	t.Run("when the expected response is not []device.Device", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(payload)
		}))

		lifx.LIFXClient.AccessToken = "someRandomToken"
		lifx.URL = server.URL

		devices, err := filament.LIFXGetRequest(&lifx)
		if len(devices) == 0 {
			t.Errorf("it should have returned 1 empty device, got %d", err)
		}
	})
}

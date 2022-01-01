package lifx

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestServiceGet(t *testing.T) {
	var client lifx.Client
	var err error

	payload := []byte(`[
		{
		  "id": "123",
		  "uuid": "123",
		  "label": "Main",
		  "connected": true,
		  "power": "on",
		  "color": {
			"hue": 240,
			"saturation": 1,
			"kelvin": 4000
		  },
		  "brightness": 1,
		  "group": {
			"id": "123",
			"name": "My Room"
		  },
		  "location": {
			"id": "123",
			"name": "My Home"
		  },
		  "product": {
			"name": "LIFX BR30",
			"identifier": "lifx_br30",
			"company": "LIFX",
			"capabilities": {
			  "has_color": true,
			  "has_variable_color_temp": true,
			  "has_ir": false,
			  "has_chain": false,
			  "has_multizone": false,
			  "min_kelvin": 2500,
			  "max_kelvin": 9000
			}
		  },
		  "last_seen": "2018-10-21T04:13:04Z",
		  "seconds_since_seen": 0
		}]
	`)

	t.Run("when AccessToken is missing from lifx.Client", func(t *testing.T) {
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

	t.Run("when Endpoint is missing from lifx.Client", func(t *testing.T) {
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

	t.Run("when the expected response is not []byte", func(t *testing.T) {
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
		if reflect.TypeOf(body).String() != "[]uint8" {
			t.Errorf("it should have returned a []byte, got %d", reflect.TypeOf(body))
		}
	})
}

func TestServicePut(t *testing.T) {
	var client lifx.Client
	var err error

	t.Run("when AccessToken is missing from lifx.Client", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMultiStatus)
			w.Write(response)
		}))

		client.Endpoint = server.URL

		_, err = service.Put(&client, payload)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an AccessToken, got %d", err)
		}
	})

	t.Run("when Endpoint is missing from lifx.Client", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMultiStatus)
			w.Write(response)
		}))

		client.AccessToken = server.URL
		client.Endpoint = ""

		_, err = service.Put(&client, payload)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an Endpoint, got %d", err)
		}
	})

	t.Run("when the LIFX API returns a non-207 HTTP status code for PUT request", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
		}))

		client.AccessToken = "someRandomToken"
		client.Endpoint = server.URL

		_, err = service.Put(&client, payload)
		if err == nil {
			t.Errorf("it should have thrown an error for returning a 500 HTTP status, got %d", err)
		}
	})

	t.Run("when the expected response is not []byte", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		}))

		client.AccessToken = "someRandomToken"
		client.Endpoint = server.URL

		body, err := service.Put(&client, payload)
		if len(body) == 0 {
			t.Errorf("it should have returned 1 empty device, got %d", err)
		}
		if reflect.TypeOf(body).String() != "[]uint8" {
			t.Errorf("it should have returned a []byte, got %d", reflect.TypeOf(body))
		}
	})
}

func TestServicePost(t *testing.T) {
	var client lifx.Client
	var err error

	t.Run("when AccessToken is missing from lifx.Client", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMultiStatus)
			w.Write(response)
		}))

		client.Endpoint = server.URL

		_, err = service.Post(&client, payload)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an AccessToken, got %d", err)
		}
	})

	t.Run("when Endpoint is missing from lifx.Client", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusMultiStatus)
			w.Write(response)
		}))

		client.AccessToken = server.URL
		client.Endpoint = ""

		_, err = service.Post(&client, payload)
		if err == nil {
			t.Errorf("it should have thrown an error for not supplying an Endpoint, got %d", err)
		}
	})

	t.Run("when the LIFX API returns a non-207 HTTP status code for POST request", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(response)
		}))

		client.AccessToken = "someRandomToken"
		client.Endpoint = server.URL

		_, err = service.Post(&client, payload)
		if err == nil {
			t.Errorf("it should have thrown an error for returning a 500 HTTP status, got %d", err)
		}
	})

	t.Run("when the expected response is not []byte", func(t *testing.T) {
		response := []byte(`
			{
				"results": [
					{
						"id": "d073d52260ef",
						"status": "ok",
						"label": "Main"
					}
				]
			}
		`)

		payload := []byte(`
			{
				"power":"on"
			}
		`)

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write(response)
		}))

		client.AccessToken = "someRandomToken"
		client.Endpoint = server.URL

		body, err := service.Post(&client, payload)
		if len(body) == 0 {
			t.Errorf("it should have returned 1 empty device, got %d", err)
		}
		if reflect.TypeOf(body).String() != "[]uint8" {
			t.Errorf("it should have returned a []byte, got %d", reflect.TypeOf(body))
		}
	})
}

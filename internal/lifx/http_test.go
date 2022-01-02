package lifx

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGet(t *testing.T) {
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

	tests := []struct {
		accessToken string
		endpoint    string
		name        string
		expectedErr bool
	}{
		{
			name:        "when AccessToken is missing",
			endpoint:    "testEndpoint",
			accessToken: "",
			expectedErr: true,
		},
		{
			name:        "when endpoint is missing",
			endpoint:    "",
			accessToken: "1234",
			expectedErr: true,
		},
		{
			name:        "when the LIFX API returns a non-200 HTTP status code",
			endpoint:    "testEndpoint",
			accessToken: "1234",
			expectedErr: true,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write(payload)
			}))

			os.Setenv("LIFX_API_ACCESS_TOKEN", tt.accessToken)

			url := fmt.Sprintf("%s/%s", server.URL, tt.endpoint)

			if _, err := Get(url); (err != nil) != tt.expectedErr {
				t.Errorf("expected to get error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

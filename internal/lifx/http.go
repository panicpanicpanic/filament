package lifx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	// APIEndpoint is the URL for the latest LIFX HTTP API
	APIEndpoint           = "https://api.lifx.com/v1"
	GetScenesEndpoint     = "%s/scenes"
	ActivateSceneEndpoint = "%s/scenes/scene_id:%s/activate"
	ValidateColorEndpoint = "%s/color?string=%s"
	GetLightsEndpoint     = "%s/lights/%s"
	SetStateEndpoint      = "%s/lights/%s/state"
	SetStatesEndpoint     = "%s/lights/states"
	StateDeltaEndpoint    = "%s/lights/%s/state/delta"
	CycleEndpoint         = "%s/lights/%s/cycle"
	ToggleEndpoint        = "%s/lights/%s/toggle"
	PulseEndpoint         = "%s/lights/%s/effects/pulse"
	BreatheEndpoint       = "%s/lights/%s/effects/breathe"
)

var (
	// AccessToken references the LIFX API Access Token
	AccessToken = os.Getenv("LIFX_API_ACCESS_TOKEN")
)

// Get makes a GET request to the LIFX HTTP API and returns []byte or error
func Get(endpoint string) ([]byte, error) {
	var (
		body       []byte
		err        error
		httpClient http.Client
		statusCode int
	)

	if endpoint == "" || AccessToken == "" {
		return body, MissingTokenEndpointError
	}

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))
	if err != nil {
		return body, fmt.Errorf(err.Error())
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return body, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return body, fmt.Errorf(err.Error())
	}

	statusCode = response.StatusCode
	responseString := string(body)

	if statusCode > 207 {
		return body, fmt.Errorf("received a %d status code. error: %s", statusCode, responseString)
	}

	return body, nil
}

// Put makes a PUT request to the LIFX HTTP API and returns []byte or error
func Put(endpoint string, payload interface{}) ([]byte, error) {
	var (
		body       []byte
		err        error
		httpClient http.Client
		statusCode int
	)

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if endpoint == "" || AccessToken == "" {
		return body, MissingTokenEndpointError
	}

	request, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(data))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	statusCode = response.StatusCode
	responseString := string(body)

	if statusCode > 207 {
		return body, fmt.Errorf("received a %d status code. error: %s", statusCode, responseString)
	}

	return body, nil
}

// Post makes a POST request to the LIFX HTTP API and returns []byte or error
func Post(endpoint string, payload interface{}) ([]byte, error) {
	var (
		body       []byte
		err        error
		httpClient http.Client
		statusCode int
	)

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if endpoint == "" || AccessToken == "" {
		return body, MissingTokenEndpointError
	}

	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(data))
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	statusCode = response.StatusCode
	responseString := string(body)

	if statusCode > 207 {
		return body, fmt.Errorf("received a %d status code. error: %s", statusCode, responseString)
	}

	return body, nil
}

// ReturnAPIEndpoint constructs and returns the appropriate LIFX
// API endpoint
func ReturnAPIEndpoint(base, args string) string {
	return fmt.Sprintf(base, APIEndpoint, args)
}

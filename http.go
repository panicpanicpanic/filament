package filament

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	getScenesEndpoint     = "%s/scenes"
	activateSceneEndpoint = "%s/scenes/scene_id:%s/activate"
	validateColorEndpoint = "%s/color?string=%s"
	getLightsEndpoint     = "%s/lights/%s"
	setStateEndpoint      = "%s/lights/%s/state"
	setStatesEndpoint     = "%s/lights/states"
	stateDeltaEndpoint    = "%s/lights/%s/state/delta"
	cycleEndpoint         = "%s/lights/%s/cycle"
	toggleEndpoint        = "%s/lights/%s/toggle"
	pulseEndpoint         = "%s/lights/%s/effects/pulse"
	breatheEndpoint       = "%s/lights/%s/effects/breathe"
)

func get(endpoint string, token string) ([]byte, error) {
	if endpoint == "" || token == "" {
		return []byte{}, MissingTokenEndpointError
	}

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return []byte{}, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return []byte{}, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf(err.Error())
	}

	if response.StatusCode > 207 {
		return body, fmt.Errorf("received a %d status code. error: %s", response.StatusCode, string(body))
	}

	return body, nil
}

// put makes a PUT request to the LIFX HTTP API and returns []byte or error
func put(endpoint, token string, payload any) ([]byte, error) {
	if endpoint == "" || token == "" {
		return []byte{}, MissingTokenEndpointError
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	request, err := http.NewRequest(http.MethodPut, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf(err.Error())
	}

	if response.StatusCode > 207 {
		return body, fmt.Errorf("received a %d status code. error: %s", response.StatusCode, string(body))
	}

	return body, nil
}

// post makes a POST request to the LIFX HTTP API and returns []byte or error
func post(endpoint, token string, payload any) ([]byte, error) {
	if endpoint == "" || token == "" {
		return []byte{}, MissingTokenEndpointError
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	request, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	httpClient := http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if response.StatusCode > 207 {
		return body, fmt.Errorf("received a %d status code. error: %s", response.StatusCode, string(body))
	}

	return body, nil
}

// returnAPIEndpoint constructs and returns the appropriate LIFX API endpoint
func (c *Client) returnAPIEndpoint(base, args string) string {
	if args == "" {
		args = "all"
	}
	return fmt.Sprintf(base, c.apiURL, args)
}

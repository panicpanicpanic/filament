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
	APIEndpoint = "https://api.lifx.com/v1"
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

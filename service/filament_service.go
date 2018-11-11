package service

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/panicpanicpanic/filament/lifx"
)

// Get makes a GET request to the LIFX HTTP API and returns []byte or error
func Get(client *lifx.Client) ([]byte, error) {
	var body []byte
	var httpClient http.Client
	var err error
	var statusCode int

	if client.AccessToken == "" || client.Endpoint == "" {
		return body, fmt.Errorf("In order to access the LIFX API, you must supply a valid AccessToken and URL")
	}

	request, err := http.NewRequest(http.MethodGet, client.Endpoint, nil)
	request.Header.Set("Authorization", "Bearer "+client.AccessToken)
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
		return body, fmt.Errorf("Uh oh! You've received a %d status code. Error: %s", statusCode, responseString)
	}

	return body, nil
}

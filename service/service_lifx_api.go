package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/panicpanicpanic/filament/device"
)

// LIFXGetRequest makes a GET request to the LIFX HTTP API
func (s *Service) LIFXGetRequest(lifx *LIFXReq) ([]device.Device, error) {
	var err error
	var devices []device.Device
	var client http.Client
	var statusCode int

	if lifx.AccessToken == "" || lifx.URL == "" {
		return devices, fmt.Errorf("In order to access the LIFX API, you must supply a valid AccessToken and URL")
	}

	request, err := http.NewRequest(http.MethodGet, lifx.URL, nil)
	request.Header.Set("Authorization", lifx.AccessToken)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	statusCode = response.StatusCode
	responseString := string(body)

	if statusCode > 207 {
		return devices, fmt.Errorf("Uh oh! You've recieved a %d status code. Error: %s", statusCode, responseString)
	}

	err = json.Unmarshal(body, &devices)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	return devices, nil
}

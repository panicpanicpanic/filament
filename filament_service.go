package filament

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// LIFXClient represents the LIFX client structure for reaching LIFX HTTP API
type LIFXClient struct {
	AccessToken string
}

// LIFXReq represents the request structure for reaching LIFX HTTP API
type LIFXReq struct {
	LIFXClient LIFXClient
	URL        string
}

// Get makes a GET request to the LIFX HTTP API and returns []byte
func (f *Filament) Get(lifx *LIFXReq) ([]byte, error) {
	var body []byte
	var client http.Client
	var err error
	var statusCode int

	if lifx.LIFXClient.AccessToken == "" || lifx.URL == "" {
		return body, fmt.Errorf("In order to access the LIFX API, you must supply a valid AccessToken and URL")
	}

	request, err := http.NewRequest(http.MethodGet, lifx.URL, nil)
	request.Header.Set("Authorization", "Bearer "+lifx.LIFXClient.AccessToken)
	if err != nil {
		return body, fmt.Errorf(err.Error())
	}

	response, err := client.Do(request)
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

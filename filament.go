package filament

import (
	"fmt"
)

const (
	// LIFXAPIURL is the URL for the latest LIFX HTTP API
	LIFXAPIURL = "https://api.lifx.com/v1"
)

// Filament struct
type Filament struct {
}

// LIFXClient represents the LIFX client structure for reaching LIFX HTTP API
type LIFXClient struct {
	AccessToken string
}

// LIFXReq represents the request structure for reaching LIFX HTTP API
type LIFXReq struct {
	LIFXClient LIFXClient
	URL        string
}

// GetLights returns []Device's that belong to your LIFX account
func (f *Filament) GetLights(client *LIFXClient) ([]Device, error) {
	var devices []Device
	var err error
	var req LIFXReq

	req.LIFXClient.AccessToken = client.AccessToken
	req.URL = LIFXAPIURL + "/lights/all"

	devices, err = f.LIFXGetRequest(&req)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	return devices, nil
}

package filament

import (
	"encoding/json"
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

// GetLights returns []Device that belong to your LIFX account
func (f *Filament) GetLights(client *LIFXClient) ([]Device, error) {
	var body []byte
	var devices []Device
	var err error
	var req LIFXReq

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and URL path
	req.LIFXClient.AccessToken = client.AccessToken
	req.URL = LIFXAPIURL + "/lights/all"

	// Use Get service method to GET info from LIFX HTTP API
	body, err = f.Get(&req)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	// Unmarshal body respone to []Device
	err = json.Unmarshal(body, &devices)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	// If successful GetLights returns []Device
	return devices, nil
}

// GetScenes returns []DeviceScene that belong to your LIFX account
func (f *Filament) GetScenes(client *LIFXClient) ([]DeviceScene, error) {
	var body []byte
	var scenes []DeviceScene
	var err error
	var req LIFXReq

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and URL path
	req.LIFXClient.AccessToken = client.AccessToken
	req.URL = LIFXAPIURL + "/scenes"

	// Use Get service method to GET info from LIFX HTTP API
	body, err = f.Get(&req)
	if err != nil {
		return scenes, fmt.Errorf(err.Error())
	}

	// Unmarshal body respone to []DeviceScene
	err = json.Unmarshal(body, &scenes)
	if err != nil {
		return scenes, fmt.Errorf(err.Error())
	}

	// If successful GetLights returns []DeviceScene
	return scenes, nil
}

// ValidateColor returns a DeviceColor if a valid color string is passed
func (f *Filament) ValidateColor(client *LIFXClient, color string) (DeviceColor, error) {
	var body []byte
	var deviceColor DeviceColor
	var err error
	var req LIFXReq

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and URL path
	req.LIFXClient.AccessToken = client.AccessToken
	req.URL = LIFXAPIURL + "/color?string=" + color

	// Use Get service method to GET info from LIFX HTTP API
	body, err = f.Get(&req)
	if err != nil {
		return deviceColor, fmt.Errorf(err.Error())
	}

	// Unmarshal body respone to []DeviceScene
	err = json.Unmarshal(body, &color)
	if err != nil {
		return deviceColor, fmt.Errorf(err.Error())
	}

	// If successful ValidateColor returns DeviceColor
	return deviceColor, nil
}

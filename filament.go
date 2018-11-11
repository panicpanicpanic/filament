package filament

import (
	"encoding/json"
	"fmt"

	"github.com/panicpanicpanic/filament/device"
	"github.com/panicpanicpanic/filament/lifx"
	"github.com/panicpanicpanic/filament/service"
)

// GetLights returns []device.Device that belong to your LIFX account
func GetLights(client *lifx.Client) ([]device.Device, error) {
	var body []byte
	var devices []device.Device
	var err error

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and URL path
	client.Endpoint = lifx.LIFXAPIURL + "/lights/all"

	body, err = service.Get(client)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &devices)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	// Return []device.Device or return an error
	return devices, nil
}

// GetScenes returns []device.Scene that belong to your LIFX account
func GetScenes(client *lifx.Client) ([]device.Scene, error) {
	var body []byte
	var scenes []device.Scene
	var err error

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/scenes"

	body, err = service.Get(client)
	if err != nil {
		return scenes, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &scenes)
	if err != nil {
		return scenes, fmt.Errorf(err.Error())
	}

	// Return []device.Scene or return an error
	return scenes, nil
}

// ValidateColor returns a device.Color if a valid color string is passed
func ValidateColor(client *lifx.Client, color string) (device.Color, error) {
	var body []byte
	var deviceColor device.Color
	var err error

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/color?string=" + color

	body, err = service.Get(client)
	if err != nil {
		return deviceColor, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &deviceColor)
	if err != nil {
		return deviceColor, fmt.Errorf(err.Error())
	}

	// Return device.Color or return error
	return deviceColor, nil
}

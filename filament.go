package filament

import (
	"encoding/json"
	"fmt"

	"github.com/panicpanicpanic/filament/device"
	"github.com/panicpanicpanic/filament/lifx"
	"github.com/panicpanicpanic/filament/service"
)

// GetLights returns []device.Device that belong to your LIFX account
func GetLights(client *lifx.Client, selector string) ([]device.Device, error) {
	var body []byte
	var devices []device.Device
	var err error

	// If no selector is passed, default to retrieving all lights for your LIFX account
	if selector == "" {
		selector = "all"
	}

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and URL path
	client.Endpoint = lifx.LIFXAPIURL + "/lights/" + selector

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

// SetState sets the state of the lights within the given selector, and returns a LIFX Response
func SetState(client *lifx.Client, selector string, payload interface{}) (lifx.Response, error) {
	var body []byte
	var err error
	var response lifx.Response

	// If no selector is passed, default to setting state for all lights
	if selector == "" {
		selector = "all"
	}

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/lights/" + selector + "/state"

	body, err = service.Put(client, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	// Return lifx.Response or return error
	return response, nil
}

// SetStates sets multiple states across multiple selectors, and returns a LIFX Response
func SetStates(client *lifx.Client, payload interface{}) (lifx.Response, error) {
	var body []byte
	var err error
	var response lifx.Response

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/lights/states"

	body, err = service.Put(client, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	// Return lifx.Response or return error
	return response, nil
}

// ActivateScene activates a scene from your LIFX account
func ActivateScene(client *lifx.Client, sceneUUID string, payload interface{}) (lifx.Response, error) {
	var body []byte
	var err error
	var response lifx.Response

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/scenes/scene_id:" + sceneUUID + "/activate"

	body, err = service.Put(client, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	// Return lifx.Response or return error
	return response, nil
}

// Cycle makes the light(s) cycle to the next or previous state in a list of states
func Cycle(client *lifx.Client, selector string, payload interface{}) (lifx.Response, error) {
	var body []byte
	var err error
	var response lifx.Response

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/lights/" + selector + "/cycle"

	body, err = service.Post(client, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	// Return lifx.Response or return error
	return response, nil
}

// PulseEffect performs a pulse effect by quickly flashing between the given colors
func PulseEffect(client *lifx.Client, selector string, payload interface{}) (lifx.Response, error) {
	var body []byte
	var err error
	var response lifx.Response

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/lights/" + selector + "/effects/pulse"

	body, err = service.Post(client, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	// Return lifx.Response or return error
	return response, nil
}

// BreatheEffect performs a breathe effect by slowly fading between the given colors.
func BreatheEffect(client *lifx.Client, selector string, payload interface{}) (lifx.Response, error) {
	var body []byte
	var err error
	var response lifx.Response

	// In order to access LIFX HTTP API, you must pass a valid AccessToken and Endpoint
	client.Endpoint = lifx.LIFXAPIURL + "/lights/" + selector + "/effects/pulse"

	body, err = service.Post(client, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	// Return lifx.Response or return error
	return response, nil
}

package filament

import (
	"encoding/json"
	"errors"
	"fmt"
)

// GetLights returns []Device that belong to your LIFX account
func GetLights(selector string) ([]Device, error) {
	var (
		body    []byte
		devices []Device
		err     error
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := lifx.returnAPIEndpoint(GetLightsEndpoint, selector); endpoint == "" {
		return devices, errors.New("not a valid endpoint")
	}

	body, err = lifx.Get(endpoint)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &devices)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	return devices, nil
}

// GetScenes returns []Scene that belong to your LIFX account
func GetScenes() ([]Scene, error) {
	var (
		body   []byte
		err    error
		scenes []Scene
	)

	if endpoint := lifx.returnAPIEndpoint(GetScenesEndpoint); endpoint == "" {
		return scenes, errors.New("not a valid endpoint")
	}

	body, err = lifx.Get(endpoint)
	if err != nil {
		return scenes, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &scenes)
	if err != nil {
		return scenes, fmt.Errorf(err.Error())
	}

	return scenes, nil
}

// ValidateColor returns a Color if a valid color string is passed
func ValidateColor(color string) (Color, error) {
	var (
		body        []byte
		deviceColor Color
		err         error
	)

	if endpoint := lifx.returnAPIEndpoint(ValidateColorEndpoint, color); endpoint == "" {
		return deviceColor, errors.New("not a valid endpoint")
	}

	body, err = lifx.Get(endpoint)
	if err != nil {
		return deviceColor, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &deviceColor)
	if err != nil {
		return deviceColor, fmt.Errorf(err.Error())
	}

	return deviceColor, nil
}

// SetState sets the state of the lights within the given selector, and returns a LIFX Response
func SetState(selector string, payload interface{}) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := lifx.returnAPIEndpoint(SetStateEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Put(endpoint, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

// SetStates sets multiple states across multiple selectors, and returns a LIFX Response
func SetStates(payload interface{}) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if endpoint := lifx.returnAPIEndpoint(SetStatesEndpoint); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Put(endpoint, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

// ActivateScene activates a scene from your LIFX account
func ActivateScene(sceneUUID string, payload interface{}) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if endpoint := lifx.returnAPIEndpoint(ActivateSceneEndpoint, sceneUUID); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Put(endpoint, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

// Cycle makes the light(s) cycle to the next or previous state in a list of states
func Cycle(selector string, payload interface{}) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := lifx.returnAPIEndpoint(CycleEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Post(client, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

// PulseEffect performs a pulse effect by quickly flashing between the given colors
func PulseEffect(selector string, payload interface{}) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := lifx.returnAPIEndpoint(PulseEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Post(endpoint, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

// BreatheEffect performs a breathe effect by slowly fading between the given colors.
func BreatheEffect(selector string, payload interface{}) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := lifx.returnAPIEndpoint(BreatheEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Post(endpoint, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

// TogglePower turns off lights if any of them are on, or turns them on if they are all off.
func TogglePower(selector string) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := lifx.returnAPIEndpoint(ToggleEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Post(endpoint, nil)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

// StateDelta changes the state of the lights by the amount specified
func StateDelta(selector string, payload interface{}) (Response, error) {
	var (
		body     []byte
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := lifx.returnAPIEndpoint(StateDeltaEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = lifx.Post(endpoint, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

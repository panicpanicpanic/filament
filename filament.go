package filament

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	lifxAPIVersion = "v1"
)

type Client struct {
	token   string
	version string
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("token can't be empty")
	}

	return &Client{
		token:   token,
		version: lifxAPIVersion,
	}, nil
}

// GetLights returns []Device that belong to your LIFX account
func (c *Client) GetLights(selector string) ([]Device, error) {
	return getLights(selector)
}

// GetScenes returns []Scene that belong to your LIFX account
func (c *Client) GetScenes() ([]Scene, error) {
	var (
		body     []byte
		endpoint string
		err      error
		scenes   []Scene
	)

	if endpoint := returnAPIEndpoint(GetScenesEndpoint, ""); endpoint == "" {
		return scenes, errors.New("not a valid endpoint")
	}

	body, err = get(endpoint)
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
func (c *Client) ValidateColor(color string) (Color, error) {
	var (
		body        []byte
		deviceColor Color
		endpoint    string
		err         error
	)

	if endpoint := returnAPIEndpoint(ValidateColorEndpoint, color); endpoint == "" {
		return deviceColor, errors.New("not a valid endpoint")
	}

	body, err = get(endpoint)
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
func (c *Client) SetState(selector string, payload any) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := returnAPIEndpoint(SetStateEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = put(endpoint, payload)
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
func (c *Client) SetStates(payload any) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if endpoint := returnAPIEndpoint(SetStatesEndpoint, ""); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = put(endpoint, payload)
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
func (c *Client) ActivateScene(sceneUUID string, payload any) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if endpoint := returnAPIEndpoint(ActivateSceneEndpoint, sceneUUID); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = put(endpoint, payload)
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
func (c *Client) Cycle(selector string, payload any) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := returnAPIEndpoint(CycleEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = post(endpoint, payload)
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
func (c *Client) PulseEffect(selector string, payload any) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := returnAPIEndpoint(PulseEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = post(endpoint, payload)
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
func (c *Client) BreatheEffect(selector string, payload any) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if endpoint := returnAPIEndpoint(BreatheEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = post(endpoint, payload)
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
func (c *Client) TogglePower(selector string) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := returnAPIEndpoint(ToggleEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = post(endpoint, nil)
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
func (c *Client) StateDelta(selector string, payload any) (Response, error) {
	var (
		body     []byte
		endpoint string
		err      error
		response Response
	)

	if selector == "" {
		selector = "all"
	}

	if endpoint := returnAPIEndpoint(StateDeltaEndpoint, selector); endpoint == "" {
		return response, errors.New("not a valid endpoint")
	}

	body, err = post(endpoint, payload)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

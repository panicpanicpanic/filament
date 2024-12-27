package filament

import (
	"errors"
	"fmt"
)

const (
	apiEndpoint    = "https://api.lifx.com/"
	lifxAPIVersion = "v1"
)

type Client struct {
	token   string
	version string
	apiURL  string
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("token can't be empty")
	}

	return &Client{
		token:   token,
		version: lifxAPIVersion,
		apiURL:  fmt.Sprintf("%s%s", apiEndpoint, lifxAPIVersion),
	}, nil
}

// GetLights returns []Device that belong to your LIFX account
func (c *Client) GetLights(selector string) ([]Device, error) {
	return c.getLights(selector)
}

// GetScenes returns []Scene that belong to your LIFX account
func (c *Client) GetScenes() ([]Scene, error) {
	return c.getScenes()
}

// ValidateColor returns a Color if a valid color string is passed
func (c *Client) ValidateColor(color string) (Color, error) {
	return c.validateColor(color)
}

// SetState sets the state of the lights within the given selector, and returns a LIFX Response
func (c *Client) SetState(selector string, payload any) (Response, error) {
	return c.setState(selector, payload)
}

// SetStates sets multiple states across multiple selectors, and returns a LIFX Response
func (c *Client) SetStates(payload any) (Response, error) {
	return c.setStates(payload)
}

// ActivateScene activates a scene from your LIFX account
func (c *Client) ActivateScene(sceneUUID string, payload any) (Response, error) {
	return c.activateScene(sceneUUID, payload)
}

// Cycle makes the light(s) cycle to the next or previous state in a list of states
func (c *Client) Cycle(selector string, payload any) (Response, error) {
	return c.cycle(selector, payload)
}

// PulseEffect performs a pulse effect by quickly flashing between the given colors
func (c *Client) PulseEffect(selector string, payload any) (Response, error) {
	return c.pulseEffect(selector, payload)
}

// BreatheEffect performs a breathe effect by slowly fading between the given colors.
func (c *Client) BreatheEffect(selector string, payload any) (Response, error) {
	return c.breatheEffect(selector, payload)
}

// TogglePower turns off lights if any of them are on, or turns them on if they are all off.
func (c *Client) TogglePower(selector string) (Response, error) {
	return c.togglePower(selector)
}

// StateDelta changes the state of the lights by the amount specified
func (c *Client) StateDelta(selector string, payload any) (Response, error) {
	return c.stateDelta(selector, payload)
}

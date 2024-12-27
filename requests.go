package filament

import (
	"encoding/json"
	"fmt"
)

func (c *Client) getLights(selector string) ([]Device, error) {
	endpoint := c.returnAPIEndpoint(getLightsEndpoint, selector)

	body, err := get(endpoint, c.token)
	if err != nil {
		return []Device{}, fmt.Errorf(err.Error())
	}

	var devices []Device
	if err = json.Unmarshal(body, &devices); err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	return devices, nil
}

func (c *Client) getScenes() ([]Scene, error) {
	endpoint := c.returnAPIEndpoint(getScenesEndpoint, "")

	body, err := get(endpoint, c.token)
	if err != nil {
		return []Scene{}, fmt.Errorf(err.Error())
	}

	var scenes []Scene
	if err = json.Unmarshal(body, &scenes); err != nil {
		return scenes, fmt.Errorf(err.Error())
	}

	return scenes, nil
}

func (c *Client) validateColor(color string) (Color, error) {
	endpoint := c.returnAPIEndpoint(validateColorEndpoint, color)

	body, err := get(endpoint, c.token)
	if err != nil {
		return Color{}, fmt.Errorf(err.Error())
	}

	var deviceColor Color
	if err = json.Unmarshal(body, &deviceColor); err != nil {
		return deviceColor, fmt.Errorf(err.Error())
	}

	return deviceColor, nil
}

func (c *Client) setState(selector string, payload any) (Response, error) {
	endpoint := c.returnAPIEndpoint(setStateEndpoint, selector)

	body, err := put(endpoint, c.token, payload)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

func (c *Client) setStates(payload any) (Response, error) {
	endpoint := c.returnAPIEndpoint(setStatesEndpoint, "")

	body, err := put(endpoint, c.token, payload)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

func (c *Client) activateScene(sceneUID string, payload any) (Response, error) {
	endpoint := c.returnAPIEndpoint(activateSceneEndpoint, sceneUID)

	body, err := put(endpoint, c.token, payload)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

func (c *Client) cycle(selector string, payload any) (Response, error) {
	endpoint := c.returnAPIEndpoint(cycleEndpoint, selector)

	body, err := post(endpoint, c.token, payload)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

func (c *Client) pulseEffect(selector string, payload any) (Response, error) {
	endpoint := c.returnAPIEndpoint(pulseEndpoint, selector)

	body, err := post(endpoint, c.token, payload)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

func (c *Client) breatheEffect(selector string, payload any) (Response, error) {
	endpoint := c.returnAPIEndpoint(breatheEndpoint, selector)

	body, err := post(endpoint, c.token, payload)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

func (c *Client) togglePower(selector string) (Response, error) {
	endpoint := c.returnAPIEndpoint(toggleEndpoint, selector)

	body, err := post(endpoint, c.token, nil)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

func (c *Client) stateDelta(selector string, payload any) (Response, error) {
	endpoint := c.returnAPIEndpoint(stateDeltaEndpoint, selector)

	body, err := post(endpoint, c.token, payload)
	if err != nil {
		return Response{}, fmt.Errorf(err.Error())
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return response, fmt.Errorf(err.Error())
	}

	return response, nil
}

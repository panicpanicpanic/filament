package filament

import (
	"encoding/json"
	"fmt"
)

func getLights(selector string) ([]Device, error) {
	endpoint := returnAPIEndpoint(GetLightsEndpoint, selector)

	body, err := get(endpoint)
	if err != nil {
		return []Device{}, fmt.Errorf(err.Error())
	}

	var devices []Device
	if err = json.Unmarshal(body, &devices); err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	return devices, nil
}

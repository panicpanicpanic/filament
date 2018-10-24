package filament

import (
	"fmt"

	"github.com/panicpanicpanic/filament/service"
)

//Filament does stuff
type Filament struct {
}

// GetLights returns []Device's that belong to your LIFX account
func (f *Filament) GetLights(client *service.LIFXReq) ([]service.Device, error) {
	var devices []service.Device
	var err error
	var service service.Service

	devices, err = service.LIFXGetRequest(client)
	if err != nil {
		return devices, fmt.Errorf(err.Error())
	}

	return devices, nil
}
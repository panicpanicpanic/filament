package device

import "time"

type DevicePayload struct {
	Devices []Device
}

type Device struct {
	ID               string         `json:"json"`
	UUID             string         `json:"uuid"`
	Label            string         `json:"label"`
	Connected        bool           `json:"connected"`
	Power            string         `json:"power"`
	Brightness       int            `json:"brightness"`
	LastSeen         time.Time      `json:"last_seen"`
	SecondsSinceSeen int            `json:"seconds_since_seen"`
	Group            DeviceGroup    `json:"group"`
	Color            DeviceColor    `json:"color"`
	Location         DeviceLocation `json:"location"`
	Product          DeviceProduct  `json:"product"`
}

type DeviceGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeviceColor struct {
	Hue        int `json:"hue"`
	Saturation int `json:"saturation"`
	Kelvin     int `json:"kelvin"`
}

type DeviceLocation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type DeviceProduct struct {
	Name         string             `json:"name"`
	Identifier   string             `json:"identifier"`
	Company      string             `json:"company"`
	Capabilities DeviceCapabilities `json:"capabilities"`
}

type DeviceCapabilities struct {
	HasColor             bool `json:"has_color"`
	HasVariableColorTemp bool `json:"has_variable_color_temp"`
	HasIR                bool `json:"has_ir"`
	HasChain             bool `json:"has_chain"`
	HasMultizone         bool `json:"has_multizone"`
	MinKelvin            int  `json:"min_kelvin"`
	MaxKelvin            int  `json:"max_kelvin"`
}

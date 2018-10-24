package filament

import (
	"time"
)

// Device represents the core fields for LIFX lightbulb info
type Device struct {
	ID               string         `json:"json"`
	UUID             string         `json:"uuid"`
	Label            string         `json:"label"`
	Connected        bool           `json:"connected"`
	Power            string         `json:"power"`
	Brightness       float64        `json:"brightness"`
	LastSeen         time.Time      `json:"last_seen"`
	SecondsSinceSeen int            `json:"seconds_since_seen"`
	Group            DeviceGroup    `json:"group"`
	Color            DeviceColor    `json:"color"`
	Location         DeviceLocation `json:"location"`
	Product          DeviceProduct  `json:"product"`
}

// DeviceGroup represents which group a Device belongs to
type DeviceGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DeviceColor represents which colors a Device is currently set to
type DeviceColor struct {
	Hue        float64 `json:"hue"`
	Saturation float64 `json:"saturation"`
	Kelvin     float64 `json:"kelvin"`
}

// DeviceLocation represents what Location a Device belongs to
type DeviceLocation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// DeviceProduct represents the type of LIFX product a Device is
type DeviceProduct struct {
	Name         string             `json:"name"`
	Identifier   string             `json:"identifier"`
	Company      string             `json:"company"`
	Capabilities DeviceCapabilities `json:"capabilities"`
}

// DeviceCapabilities represents all of the current features a Device has
type DeviceCapabilities struct {
	HasColor             bool    `json:"has_color"`
	HasVariableColorTemp bool    `json:"has_variable_color_temp"`
	HasIR                bool    `json:"has_ir"`
	HasChain             bool    `json:"has_chain"`
	HasMultizone         bool    `json:"has_multizone"`
	MinKelvin            float64 `json:"min_kelvin"`
	MaxKelvin            float64 `json:"max_kelvin"`
}

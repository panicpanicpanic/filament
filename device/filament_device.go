package device

import (
	"time"
)

// Device represents the core fields for a LIFX light source
type Device struct {
	ID               string    `json:"json"`
	UUID             string    `json:"uuid"`
	Label            string    `json:"label"`
	Connected        bool      `json:"connected"`
	Power            string    `json:"power"`
	Brightness       float64   `json:"brightness"`
	LastSeen         time.Time `json:"last_seen"`
	SecondsSinceSeen int       `json:"seconds_since_seen"`
	Group            Group     `json:"group"`
	Color            Color     `json:"color"`
	Location         Location  `json:"location"`
	Product          Product   `json:"product"`
}

// Group represents which group a Device belongs to
type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Color represents which colors a Device is currently set to
type Color struct {
	Hue        float64 `json:"hue"`
	Saturation float64 `json:"saturation"`
	Kelvin     float64 `json:"kelvin"`
	Name       string  `json:"name"`
}

// Location represents what Location a Device belongs to
type Location struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Product represents the type of LIFX product a Device is
type Product struct {
	Name         string       `json:"name"`
	Identifier   string       `json:"identifier"`
	Company      string       `json:"company"`
	Capabilities Capabilities `json:"capabilities"`
}

// Capabilities represents all of the current features a Device has
type Capabilities struct {
	HasColor             bool    `json:"has_color"`
	HasVariableColorTemp bool    `json:"has_variable_color_temp"`
	HasIR                bool    `json:"has_ir"`
	HasChain             bool    `json:"has_chain"`
	HasMultizone         bool    `json:"has_multizone"`
	MinKelvin            float64 `json:"min_kelvin"`
	MaxKelvin            float64 `json:"max_kelvin"`
}

// Scene represents currently available scenes for your LIFX bulb
type Scene struct {
	UUID      string            `json:"uuid"`
	Name      string            `json:"name"`
	Account   map[string]string `json:"account"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
	States    []State           `json:"states"`
}

// State represents the states of a Scene
type State struct {
	Color      Color   `json:"color"`
	Selector   string  `json:"selector"`
	Power      string  `json:"power"`
	Fast       bool    `json:"fast"`
	Brightness float64 `json:"brightnness"`
	Duration   float64 `json:"duration"`
	Infared    float64 `json:"infared"`
}

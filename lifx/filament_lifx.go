package lifx

const (
	// LIFXAPIURL is the URL for the latest LIFX HTTP API
	LIFXAPIURL = "https://api.lifx.com/v1"
)

// Client contains the LIFX AccessToken and URL endpoint needed to reach the LIFX HTTP API
type Client struct {
	AccessToken string
	Endpoint    string
}

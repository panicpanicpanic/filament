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

// Response is a generic slice of results from LIFX API
type Response struct {
	Results []Result `json:"results"`
}

// Result returns ID, Status and Label from LIFX API
type Result struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Label  string `json:"label"`
}

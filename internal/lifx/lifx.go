package lifx

const (
	// APIEndpoint is the URL for the latest LIFX HTTP API
	APIEndpoint = "https://api.lifx.com/v1"
)

// Response is a generic slice of results from LIFX API
type Response struct {
	Results []struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Label  string `json:"label"`
	} `json:"results"`
}

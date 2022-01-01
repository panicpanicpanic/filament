package lifx

import "errors"

const (
	// APIEndpoint is the URL for the latest LIFX HTTP API
	APIEndpoint = "https://api.lifx.com/v1"
)

// Client contains the LIFX access token needed to authenticate
// against the LIFX HTTP API
type Client struct {
	accessToken string
}

// Response is a generic slice of results from LIFX API
type Response struct {
	Results []struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Label  string `json:"label"`
	} `json:"results"`
}

// NewClient accepts an access token and returns a Client
// If the token is empty, it'll also return an error
func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("token can't be empty")
	}

	return &Client{accessToken: token}, nil
}

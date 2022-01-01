package filament

import "errors"

// Client contains the LIFX access token needed to authenticate
// against the LIFX HTTP API
type Client struct {
	accessToken string
}

// NewClient accepts an access token and returns a Client
// If the token is empty, it'll also return an error
func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("token can't be empty")
	}

	return &Client{accessToken: token}, nil
}

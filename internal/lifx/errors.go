package lifx

import "errors"

var (
	MissingTokenEndpointError = errors.New("missing access token or valid API endpoint")
)

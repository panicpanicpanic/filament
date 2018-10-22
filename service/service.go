package service

// LIFXReq represents the request structure for reaching LIFX HTTP API
type LIFXReq struct {
	AccessToken string
	URL         string
}

// Service is being exported for access to rest of application
type Service struct{}

package service

const APIUrl = "https://api.lifx.com/v1"

type LIFXReq struct {
	AccessToken string
	URL         string
}

type Service struct{}

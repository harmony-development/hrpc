package v1

import "net/http"

type VoiceServiceClient struct {
	client    *http.Client
	serverURL string
}

func NewVoiceServiceClient(url string) *VoiceServiceClient {
	return &VoiceServiceClient{
		client:    &http.Client{},
		serverURL: url,
	}
}

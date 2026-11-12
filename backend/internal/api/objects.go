package api

import(
	lksdk "github.com/livekit/server-sdk-go/v2"
)

type Server struct {
	APIKey string
	APISecret string
	LiveKitURL string
	Rooms *lksdk.RoomServiceClient
}

type TokenRequest struct {
	Name string `json:"name"`
	Room string `json:"room"`
}

type TokenResponse struct {
	Token string `json:"token"`
	URL string `json:"url"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}


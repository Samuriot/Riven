package api

import (
	"net/http"
	lksdk "github.com/livekit/server-sdk-go/v2"
)

type Server struct {
	APIKey string
	APISecret string
	LiveKitURL string
	Rooms *lksdk.RoomServiceClient
}

func (s *Server) HandleToken(writer http.ResponseWriter, req *http.Request) {

}

func (s *Server) HandleWebhook(writer http.ResponseWriter, req *http.Request) {

}

func (s *Server) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", s.HandleToken)
	mux.HandleFunc("Post /api/livekit/webhook", s.HandleWebhook)
	return mux
}
package api

import (
	"crypto/rand"
	"log"
	"net/http"
	"strings"

	"github.com/Samuriot/Riven/backend/internal/livekit"
)

var defaultRoom = "test"


func (s *Server) HandleToken(writer http.ResponseWriter, req *http.Request) {
	var request TokenRequest
	if err := DecodeJson(writer, req, &request); err != nil {
		WriteError(writer, http.StatusBadRequest, "Invalid JSON: " + err.Error())
		return
	}

	request.Name = strings.TrimSpace(request.Name)
	request.Room = strings.TrimSpace(request.Room)

	if request.Name == "" {
		WriteError(writer, http.StatusBadRequest, "Name is required")
		return
	}

	if len(request.Name) > 32 {
		WriteError(writer, http.StatusBadRequest, "Name is too long, try again")
		return
	}

	if request.Room == "" {
		request.Room = defaultRoom
	}

	identity := rand.Text()

	token, err := livekit.CreateToken(s.APIKey, s.APISecret, request.Room, identity, request.Name, false)

	if err != nil {
		log.Printf("create token: %v", err)
		WriteError(writer, http.StatusInternalServerError, "could not create a token")
		return
	}

	WriteJson(writer, http.StatusOK, TokenResponse{Token: token, URL: s.LiveKitURL})
}

func (s *Server) HandleWebhook(writer http.ResponseWriter, req *http.Request) {

}

func (s *Server) InitRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", s.HandleToken)
	mux.HandleFunc("POST /api/livekit/webhook", s.HandleWebhook)
	return mux
}
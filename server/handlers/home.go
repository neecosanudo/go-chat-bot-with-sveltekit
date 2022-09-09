package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/neecosanudo/go-chat-bot-with-sveltekit/server/server"
)

type HomeResponse struct {
	Message string `json:"message"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "1.21 Gigawatts!",
		})
	}
}

package handlers

import (
	"encoding/json"
	"github.com/segmentio/ksuid"
	"kevinPicon/go/rest-ws/models"
	"kevinPicon/go/rest-ws/repository"
	"kevinPicon/go/rest-ws/server"
	"net/http"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email: request.Email,
			Name:  request.Name,
			Id:    id.String(),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Name:  user.Name,
			Email: user.Email,
		})
	}

}

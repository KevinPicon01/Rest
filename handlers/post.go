package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
	"kevinPicon/go/rest-ws/models"
	"kevinPicon/go/rest-ws/repository"
	"kevinPicon/go/rest-ws/server"
	"net/http"
	"strings"
)

type InsertPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
type PostResponse struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func InsertPostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var postRequest = InsertPostRequest{}
			if err := json.NewDecoder(r.Body).Decode(&postRequest); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			id, err := ksuid.NewRandom()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			post := models.Post{
				Id:       id.String(),
				Title:    postRequest.Title,
				Content:  postRequest.Content,
				AuthorId: claims.UserId,
			}
			err = repository.InsertPost(r.Context(), &post)
			if err != nil {
				fmt.Println("hola")
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			var postResponse = PostResponse{
				Id:      post.Id,
				Title:   post.Title,
				Content: post.Content,
			}
			if err := json.NewEncoder(w).Encode(postResponse); err != nil {
				http.Error(w, err.Error()+"k", http.StatusInternalServerError)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

func GetPostByIdHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		post, err := repository.GetPostById(r.Context(), params["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		var postResponse = PostResponse{
			Id:      post.Id,
			Title:   post.Title,
			Content: post.Content,
		}
		if err := json.NewEncoder(w).Encode(postResponse); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func UpdatePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			var request = models.Post{}
			params := mux.Vars(r)
			err := json.NewDecoder(r.Body).Decode(&request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			request = models.Post{
				Id:       params["id"],
				Title:    request.Title,
				Content:  request.Content,
				AuthorId: claims.UserId,
			}
			err = repository.UpdatePost(r.Context(), &request)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}
func DeletePostHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
		tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
		token, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.Config().JWTSecret), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if claims, ok := token.Claims.(*models.AppClaims); ok && token.Valid {
			params := mux.Vars(r)
			var request = models.Post{}
			request = models.Post{
				Id:       params["id"],
				AuthorId: claims.UserId,
			}
			err = repository.DeletePost(r.Context(), &request)
			if err != nil {
				http.Error(w, err.Error()+"h", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

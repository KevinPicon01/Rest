package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"kevinPicon/go/rest-ws/Middleware"
	"kevinPicon/go/rest-ws/handlers"
	"kevinPicon/go/rest-ws/server"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	serv, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JWT_SECRET,
		DatabaseUrl: DATABASE_URL,
	})
	if err != nil {
		log.Fatal(err)
	}
	serv.Start(BindRouters)
}
func BindRouters(s server.Server, r *mux.Router) {
	//Middleware
	r.Use(Middleware.CheckAuthMiddleware(s))
	//Routes
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods("POST")
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods("GET")
	r.HandleFunc("/CreatePost", handlers.InsertPostHandler(s)).Methods("POST")
	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods("DELETE")
	r.HandleFunc("/posts", handlers.GetPostsHandler(s)).Methods("GET")
	//WebSockets
	r.HandleFunc("/ws", s.Hub().HandleWebSocket)

}

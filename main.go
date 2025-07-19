package main

import (
	"log"
	"manty-go/handler"
	"manty-go/resources"
	"manty-go/service"
	"net/http"

	"github.com/zbum/mantyboot/http/mux"
	"github.com/zbum/mantyboot/http/mux/middleware"
)

func main() {

	mux := mux.NewMantyMux()

	logger := log.Default()
	mux.AddMiddleware(middleware.AccessLogger(logger))
	mux.AddMiddleware(middleware.CORS(
		&middleware.CORSConfig{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
		},
	))

	mux.Handle("GET /", http.FileServerFS(resources.GetStaticFs()))

	indexHandler := handler.NewIndexHandler()
	mux.HandleFunc("GET /index", indexHandler.Index)

	postService := service.NewDumyPostService()
	postHandler := handler.NewPostHandler(logger, nil, postService)
	mux.HandleFunc("GET /posts", postHandler.GetPosts)

	// UserService 및 LoginHandler 추가
	userService := service.NewUserService()
	loginHandler := handler.NewLoginHandler(userService)
	mux.HandleFunc("POST /login", loginHandler.Login)
	mux.HandleFunc("OPTIONS /login", loginHandler.Login)

	log.Printf("Starting server on port : %s", ":9090")

	log.Panic(http.ListenAndServe(":9090", mux))

}

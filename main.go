package main

import (
	"log"
	"manty-go/handler"
	"manty-go/resources"
	"manty-go/service"
	"manty-go/utils"
	"net/http"

	"github.com/zbum/mantyboot/http/mux"
	"github.com/zbum/mantyboot/http/mux/middleware"
)

func main() {

	mux := mux.NewMantyMux()

	logger := log.Default()
	mux.AddMiddleware(middleware.AccessLogger(logger))
	mux.AddMiddleware(utils.CorsAllowAll())

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

	log.Panic(http.ListenAndServe(":9090", mux))

}

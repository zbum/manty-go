package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"manty-go/service"
	"net/http"
)

type PostHandler struct {
	logger      *log.Logger
	db          *sql.DB
	postService *service.DumyPostService
}

func NewPostHandler(logger *log.Logger, db *sql.DB, postService *service.DumyPostService) *PostHandler {
	return &PostHandler{
		logger:      logger,
		db:          db,
		postService: postService,
	}
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {

	posts := h.postService.GetPosts()

	json.NewEncoder(w).Encode(posts)
}

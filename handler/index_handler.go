package handler

import "net/http"

type IndexHandler struct {
}

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{}
}

func (h *IndexHandler) Index(w http.ResponseWriter, r *http.Request) {

}

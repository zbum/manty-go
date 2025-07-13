package service

import (
	"manty-go/model"
	"time"
)

type DumyPostService struct {
}

func NewDumyPostService() *DumyPostService {
	return &DumyPostService{}
}

func (s *DumyPostService) GetPosts() []model.Post {
	return []model.Post{
		{
			ID:        1,
			Title:     "Hello, World!",
			Content:   "This is a test post.",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "Hello, World!",
			Content:   "This is a test post.",
			CreatedAt: time.Now(),
		},
	}
}

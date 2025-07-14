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
			Title:     "loki-stack 을 사용하면서 label 을 추가하고 싶으면.",
			Content:   "This is a test post.This is a test post.This is a test post.This is a test post.This is a test post.This is a test post.This is a test post.This is a test post.This is a test post.This is a test post.",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Title:     "go 1.22 버전에서 for 루프의 2가지 변경 내용",
			Content:   "This is a test post.",
			CreatedAt: time.Now(),
		},
	}
}

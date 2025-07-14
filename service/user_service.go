package service

import "manty-go/model"

// 임시 사용자 목록 (실제 서비스에서는 DB 사용)
var users = []model.User{
	{ID: 1, Username: "admin", Password: "password"},
	{ID: 2, Username: "user", Password: "1234"},
}

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// username, password가 맞으면 User 반환, 아니면 nil
func (s *UserService) Authenticate(username, password string) *model.User {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return &u
		}
	}
	return nil
}

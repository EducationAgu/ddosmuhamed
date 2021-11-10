package interfaces

import "backend/service/model"

type Service interface {
	User() UserService
	TokenManager() TokenManagerService
	Account() AccountService
}

type TokenManagerService interface {
	GenerateToken(user model.User) (string, error)
	Parse(accessToken string) (model.User, error)
	NewRefreshToken() (string, error)
}

type UserService interface {
	Register(user model.User) (model.Tokens, error)
	Login(user model.User) (model.Tokens, error)
	RefreshToken(token string) (model.Tokens, error)
}

type AccountService interface {
	Get(id int64) (model.User, error)
	GetStudents(id int64) ([]model.User, error)
}

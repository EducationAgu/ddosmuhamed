package interfaces

import (
	"backend/service/model"
)

type Provider interface {
	User() UserProvider
	Session() SessionProvider
	Account() AccountProvider
}

type UserProvider interface {
	Register(user model.User) (int64, error)
	GetByCredentials(model.User) (user model.User, err error)
}

type SessionProvider interface {
	Create(token string, userId int64) error
	Refresh(oldToken, newToken string) error
	GetUserByToken(token string) (model.User, error)
}

type AccountProvider interface {
	Get(int64) (model.User, error)
	GetStudents(int64) ([]model.User, error)
}

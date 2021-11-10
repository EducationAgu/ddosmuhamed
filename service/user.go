package service

import (
	"crypto/sha1"
	"fmt"

	"backend/internal/interfaces"
	"backend/service/model"
)

var _ interfaces.UserService = (*User)(nil)

type User struct {
	db   interfaces.Provider
	salt string
	interfaces.TokenManagerService
}

func NewUser(db interfaces.Provider, manager interfaces.TokenManagerService) *User {
	return &User{
		db:                  db,
		salt:                "salt",
		TokenManagerService: manager,
	}
}

func (u *User) Login(user model.User) (tokens model.Tokens, err error) {
	user.Password = u.hash(user.Password)
	user, err = u.db.User().GetByCredentials(user)
	if err != nil {
		return model.Tokens{}, err
	}

	tokens.AccessToken, err = u.TokenManagerService.GenerateToken(user)
	if err != nil {
		return model.Tokens{}, err
	}

	tokens.RefreshToken, err = u.TokenManagerService.NewRefreshToken()
	if err != nil {
		return model.Tokens{}, err
	}

	err = u.db.Session().Create(tokens.RefreshToken, user.Id)
	if err != nil {
		return model.Tokens{}, err
	}

	return tokens, err
}

func (u *User) RefreshToken(token string) (model.Tokens, error) {
	user, err := u.db.Session().GetUserByToken(token)
	if err != nil {
		return model.Tokens{}, err
	}
	var tokens model.Tokens
	tokens.AccessToken, err = u.GenerateToken(user)
	if err != nil {
		return model.Tokens{}, err
	}
	tokens.RefreshToken, err = u.NewRefreshToken()
	if err != nil {
		return model.Tokens{}, err
	}

	err = u.db.Session().Refresh(token, tokens.RefreshToken)
	if err != nil {
		return model.Tokens{}, err
	}
	return tokens, nil
}

func (u *User) hash(in string) string {
	hash := sha1.New()
	hash.Write([]byte(in))
	in = fmt.Sprintf("%x", hash.Sum([]byte(u.salt)))
	return in
}

func (u *User) Register(user model.User) (tokens model.Tokens, err error) {
	user.Password = u.hash(user.Password)
	user.Id, err = u.db.User().Register(user)
	if err != nil {
		return model.Tokens{}, err
	}
	tokens.AccessToken, err = u.GenerateToken(user)
	if err != nil {
		return model.Tokens{}, err
	}

	tokens.RefreshToken, err = u.NewRefreshToken()
	if err != nil {
		return model.Tokens{}, err
	}
	err = u.db.Session().Create(tokens.RefreshToken, user.Id)
	if err != nil {
		return tokens, err
	}
	return tokens, nil
}

package data

import (
	"backend/data/entity"
	"backend/internal/interfaces"
	"backend/service/model"

	"github.com/go-pg/pg/v10"
)

var _ interfaces.UserProvider = (*User)(nil)

type User struct {
	*pg.DB
}

func NewUser(db *pg.DB) *User {
	return &User{db}
}

func (u *User) Register(user model.User) (int64, error) {
	ent := entity.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	_, err := u.DB.Model(&ent).Insert()
	if err != nil {
		return 0, err
	}
	return ent.Id, nil
}

func (u *User) GetByCredentials(in model.User) (user model.User, err error) {
	var ent entity.User
	err = u.DB.Model(&ent).
		Where("username = ? AND password = ?",
			in.Username, in.Password).
		Select()
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Id:       ent.Id,
		Username: ent.Username,
		Name:     ent.Name,
		Email:    ent.Email,
	}, nil
}

package data

import (
	"errors"

	"backend/data/entity"
	"backend/internal/interfaces"
	"backend/service/model"

	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
)

var _ interfaces.UserProvider = (*User)(nil)

type User struct {
	*pg.DB
	salt int
}

func NewUser(db *pg.DB, salt int) *User {
	return &User{DB: db, salt: salt}
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
		Where("username = ?",
			in.Username, in.Password).
		Select()
	if err != nil {
		return model.User{}, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(in.Password), []byte(ent.Password)); err != nil {
		return model.User{}, errors.New("Invalid password")
	}

	return model.User{
		Id:       ent.Id,
		Username: ent.Username,
		Name:     ent.Name,
		Email:    ent.Email,
	}, nil
}

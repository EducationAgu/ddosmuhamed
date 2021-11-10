package data

import (
	"backend/data/entity"
	"backend/internal/interfaces"
	"backend/service/model"

	"github.com/go-pg/pg/v10"
)

type AccountProvider struct {
	*pg.DB
}

func NewAccount(db *pg.DB) *AccountProvider {
	return &AccountProvider{db}
}

var _ interfaces.AccountProvider = (*AccountProvider)(nil)

func (a *AccountProvider) Get(id int64) (model.User, error) {
	var user entity.User
	err := a.Model(&user).Where("id = ?", id).Select()
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (a *AccountProvider) GetStudents(id int64) ([]model.User, error) {
	var user []entity.User

	err := a.Model(&user).Where("p_id = ?", id).Select()
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0, len(user))
	for _, item := range user {
		users = append(users, model.User{
			Id:       item.Id,
			Name:     item.Name,
			Username: item.Username,
			Email:    item.Email,
			Password: item.Password,
		})
	}
	return users, nil
}

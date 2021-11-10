package service

import (
	"backend/internal/interfaces"
	"backend/service/model"
)

type AccountService struct {
	interfaces.Provider
}

func NewAccount(provider interfaces.Provider) *AccountService {
	return &AccountService{provider}
}

var _ interfaces.AccountService = (*AccountService)(nil)

func (a *AccountService) Get(id int64) (model.User, error) {
	return a.Provider.Account().Get(id)
}

func (a *AccountService) GetStudents(id int64) ([]model.User, error) {
	return a.Provider.Account().GetStudents(id)
}

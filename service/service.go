package service

import "backend/internal/interfaces"

type Service struct {
	interfaces.UserService
	interfaces.TokenManagerService
	interfaces.AccountService
}

func New(db interfaces.Provider, salt int) *Service {
	tokenManager := NewManager("key")
	return &Service{
		UserService:         NewUser(db, tokenManager, salt),
		TokenManagerService: tokenManager,
		AccountService:      NewAccount(db),
	}
}

func (s *Service) User() interfaces.UserService {
	return s.UserService
}

func (s *Service) TokenManager() interfaces.TokenManagerService {
	return s.TokenManagerService
}

func (s *Service) Account() interfaces.AccountService {
	return s.AccountService
}

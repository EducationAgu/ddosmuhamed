package data

import (
	"time"

	"backend/data/entity"
	"backend/internal/interfaces"
	"backend/internal/utils"
	"backend/service/model"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
)

var _ interfaces.SessionProvider = (*Session)(nil)

type Session struct {
	*pg.DB
}

func NewSession(db *pg.DB) *Session {
	return &Session{db}
}

func (s *Session) Create(token string, userId int64) error {
	session := entity.Session{
		UserId:    userId,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}
	_, err := s.Model(&session).Insert()

	go func() {
		_, err = s.Exec(`DELETE FROM sessions WHERE expires_at < now() and user_id = ?`, userId)
		logrus.Error(err)
	}()
	return err
}

func (s *Session) Refresh(oldToken, newToken string) error {
	var token entity.Session
	token.Token = newToken
	token.ExpiresAt = time.Now().Add(time.Minute * 5)
	_, err := s.Model(&token).Column("token", "expires_at").Where("token = ?", oldToken).Update()
	return err
}

func (s *Session) GetUserByToken(token string) (model.User, error) {
	var session entity.Session
	err := s.Model(&session).Where("token = ?", token).Select()
	if err != nil {
		return model.User{}, err
	}

	if time.Now().After(session.ExpiresAt) {
		s.Model(&session).Delete()
		return model.User{}, utils.TokenExpired
	}

	var user entity.User

	err = s.Model(&user).Where("id = ?", session.UserId).Select()
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

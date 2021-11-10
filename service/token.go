package service

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"backend/service/model"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
)

type TokenManager struct {
	signingKey string
}

func NewManager(signingKey string) *TokenManager {
	return &TokenManager{signingKey: signingKey}
}

func (m *TokenManager) GenerateToken(user model.User) (string, error) {
	data, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
		Subject:   string(data),
	})

	return token.SignedString([]byte(m.signingKey))
}

func (m *TokenManager) Parse(accessToken string) (model.User, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(m.signingKey), nil
	})
	if err != nil {
		return model.User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return model.User{}, fmt.Errorf("error get user claims from token")
	}
	var user model.User
	err = json.Unmarshal([]byte(claims["sub"].(string)), &user)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (m *TokenManager) NewRefreshToken() (string, error) {
	b := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if _, err := r.Read(b); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", b), nil
}

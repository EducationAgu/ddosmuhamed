package server

import (
	"encoding/json"
	"net/http"

	"backend/internal/interfaces"

	"github.com/sirupsen/logrus"
)

type Account struct {
	interfaces.Service
}

func (a *Account) profile(writer http.ResponseWriter, req *http.Request) {
	user, err := auth(a.TokenManager(), req)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte(err.Error()))
		return
	}
	user, err = a.Service.Account().Get(user.Id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		logrus.Error(err, "1")
		writer.Write([]byte(err.Error()))
		return
	}
	body, _ := json.Marshal(user)
	writer.Write(body)
}

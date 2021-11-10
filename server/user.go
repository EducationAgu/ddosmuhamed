package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"backend/internal/interfaces"
	"backend/service/model"

	"github.com/sirupsen/logrus"
)

type User struct {
	interfaces.Service
}

func (u *User) signup(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		logrus.Error(err)
		return
	}

	tokens, err := u.User().Register(user)
	if err != nil {
		logrus.Error(err)
		return
	}
	payload, err := json.Marshal(tokens)
	if err != nil {
		logrus.Error(err)
		return
	}

	writer.Write(payload)
	writer.WriteHeader(http.StatusOK)
}

func (u *User) signin(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logrus.Error(err)
		return
	}
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		logrus.Error(err)
		return
	}

	tokens, err := u.User().Login(user)
	if err != nil {
		logrus.Error(err)
		return
	}

	body, _ = json.Marshal(tokens)

	writer.Write(body)
}

func (u *User) refresh(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logrus.Error(err)
		return
	}

	var tokenIn model.Tokens

	err = json.Unmarshal(body, &tokenIn)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		logrus.Error(err, "2")
		return
	}

	tokens, err := u.User().RefreshToken(tokenIn.RefreshToken)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		logrus.Error(err, "1")
		writer.Write([]byte("Сессия просрочена!"))
		return
	}

	body, _ = json.Marshal(tokens)
	writer.Write(body)
}

func (u *User) ping(writer http.ResponseWriter, request *http.Request) {

}

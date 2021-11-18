package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"backend/internal/interfaces"
	"backend/service/model"

	"github.com/sirupsen/logrus"
)

type User struct {
	interfaces.Service
	Salt int
}

func (u *User) signup(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokens, err := u.User().Register(user)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	payload, err := json.Marshal(tokens)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Write(payload)
	writer.WriteHeader(http.StatusOK)
}

func (u *User) signin(writer http.ResponseWriter, request *http.Request) {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var user model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokens, err := u.User().Login(user)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	body, err = json.Marshal(tokens)
	if err != nil {
		logrus.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

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

func (u *User) salt(writer http.ResponseWriter, _ *http.Request) {
	writer.Write([]byte(strconv.Itoa(u.Salt)))
}

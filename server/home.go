package server

import (
	"encoding/json"
	"net/http"

	"backend/internal/interfaces"
	"backend/service/model"

	"github.com/sirupsen/logrus"
)

type Home struct {
	interfaces.Service
}

func (h *Home) home(writer http.ResponseWriter, req *http.Request) {
	s, e := auth(h.TokenManager(), req)
	if e != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		logrus.Error(e)
		return
	}
	stud, e := h.Service.Account().GetStudents(s.Id)
	if e != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		logrus.Error(e)
		return
	}

	output := struct {
		Name     string
		Students []model.User
	}{
		Name:     s.Username,
		Students: stud,
	}

	msg, _ := json.Marshal(output)

	writer.Write(msg)
}

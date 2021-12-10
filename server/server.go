package server

import (
	"errors"
	"log"
	"net/http"

	"backend/internal/interfaces"
	"backend/service/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	service interfaces.Service
	user    User
	home    Home
	account Account
}

func New(service interfaces.Service, salt int) *Server {
	return &Server{
		user:    User{Service: service, Salt: salt},
		home:    Home{service},
		account: Account{service},
	}
}
func (s *Server) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/api/", s.home.home)
	r.HandleFunc("/api/auth/signup", s.user.signup)
	r.HandleFunc("/api/auth/salt", s.user.salt)
	r.HandleFunc("/api/auth/signin", s.user.signin)
	r.HandleFunc("/api/auth/refreshtoken", s.user.refresh)
	r.HandleFunc("/api/user", s.account.profile)

	log.Println("Начинаем работу")
	log.Fatal(http.ListenAndServe(":500", // dir+"/certs/localhost.crt", dir+"/certs/localhost.key",
		handlers.CORS(
			handlers.AllowedHeaders(
				[]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r)))
}

func auth(tokenManager interfaces.TokenManagerService, request *http.Request) (model.User, error) {
	header := request.Header.Get("Authorization")
	if header == "" {
		return model.User{}, errors.New("empty auth header")
	}

	return tokenManager.Parse(header)
}

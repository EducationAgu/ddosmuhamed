package main

import (
	"backend/data"
	"backend/server"
	"backend/service"

	"github.com/go-pg/pg/v10"
)

func main() {
	// docker run -d --name go-auth -e POSTGRES_PASSWORD=golang -e  POSTGRES_USER=golang -e POSTGRES_DB=golang --restart always -p "501:5432" postgres
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:501",
		User:     "golang",
		Password: "golang",
		Database: "golang",
	})

	salt := 10
	provider := data.New(db, salt)
	srv := service.New(provider, salt)

	server.New(srv, salt).Start()
}

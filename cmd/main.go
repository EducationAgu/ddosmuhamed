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
		Addr:     "postgres:5432",
		User:     "golang",
		Password: "golang",
		Database: "golang",
	})

	provider := data.New(db)
	service := service.New(provider)

	server.New(service).Start()
}

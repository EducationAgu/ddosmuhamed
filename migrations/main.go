package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/spf13/pflag"
)

func main() {

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:501",
		User:     "golang",
		Password: "golang",
		Database: "golang",
	})
	var err error

	oldVersion, newVersion, err := migrations.Run(db, pflag.Args()...)
	if err != nil {
		if strings.Contains(err.Error(), "\"gopg_migrations\" does not exist") {
			args := []string{"init"}

			_, _, err = migrations.Run(db, args...)
			if err != nil {
				panic(err)
			}
		}

		oldVersion, newVersion, err = migrations.Run(db, pflag.Args()...)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Printf("Ошибка при накатке миграции! %s", err.Error())
			os.Exit(1)
		}
	}

	fmt.Println(oldVersion, " -> ", newVersion)
}

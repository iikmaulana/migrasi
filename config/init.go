package config

import (
	"fmt"
	"github.com/iikmaulana/gateway/controller"
	"github.com/iikmaulana/gateway/service"
	"github.com/jmoiron/sqlx"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

type AppConfig interface {
	Init() Config
}

type Config struct {
	Version   string
	DBRething *r.Session
	DB        *sqlx.DB
	Registry  *controller.Registry
	Server    *service.Server
	Gateway   *service.Service
}

func Init() Config {
	var cfg Config

	errx := cfg.InitCockroachdb()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitRethinkDB()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitMigrate()
	if errx != nil {
		errx.Panic()
	}

	fmt.Println("Server is running ..")
	return cfg
}

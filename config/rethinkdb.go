package config

import (
	"github.com/iikmaulana/gateway/libs"
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (cfg *Config) InitRethinkDB() serror.SError {

	db, err := r.Connect(r.ConnectOpts{
		Address:  helper.Env(libs.RethinkDBHost, "127.0.0.1:28015"),
		Database: helper.Env(libs.RethinkDBName, "test_golang"),
	})

	if err != nil {
		return serror.NewFromError(err)
	}

	db.SetMaxIdleConns(int(helper.StringToInt(helper.Env(libs.DBConnMaxIdle, "5"), 5)))
	db.SetMaxOpenConns(int(helper.StringToInt(helper.Env(libs.DBConnMaxOpen, "0"), 0)))

	cfg.DBRething = db

	return nil
}

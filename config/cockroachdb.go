package config

import (
	"fmt"
	"github.com/gearintellix/u2"
	"github.com/iikmaulana/gateway/libs"
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/jmoiron/sqlx"
	"time"
)

func (cfg *Config) InitCockroachdb() serror.SError {
	sqlConn := helper.Env(libs.DBConnStr, `
        host=__host__
		port=__port__
        user=__user__
		port=__port__
        password=__pwd__
        dbname=__name__
        sslmode=__sslMode__
        application_name=__appKey__
    `)
	sqlConn = u2.Binding(sqlConn, map[string]string{
		"host":    helper.Env(libs.DBHost, "127.0.0.1"),
		"user":    helper.Env(libs.DBUser, "root"),
		"pwd":     helper.Env(libs.DBPwd, ""),
		"name":    helper.Env(libs.DBName, "asgard"),
		"sslMode": helper.Env(libs.DBSSLMode, "disable"),
		"appKey":  helper.Env(libs.AppKey, "device"),
		"appName": helper.Env(libs.AppName, "Device"),
		"port":    helper.Env(libs.DBPort, "54320"),
	})

	db, err := sqlx.Connect(helper.Env(libs.DBEngine, "impl"), sqlConn)
	if err != nil {
		return serror.NewFromErrorc(err, fmt.Sprintf("failed connect to database %s", helper.Env(libs.DBName, "asgard")))
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(helper.StringToInt(helper.Env(libs.DBConnLifetime, "15"), 15)))
	db.SetMaxIdleConns(int(helper.StringToInt(helper.Env(libs.DBConnMaxIdle, "5"), 5)))
	db.SetMaxOpenConns(int(helper.StringToInt(helper.Env(libs.DBConnMaxOpen, "0"), 0)))

	cfg.DB = db

	return nil
}

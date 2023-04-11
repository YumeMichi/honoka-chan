// Copyright (C) 2021 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package config

import (
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

var (
	ConfName = "config.yml"
	Conf     = &AppConfigs{}
	MainEng  *xorm.Engine
	UserEng  *xorm.Engine
)

func init() {
	Conf = Load(ConfName)

	eng, err := xorm.NewEngine("sqlite3", "assets/main.db")
	if err != nil {
		panic(err)
	}
	err = eng.Ping()
	if err != nil {
		panic(err)
	}
	MainEng = eng
	MainEng.SetMaxOpenConns(50)
	MainEng.SetMaxIdleConns(10)

	eng, err = xorm.NewEngine("sqlite3", "assets/account.db")
	if err != nil {
		panic(err)
	}
	err = eng.Ping()
	if err != nil {
		panic(err)
	}
	UserEng = eng
	UserEng.SetMaxOpenConns(50)
	MainEng.SetMaxIdleConns(10)
}

// Copyright (C) 2021 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"honoka-chan/utils"
	"os"

	_ "modernc.org/sqlite"
	"xorm.io/xorm"
)

var (
	ConfName = "config.yml"
	Conf     = &AppConfigs{}

	ExampleDb = "assets/data.example.db"
	MainDb    = "assets/main.db"
	UserDb    = "assets/data.db"
	MainEng   *xorm.Engine
	UserEng   *xorm.Engine

	PackageVersion = "97.4.6"

	// LLAS
	StartUpKey    = "e0xrykyuBrLlwZhd"
	MasterVersion = "f7f2ac627227500b"
)

func init() {
	Conf = Load(ConfName)

	_, err := os.Stat(UserDb)
	if err != nil {
		utils.WriteAllText(UserDb, utils.ReadAllText(ExampleDb))
	}

	eng, err := xorm.NewEngine("sqlite", MainDb)
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

	eng, err = xorm.NewEngine("sqlite", UserDb)
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

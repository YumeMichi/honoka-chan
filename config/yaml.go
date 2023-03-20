// Copyright (C) 2021-2022 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"honoka-chan/utils"
	"honoka-chan/xclog"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

type AppConfigs struct {
	AppName string        `yaml:"app_name"`
	Server  ServerConfigs `yaml:"server"`
	Log     LogConfigs    `yaml:"log"`
	Redis   RedisConfigs  `yaml:"redis"`
	SifCap  SifCapConfigs `yaml:"sifcap"`
}

type ServerConfigs struct {
	PoweredBy     string `yaml:"powered_by"`
	VersionDate   string `yaml:"version_date"`
	VersionNumber string `yaml:"version_number"`
	VersionUp     string `yaml:"version_up"`
}

type LogConfigs struct {
	LogDir   string `yaml:"log_dir"`
	LogLevel int    `yaml:"log_level"`
	LogSave  bool   `yaml:"log_save"`
}

type RedisConfigs struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Pass string `yaml:"pass"`
	Db   int    `yaml:"db"`
}

type SifCapConfigs struct {
	Enabled bool `yaml:"enabled"`
}

func DefaultConfigs() *AppConfigs {
	return &AppConfigs{
		AppName: "LL! SIF Private Server",
		Server: ServerConfigs{
			PoweredBy:     "KLab Native APP Platform",
			VersionDate:   "20120129",
			VersionNumber: "97.4.6",
			VersionUp:     "0",
		},
		Log: LogConfigs{
			LogDir:   "logs",
			LogLevel: 5,
			LogSave:  true,
		},
		Redis: RedisConfigs{
			Host: "127.0.0.1",
			Port: "6379",
			Pass: "",
			Db:   0,
		},
		SifCap: SifCapConfigs{
			Enabled: false,
		},
	}
}

func Load(p string) *AppConfigs {
	if !utils.PathExists(p) {
		_ = DefaultConfigs().Save(p)
	}
	c := AppConfigs{}
	err := yaml.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	if err != nil {
		xclog.Error("Failed to load " + ConfName + ": " + err.Error())
		_ = os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		_ = DefaultConfigs().Save(p)
	}
	c = AppConfigs{}
	_ = yaml.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	xclog.Info(ConfName + " loaded!")
	return &c
}

func (c *AppConfigs) Save(p string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		xclog.Error("Failed to save " + ConfName + ": " + err.Error())
		return err
	}
	utils.WriteAllText(p, string(data))
	return nil
}

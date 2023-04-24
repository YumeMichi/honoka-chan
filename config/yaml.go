// Copyright (C) 2021-2022 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package config

import (
	"fmt"
	"honoka-chan/utils"
	"os"
	"strconv"
	"time"

	"gopkg.in/yaml.v3"
)

type AppConfigs struct {
	AppName string         `yaml:"app_name"`
	LevelDb LevelDbConfigs `yaml:"leveldb"`
	Cdn     CdnConfigs     `yaml:"cdn"`
}

type LevelDbConfigs struct {
	DataPath string `yaml:"data_path"`
}

type CdnConfigs struct {
	CdnUrl string `yaml:"cdn_url"`
}

func DefaultConfigs() *AppConfigs {
	return &AppConfigs{
		AppName: "LL! SIF Private Server",
		LevelDb: LevelDbConfigs{
			DataPath: "./data/honoka-chan.db",
		},
		Cdn: CdnConfigs{
			CdnUrl: "",
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
		fmt.Println("Failed to load" + ConfName + ":" + err.Error())
		_ = os.Rename(p, p+".backup"+strconv.FormatInt(time.Now().Unix(), 10))
		_ = DefaultConfigs().Save(p)
	}
	c = AppConfigs{}
	_ = yaml.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	fmt.Println(ConfName + "loaded!")
	return &c
}

func (c *AppConfigs) Save(p string) error {
	data, err := yaml.Marshal(c)
	if err != nil {
		fmt.Println("Failed to save" + ConfName + ":" + err.Error())
		return err
	}
	utils.WriteAllText(p, string(data))
	return nil
}

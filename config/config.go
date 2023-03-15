// Copyright (C) 2021 YumeMichi
//
// SPDX-License-Identifier: Apache-2.0
package config

var (
	ConfName = "config.yml"
	Conf     = &AppConfigs{}
)

func init() {
	Conf = Load(ConfName)
}

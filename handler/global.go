package handler

import (
	"honoka-chan/config"

	"xorm.io/xorm"
)

var (
	// nonce          = 0
	PackageVersion = "97.4.6"
	CdnUrl         string
	ErrorMsg       = `{"code":20001,"message":""}`
	MainEng        *xorm.Engine
	UserEng        *xorm.Engine
)

func init() {
	if config.Conf.Cdn.CdnUrl != "" {
		CdnUrl = config.Conf.Cdn.CdnUrl
	}

	MainEng = config.MainEng
	UserEng = config.UserEng
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

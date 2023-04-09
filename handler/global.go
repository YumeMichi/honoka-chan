package handler

import "honoka-chan/config"

var (
	nonce          = 0
	PackageVersion = "97.4.6"
	CdnUrl         string
	ErrorMsg       = `{"code":20001,"message":""}`
)

func init() {
	if config.Conf.Cdn.CdnUrl != "" {
		CdnUrl = config.Conf.Cdn.CdnUrl
	}
}

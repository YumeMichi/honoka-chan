package handler

import "honoka-chan/config"

var (
	nonce          = 0
	PackageVersion = "97.4.6"
	CdnUrl         string
)

func init() {
	if config.Conf.Cdn.CdnUrl != "" {
		CdnUrl = config.Conf.Cdn.CdnUrl
	}
}

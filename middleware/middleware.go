package middleware

import (
	"honoka-chan/config"

	"github.com/gin-gonic/gin"
)

func KlabHeader(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Header("X-Powered-By", config.Conf.Server.PoweredBy)
	ctx.Header("server_version", config.Conf.Server.VersionDate)
	ctx.Header("version_up", config.Conf.Server.VersionUp)
	ctx.Header("status_code", "200")
}

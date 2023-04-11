package middleware

import (
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/handler"
	"honoka-chan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonMid(ctx *gin.Context) {
	authorizeStr := ctx.Request.Header["Authorize"]
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) > 0 {
		authStr, _ := utils.GetAuthorizeToken(authorizeStr)
		res, _ := database.LevelDb.Get([]byte(userId[0]))
		if authStr != string(res) {
			ctx.String(http.StatusForbidden, handler.ErrorMsg)
			return
		}
	}
	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Header("X-Powered-By", config.Conf.Server.PoweredBy)
	ctx.Header("server_version", config.Conf.Server.VersionDate)
	ctx.Header("version_up", config.Conf.Server.VersionUp)
	ctx.Header("status_code", "200")
}

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WebAuth(ctx *gin.Context) {
	session := sessions.Default(ctx)
	requestUrl := strings.Split(ctx.Request.URL.String(), "?")[0] // 过滤 GET 参数
	userId, ok := session.Get("userid").(int)
	if ok {
		if requestUrl == "/admin/login" {
			ctx.Redirect(http.StatusFound, "/admin/index")
			ctx.Abort()
		}
	} else {
		if requestUrl != "/admin/login" {
			ctx.Redirect(http.StatusFound, "/admin/login")
			ctx.Abort()
		}
	}
	ctx.Set("userid", userId)

	ctx.Next()
}

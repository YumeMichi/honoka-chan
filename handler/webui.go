package handler

import (
	"honoka-chan/model"
	"net/http"

	"github.com/forgoer/openssl"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func WebLogin(ctx *gin.Context) {
	area := ctx.PostForm("area")
	user := ctx.PostForm("user")
	pass := ctx.PostForm("pass")
	if area == "" || user == "" || pass == "" {
		ctx.JSON(http.StatusOK, model.Msg{
			Code:     1,
			Message:  "参数不完整！",
			Redirect: "",
		})
		return
	}

	userName := " " + area + "-" + user
	exists, err := UserEng.Table("users").Where("phone = ? AND password = ?", userName, openssl.Md5ToString(pass)).Exist()
	CheckErr(err)
	if !exists {
		ctx.JSON(http.StatusOK, model.Msg{
			Code:     1,
			Message:  "账号不存在或者密码有误！",
			Redirect: "",
		})
		return
	}

	session := sessions.Default(ctx)
	session.Options(sessions.Options{
		MaxAge: 3600 * 24,
	})
	session.Set("username", userName)
	session.Save()

	ctx.JSON(http.StatusOK, model.Msg{
		Code:     0,
		Message:  "登录成功！",
		Redirect: "/admin/index",
	})
}

func WebLogout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{
		Path:   "/admin",
		MaxAge: -1,
	})
	session.Save()

	ctx.Redirect(http.StatusFound, "/admin/login")
}

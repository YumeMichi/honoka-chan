package middleware

import (
	"fmt"
	"honoka-chan/database"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ErrorMsg = `{"code":20001,"message":""}`
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Common(ctx *gin.Context) {
	ctx.Set("req_time", time.Now().Unix())

	authorize := ctx.Request.Header.Get("Authorize")
	if authorize == "" {
		ctx.String(http.StatusForbidden, ErrorMsg)
		ctx.Abort()
	}
	ctx.Set("authorize", authorize)

	params, err := url.ParseQuery(authorize)
	CheckErr(err)

	nonce, err := strconv.Atoi(params.Get("nonce"))
	CheckErr(err)
	nonce++
	ctx.Set("nonce", nonce)

	token := params.Get("token")
	ctx.Set("token", token)

	if ctx.Request.URL.String() == "/main.php/login/authkey" ||
		ctx.Request.URL.String() == "/main.php/login/login" {
		// 特殊请求
		fmt.Println("========")
	} else {
		userId := ctx.Request.Header.Get("User-ID")
		if userId == "" {
			ctx.String(http.StatusForbidden, ErrorMsg)
			ctx.Abort()
		}
		ctx.Set("userid", userId)

		rToken, err := database.LevelDb.Get([]byte(userId))
		CheckErr(err)
		if token != string(rToken) {
			ctx.String(http.StatusForbidden, ErrorMsg)
			ctx.Abort()
		}

		if !database.MatchTokenUid(token, userId) {
			ctx.String(http.StatusForbidden, ErrorMsg)
			ctx.Abort()
		}
	}

	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Header("X-Powered-By", "KLab Native APP Platform")
	ctx.Header("server_version", "20120129")
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("version_up", "0")
	ctx.Header("status_code", "200")

	ctx.Next()
}

func CommonAs(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}
	defer ctx.Request.Body.Close()
	ctx.Set("reqBody", string(body))

	ep := strings.ReplaceAll(ctx.Request.URL.String(), "/ep3071", "")
	ctx.Set("ep", ep)

	ctx.Next()
}

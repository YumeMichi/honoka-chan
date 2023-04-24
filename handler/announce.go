package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AnnounceIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "announce.tmpl", gin.H{
		"title":   "Love Live! 学园偶像祭 非官方服务器",
		"content": template.HTML("目前开发完毕的功能包括：登录、相册、Live、个人信息。<br>其他功能仍在开发中，有报错属于正常现象。"),
	})
}

func AnnounceCheckState(ctx *gin.Context) {
	announceResp := model.AnnounceResp{
		ResponseData: model.AnnounceRes{
			HasUnreadAnnounce: false,
			ServerTimestamp:   time.Now().Unix(),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(announceResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/utils"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AnnounceResp struct {
	ResponseData AnnounceData  `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}
type AnnounceData struct {
	HasUnreadAnnounce bool  `json:"has_unread_announce"`
	ServerTimestamp   int64 `json:"server_timestamp"`
}

func AnnounceIndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "announce.tmpl", gin.H{
		"title":   "Love Live! 学园偶像祭 非官方服务器",
		"content": template.HTML("目前开发完毕的功能包括：登录、相册、Live、个人信息。<br>其他功能仍在开发中，有报错属于正常现象。"),
	})
}

func AnnounceCheckStateHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	announceResp := AnnounceResp{
		ResponseData: AnnounceData{
			HasUnreadAnnounce: false,
			ServerTimestamp:   time.Now().Unix(),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(announceResp)
	if err != nil {
		panic(err)
	}
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthKey(ctx *gin.Context) {
	authResp := model.AuthKeyResp{
		ResponseData: model.AuthKeyRes{
			AuthorizeToken: ctx.GetString("authorize_token"),
			DummyToken:     ctx.GetString("dummy_token"),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(authResp)
	CheckErr(err)

	xMessageSign := base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem"))

	ctx.Header("X-Message-Sign", xMessageSign)

	ctx.JSON(http.StatusOK, authResp)
}

func Login(ctx *gin.Context) {
	loginKey := ctx.GetString("login_key")
	var userId int
	exists, err := UserEng.Table("user_key").Where("key = ?", loginKey).Cols("userid").Get(&userId)
	CheckErr(err)

	if !exists || userId == 0 {
		userId = 9999999
	}
	ctx.Set("userid", userId)

	err = database.LevelDb.Put([]byte(strconv.Itoa(userId)), []byte(ctx.GetString("authorize_token")))
	CheckErr(err)

	loginResp := model.LoginResp{
		ResponseData: model.LoginRes{
			AuthorizeToken:  ctx.GetString("authorize_token"),
			UserId:          userId,
			ServerTimestamp: time.Now().Unix(),
			AdultFlag:       2,
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(loginResp)
	CheckErr(err)

	ctx.Header("user_id", "")
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("authorize_token"), ctx.GetInt("nonce"), ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.JSON(http.StatusOK, loginResp)
}

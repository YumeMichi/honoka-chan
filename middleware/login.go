package middleware

import (
	"encoding/base64"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/utils"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func Login(ctx *gin.Context) {
	authData, err := database.LevelDb.Get([]byte(ctx.GetString("token")))
	CheckErr(err)

	clientToken, err := base64.StdEncoding.DecodeString(gjson.Get(string(authData), "client_token").String())
	CheckErr(err)
	serverToken, err := base64.StdEncoding.DecodeString(gjson.Get(string(authData), "server_token").String())
	CheckErr(err)

	xmcKey := utils.SliceXor(clientToken, serverToken)
	aesKey := xmcKey[0:16]

	req := gjson.Parse(ctx.GetString("request_data"))
	tKey, err := base64.StdEncoding.DecodeString(req.Get("login_key").String())
	CheckErr(err)
	loginKey := utils.Sub16(encrypt.AES_CBC_Decrypt(tKey, aesKey))
	ctx.Set("login_key", string(loginKey))

	tPasswd, err := base64.StdEncoding.DecodeString(req.Get("login_passwd").String())
	CheckErr(err)
	loginPasswd := utils.Sub16(encrypt.AES_CBC_Decrypt(tPasswd, aesKey))
	ctx.Set("login_passwd", string(loginPasswd))

	nonce := ctx.GetInt("nonce")
	nonce++
	ctx.Set("nonce", nonce)

	authorizeToken := base64.StdEncoding.EncodeToString([]byte(utils.RandomStr(32)))
	ctx.Set("authorize_token", authorizeToken)

	ctx.Next()

}

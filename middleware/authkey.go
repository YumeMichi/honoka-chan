package middleware

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

func AuthKey(ctx *gin.Context) {
	req := gjson.Parse(ctx.PostForm("request_data"))
	tDummyToken, err := base64.StdEncoding.DecodeString(req.Get("dummy_token").String())
	CheckErr(err)
	dummyToken := encrypt.RSA_Decrypt(tDummyToken, "privatekey.pem")

	// aesKey := dummyToken[0:16]
	// tAuthData, err := base64.StdEncoding.DecodeString(req.Get("auth_data").String())
	// CheckErr(err)
	// authData := utils.Sub16(encrypt.AES_CBC_Decrypt(tAuthData, aesKey))
	// fmt.Println(string(authData))

	clientToken := base64.StdEncoding.EncodeToString(dummyToken)
	serverToken := base64.StdEncoding.EncodeToString([]byte(utils.RandomStr(32)))
	authorizeToken := base64.StdEncoding.EncodeToString([]byte(utils.RandomStr(32)))

	ctx.Set("dummy_token", serverToken)
	ctx.Set("authorize_token", authorizeToken)

	authJson, err := json.Marshal(map[string]interface{}{
		"client_token": clientToken,
		"server_token": serverToken,
	})
	CheckErr(err)
	err = database.LevelDb.Put([]byte(authorizeToken), authJson)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++
	ctx.Set("nonce", nonce)

	authorize := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&nonce=%d&requestTimeStamp=%d", time.Now().Unix(), nonce, ctx.GetInt64("req_time"))

	ctx.Header("user_id", "")
	ctx.Header("authorize", authorize)

	ctx.Next()
}

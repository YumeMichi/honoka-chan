package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthKeyHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authReq := model.AuthKeyReq{}
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &authReq)
	if err != nil {
		panic(err)
	}

	dummyToken64, err := base64.StdEncoding.DecodeString(authReq.DummyToken)
	if err != nil {
		panic(err)
	}
	dummyTokenDecrypted := encrypt.RSA_Decrypt(dummyToken64, "privatekey.pem")

	aesKey := dummyTokenDecrypted[0:16]
	data64, err := base64.StdEncoding.DecodeString(authReq.AuthData)
	if err != nil {
		panic(err)
	}
	dataDecrypted := utils.Sub16(encrypt.AES_CBC_Decrypt(data64, aesKey))
	fmt.Println(string(dataDecrypted))

	mRandStr := utils.RandomStr(32)
	serverToken := base64.RawStdEncoding.EncodeToString([]byte(mRandStr))

	authorizeToken := utils.RandomBase64Token(32)

	nonce++

	clientToken := base64.RawStdEncoding.EncodeToString(dummyTokenDecrypted)
	authData := map[string]interface{}{
		"client_token": clientToken,
		"server_token": serverToken,
	}
	_, err = database.RedisCli.HSet(database.RedisCtx, authorizeToken, authData).Result()
	if err != nil {
		panic(err)
	}

	authResp := model.AuthKeyResp{}
	authResp.ResponseData.DummyToken = serverToken
	authResp.ResponseData.AuthorizeToken = authorizeToken
	authResp.StatusCode = 200
	resp, err := json.Marshal(authResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

	respTime := time.Now().Unix()
	authorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&nonce=%d&requestTimeStamp=%d", respTime, nonce, reqTime)
	fmt.Println(authorizeStr)

	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("user_id", "")
	ctx.Header("authorize", authorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.JSON(http.StatusOK, authResp)
}

func LoginHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	authData, err := database.RedisCli.HGetAll(database.RedisCtx, authToken).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(authData)

	clientToken, serverToken := authData["client_token"], authData["server_token"]
	clientToken64, err := base64.RawStdEncoding.DecodeString(clientToken)
	if err != nil {
		panic(err)
	}
	serverToken64, err := base64.RawStdEncoding.DecodeString(serverToken)
	if err != nil {
		panic(err)
	}

	xmcKey := utils.SliceXor([]byte(clientToken64), []byte(serverToken64))
	aesKey := xmcKey[0:16]

	loginReq := model.LoginReq{}
	err = json.Unmarshal([]byte(ctx.PostForm("request_data")), &loginReq)
	if err != nil {
		panic(err)
	}
	key64, err := base64.StdEncoding.DecodeString(loginReq.LoginKey)
	if err != nil {
		panic(err)
	}
	pass64, err := base64.StdEncoding.DecodeString(loginReq.LoginPasswd)
	if err != nil {
		panic(err)
	}
	keyDescrypted := utils.Sub16(encrypt.AES_CBC_Decrypt(key64, aesKey))
	fmt.Println(string(keyDescrypted))
	passDescrypted := utils.Sub16(encrypt.AES_CBC_Decrypt(pass64, aesKey))
	fmt.Println(string(passDescrypted))

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	nonce++

	// userId, err := database.GetUid(string(keyDescrypted))
	// if err != nil {
	// 	ctx.String(http.StatusForbidden, "Fuck you!")
	// 	return
	// }
	userId := 9999999
	authorizeToken := utils.RandomBase64Token(32)

	_, err = database.RedisCli.HSet(database.RedisCtx, "token_uid", authorizeToken, userId).Result()
	if err != nil {
		panic(err)
	}

	loginResp := model.LoginResp{}
	loginResp.ResponseData.AuthorizeToken = authorizeToken
	loginResp.ResponseData.UserId = userId
	loginResp.ResponseData.ServerTimestamp = time.Now().Unix()
	loginResp.ResponseData.AdultFlag = 2
	loginResp.StatusCode = 200
	resp, err := json.Marshal(loginResp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%d&requestTimeStamp=%d", respTime, authorizeToken, nonce, userId, reqTime)
	fmt.Println(newAuthorizeStr)

	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("user_id", "")
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.JSON(http.StatusOK, loginResp)
}

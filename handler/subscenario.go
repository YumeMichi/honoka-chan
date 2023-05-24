package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SubScenarioStartup(ctx *gin.Context) {
	startReq := model.SubScenarioReq{}
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &startReq)
	CheckErr(err)

	startResp := model.SubScenarioResp{
		ResponseData: model.SubScenarioRes{
			SubscenarioID:      startReq.SubscenarioID,
			ScenarioAdjustment: 50,
			ServerTimestamp:    time.Now().Unix(),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(startResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func SubScenarioReward(ctx *gin.Context) {
	resp := utils.ReadAllText("assets/sif/subreward.json")

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1([]byte(resp), "privatekey.pem")))

	ctx.String(http.StatusOK, resp)
}

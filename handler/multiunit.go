package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MultiUnitStartUpResp struct {
	ResponseData MultiUnitStartUpData `json:"response_data"`
	ReleaseInfo  []interface{}        `json:"release_info"`
	StatusCode   int                  `json:"status_code"`
}

type MultiUnitStartUpData struct {
	MultiUnitScenarioID int   `json:"multi_unit_scenario_id"`
	ScenarioAdjustment  int   `json:"scenario_adjustment"`
	ServerTimestamp     int64 `json:"server_timestamp"`
}

type MultiUnitStartUpReq struct {
	Module              string `json:"module"`
	Action              string `json:"action"`
	TimeStamp           int    `json:"timeStamp"`
	Mgd                 int    `json:"mgd"`
	MultiUnitScenarioID int    `json:"multi_unit_scenario_id"`
	CommandNum          string `json:"commandNum"`
}

func MultiUnitStartUp(ctx *gin.Context) {
	startReq := MultiUnitStartUpReq{}
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &startReq)
	CheckErr(err)

	startResp := MultiUnitStartUpResp{
		ResponseData: MultiUnitStartUpData{
			MultiUnitScenarioID: startReq.MultiUnitScenarioID,
			ScenarioAdjustment:  50,
			ServerTimestamp:     time.Now().Unix(),
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

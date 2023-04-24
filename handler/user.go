package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/tools"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

type NotificationResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

func SetNotificationToken(ctx *gin.Context) {
	notifResp := NotificationResp{
		ResponseData: []interface{}{},
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(notifResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func ChangeNavi(ctx *gin.Context) {
	req := gjson.Parse(ctx.PostForm("request_data"))
	pref := tools.UserPref{
		UnitOwningUserID: int(req.Get("unit_owning_user_id").Int()),
	}
	_, err := UserEng.Table("user_preference_m").Where("user_id = ?", ctx.GetString("userid")).Update(&pref)
	CheckErr(err)
	naviResp := model.UserNaviChangeResp{
		ResponseData: []interface{}{},
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(naviResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func ChangeName(ctx *gin.Context) {
	req := gjson.Parse(ctx.PostForm("request_data"))
	var oldName string
	exists, err := UserEng.Table("user_preference_m").Where("user_id = ?", ctx.GetString("userid")).Cols("user_name").Get(&oldName)
	CheckErr(err)
	if !exists {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	pref := tools.UserPref{
		UserName: req.Get("name").String(),
	}
	_, err = UserEng.Table("user_preference_m").Where("user_id = ?", ctx.GetString("userid")).Update(&pref)
	CheckErr(err)
	nameResp := model.UserNameChangeResp{
		ResponseData: model.UserNameChangeRes{
			BeforeName:      oldName,
			AfterName:       req.Get("name").String(),
			ServerTimestamp: time.Now().Unix(),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(nameResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type MuseumContent struct {
	MuseumContentsId int `xorm:"museum_contents_id"`
	SmileBuff        int `xorm:"smile_buff"`
	PureBuff         int `xorm:"pure_buff"`
	CoolBuff         int `xorm:"cool_buff"`
}

func MuseumInfo(ctx *gin.Context) {
	var contents []MuseumContent
	err := MainEng.Table("museum_contents_m").Cols("museum_contents_id,smile_buff,pure_buff,cool_buff").Find(&contents)
	CheckErr(err)
	var smileBuff, pureBuff, coolBuff int
	var contentsList []int
	for _, content := range contents {
		smileBuff += content.SmileBuff
		pureBuff += content.PureBuff
		coolBuff += content.CoolBuff
		contentsList = append(contentsList, content.MuseumContentsId)
	}
	museumResp := model.MuseumResp{
		ResponseData: model.MuseumRes{
			MuseumInfo: model.Museum{
				Parameter: model.MuseumParameter{
					Smile: smileBuff,
					Pure:  pureBuff,
					Cool:  coolBuff,
				},
				ContentsIDList: contentsList,
			},
			ServerTimestamp: time.Now().Unix(),
		},
		ReleaseInfo: []any{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(museumResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

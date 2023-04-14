package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UserInfoHandler(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.GetString("userid"))
	CheckErr(err)

	userResp := model.UserInfoResp{
		ResponseData: model.UserInfoRes{
			User: model.UserInfo{
				UserID:                         userId,
				Name:                           "\u68a6\u8def @\u65c5\u7acb\u3061\u306e\u65e5\u306b",
				Level:                          1028,
				Exp:                            28824396,
				PreviousExp:                    27734700,
				NextExp:                        28941885,
				GameCoin:                       112124104,
				SnsCoin:                        0,
				FreeSnsCoin:                    0,
				PaidSnsCoin:                    0,
				SocialPoint:                    1438395,
				UnitMax:                        5000,
				WaitingUnitMax:                 1000,
				EnergyMax:                      417,
				EnergyFullTime:                 "2023-03-20 03:58:55",
				LicenseLiveEnergyRecoverlyTime: 60,
				EnergyFullNeedTime:             0,
				OverMaxEnergy:                  0,
				TrainingEnergy:                 100,
				TrainingEnergyMax:              100,
				FriendMax:                      99,
				InviteCode:                     "377385143",
				InsertDate:                     "2015-08-10 18:58:30",
				UpdateDate:                     "2018-08-09 18:13:12",
				TutorialState:                  -1,
				DiamondCoin:                    0,
				CrystalCoin:                    0,
				LpRecoveryItem:                 []model.LpRecoveryItem{},
			},
			Birth: model.Birth{
				BirthMonth: 10,
				BirthDay:   18,
			},
			ServerTimestamp: time.Now().Unix(),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(userResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

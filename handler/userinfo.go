package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/tools"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UserInfo(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.GetString("userid"))
	CheckErr(err)

	pref := tools.UserPref{}
	exists, err := UserEng.Table("user_preference_m").Where("user_id = ?", userId).Get(&pref)
	CheckErr(err)
	if !exists {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	userResp := model.UserInfoResp{
		ResponseData: model.UserInfoRes{
			User: model.UserInfo{
				UserID:                         userId,
				Name:                           pref.UserName,
				Level:                          config.Conf.UserPrefs.Level,
				Exp:                            config.Conf.UserPrefs.ExpNumerator,
				PreviousExp:                    0,
				NextExp:                        config.Conf.UserPrefs.ExpDenominator,
				GameCoin:                       config.Conf.UserPrefs.GameCoin,
				SnsCoin:                        config.Conf.UserPrefs.SnsCoin,
				FreeSnsCoin:                    0,
				PaidSnsCoin:                    0,
				SocialPoint:                    1438395,
				UnitMax:                        5000,
				WaitingUnitMax:                 1000,
				EnergyMax:                      config.Conf.UserPrefs.EnergyMax,
				EnergyFullTime:                 "2023-03-20 03:58:55",
				LicenseLiveEnergyRecoverlyTime: 60,
				EnergyFullNeedTime:             0,
				OverMaxEnergy:                  config.Conf.UserPrefs.OverMaxEnergy,
				TrainingEnergy:                 100,
				TrainingEnergyMax:              100,
				FriendMax:                      99,
				InviteCode:                     config.Conf.UserPrefs.InviteCode,
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
		ReleaseInfo: []any{},
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

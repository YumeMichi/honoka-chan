package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserInfoResp struct {
	ResponseData UserInfoData  `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

type LpRecoveryItem struct {
	ItemID int `json:"item_id"`
	Amount int `json:"amount"`
}

type User struct {
	UserID                         int              `json:"user_id"`
	Name                           string           `json:"name"`
	Level                          int              `json:"level"`
	Exp                            int              `json:"exp"`
	PreviousExp                    int              `json:"previous_exp"`
	NextExp                        int              `json:"next_exp"`
	GameCoin                       int              `json:"game_coin"`
	SnsCoin                        int              `json:"sns_coin"`
	FreeSnsCoin                    int              `json:"free_sns_coin"`
	PaidSnsCoin                    int              `json:"paid_sns_coin"`
	SocialPoint                    int              `json:"social_point"`
	UnitMax                        int              `json:"unit_max"`
	WaitingUnitMax                 int              `json:"waiting_unit_max"`
	EnergyMax                      int              `json:"energy_max"`
	EnergyFullTime                 string           `json:"energy_full_time"`
	LicenseLiveEnergyRecoverlyTime int              `json:"license_live_energy_recoverly_time"`
	EnergyFullNeedTime             int              `json:"energy_full_need_time"`
	OverMaxEnergy                  int              `json:"over_max_energy"`
	TrainingEnergy                 int              `json:"training_energy"`
	TrainingEnergyMax              int              `json:"training_energy_max"`
	FriendMax                      int              `json:"friend_max"`
	InviteCode                     string           `json:"invite_code"`
	InsertDate                     string           `json:"insert_date"`
	UpdateDate                     string           `json:"update_date"`
	TutorialState                  int              `json:"tutorial_state"`
	DiamondCoin                    int              `json:"diamond_coin"`
	CrystalCoin                    int              `json:"crystal_coin"`
	LpRecoveryItem                 []LpRecoveryItem `json:"lp_recovery_item"`
}

type Birth struct {
	BirthMonth int `json:"birth_month"`
	BirthDay   int `json:"birth_day"`
}

type UserInfoData struct {
	User            User  `json:"user"`
	Birth           Birth `json:"birth"`
	ServerTimestamp int64 `json:"server_timestamp"`
}

func UserInfoHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	userResp := UserInfoResp{
		ResponseData: UserInfoData{
			User: User{
				UserID:                         9999999,
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
				LpRecoveryItem:                 []LpRecoveryItem{},
			},
			Birth: Birth{
				BirthMonth: 10,
				BirthDay:   18,
			},
			ServerTimestamp: time.Now().Unix(),
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(userResp)
	if err != nil {
		panic(err)
	}
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

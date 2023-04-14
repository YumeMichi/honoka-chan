package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/tools"
	"honoka-chan/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SetDisplayRankResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

func SetDisplayRankHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	dispResp := SetDisplayRankResp{
		ResponseData: []interface{}{},
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(dispResp)
	CheckErr(err)
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

func SetDeckHandler(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.GetString("userid"))
	CheckErr(err)

	deckReq := model.UnitDeckReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &deckReq); err != nil {
		panic(err)
	}

	// 开始事务
	// UserEng.ShowSQL(true)
	session := UserEng.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		session.Rollback()
		panic(err)
	}

	// 原有队伍信息
	var userDeckId []int
	err = session.Table("user_deck_m").Cols("id").Where("user_id = ?", userId).Find(&userDeckId)
	if err != nil {
		session.Rollback()
		panic(err)
	}

	// 删除全部原有队伍成员
	_, err = session.Table("deck_unit_m").In("user_deck_id", userDeckId).Delete()
	if err != nil {
		session.Rollback()
		panic(err)
	}

	// 删除全部原有队伍
	_, err = session.Table("user_deck_m").In("id", userDeckId).Delete()
	if err != nil {
		session.Rollback()
		panic(err)
	}

	// 遍历新队伍
	for _, deck := range deckReq.UnitDeckList {
		// 新队伍信息
		userDeck := tools.UserDeckData{
			DeckID:     deck.UnitDeckID,
			MainFlag:   deck.MainFlag,
			DeckName:   deck.DeckName,
			UserID:     userId,
			InsertDate: time.Now().Unix(),
		}
		_, err = session.Table("user_deck_m").Insert(&userDeck)
		if err != nil {
			session.Rollback()
			panic(err)
		}
		userDeckId := userDeck.ID
		// fmt.Println("新队伍 ID:", userDeckId)

		// 队伍成员信息
		for _, unit := range deck.UnitDeckDetail {
			// 成员信息
			newUnitData := tools.UnitData{}
			exists, err := session.Table("user_unit_m").Where("unit_owning_user_id = ?", unit.UnitOwningUserID).Exist()
			if err != nil {
				session.Rollback()
				panic(err)
			}
			if exists {
				// fmt.Println("新成员为用户增加成员")
				_, err = session.Table("user_unit_m").Where("unit_owning_user_id = ?", unit.UnitOwningUserID).Get(&newUnitData)
				if err != nil {
					session.Rollback()
					panic(err)
				}
			} else {
				exists, err := MainEng.Table("common_unit_m").Where("unit_owning_user_id = ?", unit.UnitOwningUserID).Exist()
				if err != nil {
					session.Rollback()
					panic(err)
				}
				if exists {
					// fmt.Println("新成员为公共成员")
					_, err = MainEng.Table("common_unit_m").Where("unit_owning_user_id = ?", unit.UnitOwningUserID).Get(&newUnitData)
					if err != nil {
						session.Rollback()
						panic(err)
					}
				} else {
					// fmt.Println("新成员不存在")
					session.Rollback()
					panic("unexpected operation")
				}
			}
			// fmt.Println("新的成员信息:", newUnitData)

			// 插入新成员信息
			newUnitDeckData := tools.UnitDeckData{}
			b, err := json.Marshal(newUnitData)
			if err != nil {
				session.Rollback()
				panic(err)
			}
			if err = json.Unmarshal(b, &newUnitDeckData); err != nil {
				session.Rollback()
				panic(err)
			}
			newUnitDeckData.BeforeLove = newUnitDeckData.MaxLove
			newUnitDeckData.Position = unit.Position
			newUnitDeckData.UserDeckID = userDeckId
			newUnitDeckData.InsertData = time.Now().Unix()

			_, err = session.Table("deck_unit_m").Insert(&newUnitDeckData)
			if err != nil {
				session.Rollback()
				panic(err)
			}
		}
	}

	// 结束事务
	if err = session.Commit(); err != nil {
		session.Rollback()
		panic(err)
	}

	dispResp := SetDisplayRankResp{
		ResponseData: []interface{}{},
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(dispResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func SetDeckNameHandler(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.GetString("userid"))
	CheckErr(err)

	deckReq := model.DeckNameReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &deckReq); err != nil {
		panic(err)
	}

	exists, err := UserEng.Table("user_deck_m").Where("user_id = ? AND deck_id = ?", userId, deckReq.UnitDeckID).Exist()
	CheckErr(err)
	if !exists {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	userDeck := tools.UserDeckData{
		DeckName: deckReq.DeckName,
	}
	_, err = UserEng.Table("user_deck_m").Update(&userDeck, &tools.UserDeckData{
		UserID: userId,
		DeckID: deckReq.UnitDeckID,
	})
	CheckErr(err)

	dispResp := SetDisplayRankResp{
		ResponseData: []interface{}{},
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(dispResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

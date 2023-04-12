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
	uId, _ := strconv.Atoi(userId[0])

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

	deckReq := model.UnitDeckReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &deckReq); err != nil {
		panic(err)
	}

	UserEng.ShowSQL(true)
	session := UserEng.NewSession()
	defer session.Close()

	for _, deck := range deckReq.UnitDeckList {
		exists, err := UserEng.Table("user_deck_m").Where("user_id = ? AND deck_id = ?", uId, deck.UnitDeckID).Exist()
		CheckErr(err)

		if exists { // 原队伍
			// 原队伍信息
			var userDeckId int
			_, err = UserEng.Table("user_deck_m").Cols("id").Where("user_id = ? AND deck_id = ?", uId, deck.UnitDeckID).Get(&userDeckId)
			CheckErr(err)

			// fmt.Println(userDeckId)

			// 事务
			if err = session.Begin(); err != nil {
				session.Rollback()
				panic(err)
			}
			// 更新队伍信息
			userDeck := tools.UserDeckData{
				MainFlag: deck.MainFlag,
				DeckName: deck.DeckName,
			}
			_, err := session.Table("user_deck_m").Update(&userDeck, &tools.UserDeckData{
				UserID: uId,
				DeckID: deck.UnitDeckID,
			})
			if err != nil {
				session.Rollback()
				panic(err)
			}
			// fmt.Println("Aff:", r)

			// 更新队伍成员信息
			for _, unit := range deck.UnitDeckDetail {
				// 原队伍成员位置信息
				unitDeckData := tools.UnitDeckData{}
				_, err = session.Table("deck_unit_m").Where("user_deck_id = ? AND position = ?", userDeckId, unit.Position).Get(&unitDeckData)
				if err != nil {
					session.Rollback()
					panic(err)
				}

				// 该位置成员信息不变
				if unitDeckData.UnitOwningUserID == unit.UnitOwningUserID {
					continue
				}

				// 新成员信息
				// fmt.Printf("位置: %d 成员发生改变, %d -> %d\n", unit.Position, unitDeckData.UnitOwningUserID, unit.UnitOwningUserID)
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
						panic("wtf?")
					}
				}
				// fmt.Println("新的成员信息:", newUnitData)

				// 更新新成员信息
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

				_, err = session.Table("deck_unit_m").Update(&newUnitDeckData, &tools.UnitDeckData{
					ID: unitDeckData.ID,
				})
				if err != nil {
					session.Rollback()
					panic(err)
				}
			}
			if err = session.Commit(); err != nil {
				session.Rollback()
				panic(err)
			}
		} else { // 新队伍
			if err := session.Begin(); err != nil {
				session.Rollback()
				panic(err)
			}

			// 新队伍信息
			userDeck := tools.UserDeckData{
				DeckID:     deck.UnitDeckID,
				MainFlag:   deck.MainFlag,
				DeckName:   deck.DeckName,
				UserID:     uId,
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
						panic("wtf?")
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
			if err = session.Commit(); err != nil {
				session.Rollback()
				panic(err)
			}
		}
	}

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

func SetDeckNameHandler(ctx *gin.Context) {
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
	uId, _ := strconv.Atoi(userId[0])

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

	deckReq := model.DeckNameReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &deckReq); err != nil {
		panic(err)
	}

	exists, err := UserEng.Table("user_deck_m").Where("user_id = ? AND deck_id = ?", uId, deckReq.UnitDeckID).Exist()
	CheckErr(err)
	if !exists {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	userDeck := tools.UserDeckData{
		DeckName: deckReq.DeckName,
	}
	_, err = UserEng.Table("user_deck_m").Update(&userDeck, &tools.UserDeckData{
		UserID: uId,
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
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

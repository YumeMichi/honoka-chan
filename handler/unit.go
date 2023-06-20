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

func SetDisplayRank(ctx *gin.Context) {
	dispResp := model.SetDisplayRankResp{
		ResponseData: []any{},
		ReleaseInfo:  []any{},
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

func SetDeck(ctx *gin.Context) {
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
		userDeck := model.UserDeckData{
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
			newUnitData := model.UnitData{}
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
			newUnitDeckData := model.UnitDeckData{}
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

	dispResp := model.SetDeckResp{
		ResponseData: []any{},
		ReleaseInfo:  []any{},
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

func SetDeckName(ctx *gin.Context) {
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
	userDeck := model.UserDeckData{
		DeckName: deckReq.DeckName,
	}
	_, err = UserEng.Table("user_deck_m").Update(&userDeck, &model.UserDeckData{
		UserID: userId,
		DeckID: deckReq.UnitDeckID,
	})
	CheckErr(err)

	dispResp := model.SetDeckResp{
		ResponseData: []any{},
		ReleaseInfo:  []any{},
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

func WearAccessory(ctx *gin.Context) {
	fmt.Println(ctx.PostForm("request_data"))
	req := model.WearAccessoryReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &req); err != nil {
		panic(err)
	}

	// UserEng.ShowSQL(true)
	// 开始事务
	session := UserEng.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		session.Rollback()
		panic(err)
	}

	// 取下饰品
	for _, v := range req.Remove {
		fmt.Println("Remove:", v.AccessoryOwningUserID, v.UnitOwningUserID)
		_, err := session.Table("accessory_wear_m").
			Where("accessory_owning_user_id = ? AND unit_owning_user_id = ? AND user_id = ?", v.AccessoryOwningUserID, v.UnitOwningUserID, ctx.GetString("userid")).
			Delete()
		if err != nil {
			session.Rollback()
			panic(err)
		}
	}

	// 佩戴饰品
	for _, v := range req.Wear {
		fmt.Println("Wear:", v.AccessoryOwningUserID, v.UnitOwningUserID)
		data := model.AccessoryWearData{
			AccessoryOwningUserID: v.AccessoryOwningUserID,
			UnitOwningUserID:      v.UnitOwningUserID,
			UserId:                ctx.GetString("userid"),
		}
		_, err := session.Table("accessory_wear_m").Insert(&data)
		CheckErr(err)
	}

	// 结束事务
	if err := session.Commit(); err != nil {
		session.Rollback()
		panic(err)
	}

	wearResp := model.AwardSetResp{
		ResponseData: []any{},
		ReleaseInfo:  []any{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(wearResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func RemoveSkillEquip(ctx *gin.Context) {
	fmt.Println(ctx.PostForm("request_data"))
	req := model.SkillEquipReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &req); err != nil {
		panic(err)
	}

	// UserEng.ShowSQL(true)
	// 开始事务
	session := UserEng.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		session.Rollback()
		panic(err)
	}

	// 取下宝石
	for _, v := range req.Remove {
		fmt.Println("Remove:", v.UnitOwningUserID, v.UnitRemovableSkillID)
		_, err := session.Table("skill_equip_m").
			Where("unit_removable_skill_id = ? AND unit_owning_user_id = ? AND user_id = ?", v.UnitRemovableSkillID, v.UnitOwningUserID, ctx.GetString("userid")).
			Delete()
		if err != nil {
			session.Rollback()
			panic(err)
		}
	}

	// 佩戴宝石
	for _, v := range req.Equip {
		fmt.Println("Equip:", v.UnitOwningUserID, v.UnitRemovableSkillID)
		data := model.SkillEquipData{
			UnitRemovableSkillId: v.UnitRemovableSkillID,
			UnitOwningUserID:     v.UnitOwningUserID,
			UserId:               ctx.GetString("userid"),
		}
		_, err := session.Table("skill_equip_m").Insert(&data)
		CheckErr(err)
	}

	// 结束事务
	if err := session.Commit(); err != nil {
		session.Rollback()
		panic(err)
	}

	wearResp := model.AwardSetResp{
		ResponseData: []any{},
		ReleaseInfo:  []any{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(wearResp)
	CheckErr(err)
	fmt.Println(string(resp))

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

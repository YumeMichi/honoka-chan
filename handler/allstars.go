package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func AsLogin(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var mask string
	req := gjson.Parse(reqBody)
	req.ForEach(func(key, value gjson.Result) bool {
		if value.Get("mask").String() != "" {
			mask = value.Get("mask").String()
			return false
		}
		return true
	})
	// fmt.Println("Request data:", req.String())
	// fmt.Println("Mask:", mask)

	mask64, err := base64.StdEncoding.DecodeString(mask)
	if err != nil {
		panic(err)
	}
	randomBytes := encrypt.RSA_DecryptOAEP(mask64, "privatekey.pem")
	// fmt.Println("Random Bytes:", randomBytes)

	newKey := utils.SliceXor(randomBytes, []byte(sessionKey))
	newKey64 := base64.StdEncoding.EncodeToString(newKey)

	loginBody := GetUserData("login.json")
	loginBody, _ = sjson.Set(loginBody, "session_key", newKey64)
	loginBody, _ = sjson.Set(loginBody, "user_model.user_status", GetUserStatus())

	/* ======== UserData ======== */
	liveDeckData := gjson.Parse(GetUserData("liveDeck.json"))
	loginBody, _ = sjson.Set(loginBody, "user_model.user_live_deck_by_id", liveDeckData.Get("user_live_deck_by_id").Value())

	var liveParty []any
	decoder := json.NewDecoder(strings.NewReader(liveDeckData.Get("user_live_party_by_id").String()))
	decoder.UseNumber()
	err = decoder.Decode(&liveParty)
	CheckErr(err)
	// fmt.Println(liveParty)
	loginBody, _ = sjson.Set(loginBody, "user_model.user_live_party_by_id", liveParty)

	memberData := gjson.Parse(GetUserData("memberSettings.json"))
	loginBody, _ = sjson.Set(loginBody, "user_model.user_member_by_member_id", memberData.Get("user_member_by_member_id").Value())

	lessonData := gjson.Parse(GetUserData("lessonDeck.json"))
	loginBody, _ = sjson.Set(loginBody, "user_model.user_lesson_deck_by_id", lessonData.Get("user_lesson_deck_by_id").Value())

	cardData := gjson.Parse(GetUserData("userCard.json"))
	loginBody, _ = sjson.Set(loginBody, "user_model.user_card_by_card_id", cardData.Get("user_card_by_card_id").Value())
	/* ======== UserData ======== */

	resp := SignResp(ctx.GetString("ep"), loginBody, config.StartUpKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchBootstrap(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("fetchBootstrap.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchBillingHistory(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("fetchBillingHistory.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchNotice(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("fetchNotice.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsGetPackUrl(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	var packNames []string
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			if err := json.Unmarshal([]byte(value.Get("pack_names").String()), &packNames); err != nil {
				panic(err)
			}
			return false
		}
		return true
	})

	var packUrls []string
	for _, pack := range packNames {
		packUrls = append(packUrls, AsCdnServer+"/"+config.MasterVersion+"/"+pack)
	}

	packBody, _ := sjson.Set("{}", "url_list", packUrls)
	resp := SignResp(ctx.GetString("ep"), packBody, sessionKey)
	// fmt.Println("Response:", resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsUpdateCardNewFlag(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("updateCardNewFlag.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsGetClearedPlatformAchievement(ctx *gin.Context) {
	signBody := GetUserData("getClearedPlatformAchievement.json")
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchLiveMusicSelect(ctx *gin.Context) {
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	liveDailyList := []model.LiveDaily{}
	err := MainEng.Table("m_live_daily").Where("weekday = ?", weekday).Cols("id,live_id").Find(&liveDailyList)
	CheckErr(err)
	for k := range liveDailyList {
		liveDailyList[k].EndAt = int(tomorrow)
		liveDailyList[k].RemainingPlayCount = 5
		liveDailyList[k].RemainingRecoveryCount = 9
	}

	signBody := GetUserData("fetchLiveMusicSelect.json")
	signBody, _ = sjson.Set(signBody, "weekday_state.weekday", weekday)
	signBody, _ = sjson.Set(signBody, "weekday_state.next_weekday_at", tomorrow)
	signBody, _ = sjson.Set(signBody, "live_daily_list", liveDailyList)
	signBody, _ = sjson.Set(signBody, "user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsLiveMvStart(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("liveMvStart.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsLiveMvSaveDeck(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	reqData := gjson.Parse(reqBody).Array()[0]
	// fmt.Println(reqData)

	saveReq := model.LiveSaveDeckReq{}
	err := json.Unmarshal([]byte(reqData.String()), &saveReq)
	if err != nil {
		panic(err)
	}
	// fmt.Println(saveReq)

	userLiveMvDeckInfo := model.UserLiveMvDeckInfo{
		LiveMasterID: saveReq.LiveMasterID,
	}

	memberIds := map[int]int{}
	for k, v := range saveReq.MemberMasterIDByPos {
		if k%2 == 0 {
			memberId := saveReq.MemberMasterIDByPos[k+1]
			memberIds[v] = memberId

			switch v {
			case 1:
				userLiveMvDeckInfo.MemberMasterID1 = memberId
			case 2:
				userLiveMvDeckInfo.MemberMasterID2 = memberId
			case 3:
				userLiveMvDeckInfo.MemberMasterID3 = memberId
			case 4:
				userLiveMvDeckInfo.MemberMasterID4 = memberId
			case 5:
				userLiveMvDeckInfo.MemberMasterID5 = memberId
			case 6:
				userLiveMvDeckInfo.MemberMasterID6 = memberId
			case 7:
				userLiveMvDeckInfo.MemberMasterID7 = memberId
			case 8:
				userLiveMvDeckInfo.MemberMasterID8 = memberId
			case 9:
				userLiveMvDeckInfo.MemberMasterID9 = memberId
			case 10:
				userLiveMvDeckInfo.MemberMasterID10 = memberId
			case 11:
				userLiveMvDeckInfo.MemberMasterID11 = memberId
			case 12:
				userLiveMvDeckInfo.MemberMasterID12 = memberId
			}
		}
	}
	// fmt.Println(memberIds)

	suitIds := map[int]int{}
	for k, v := range saveReq.SuitMasterIDByPos {
		if k%2 == 0 {
			suitId := saveReq.SuitMasterIDByPos[k+1]
			suitIds[v] = suitId

			switch v {
			case 1:
				userLiveMvDeckInfo.SuitMasterID1 = suitId
			case 2:
				userLiveMvDeckInfo.SuitMasterID2 = suitId
			case 3:
				userLiveMvDeckInfo.SuitMasterID3 = suitId
			case 4:
				userLiveMvDeckInfo.SuitMasterID4 = suitId
			case 5:
				userLiveMvDeckInfo.SuitMasterID5 = suitId
			case 6:
				userLiveMvDeckInfo.SuitMasterID6 = suitId
			case 7:
				userLiveMvDeckInfo.SuitMasterID7 = suitId
			case 8:
				userLiveMvDeckInfo.SuitMasterID8 = suitId
			case 9:
				userLiveMvDeckInfo.SuitMasterID9 = suitId
			case 10:
				userLiveMvDeckInfo.SuitMasterID10 = suitId
			case 11:
				userLiveMvDeckInfo.SuitMasterID11 = suitId
			case 12:
				userLiveMvDeckInfo.SuitMasterID12 = suitId
			}
		}
	}
	// fmt.Println(suitIds)

	var userLiveMvDeckCustomByID []any
	userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, saveReq.LiveMasterID)
	userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, userLiveMvDeckInfo)
	// fmt.Println(userLiveMvDeckCustomByID)

	signBody := GetUserData("liveMvSaveDeck.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_live_mv_deck_custom_by_id", userLiveMvDeckCustomByID)

	resp := SignResp(ctx.GetString("ep"), string(signBody), sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchCommunicationMemberDetail(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	var memberId int64
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("member_id").String() != "" {
			memberId = value.Get("member_id").Int()
			return false
		}
		return true
	})

	lovePanelCellIds := []int{}
	err := MainEng.Table("m_member_love_panel_cell").
		Join("LEFT", "m_member_love_panel", "m_member_love_panel_cell.member_love_panel_master_id = m_member_love_panel.id").
		Cols("m_member_love_panel_cell.id").Where("m_member_love_panel.member_master_id = ?", memberId).
		OrderBy("m_member_love_panel_cell.id ASC").Find(&lovePanelCellIds)
	CheckErr(err)

	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	tomorrow := time.Date(year, month, day+1, 0, 0, 0, 0, now.Location()).Unix()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	signBody := GetUserData("fetchCommunicationMemberDetail.json")
	signBody, _ = sjson.Set(signBody, "member_love_panels.0.member_id", memberId)
	signBody, _ = sjson.Set(signBody, "member_love_panels.0.member_love_panel_cell_ids", lovePanelCellIds)
	signBody, _ = sjson.Set(signBody, "weekday_state.weekday", weekday)
	signBody, _ = sjson.Set(signBody, "weekday_state.next_weekday_at", tomorrow)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsTapLovePoint(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("tapLovePoint.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsUpdateUserCommunicationMemberDetailBadge(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	var memberMasterId int64
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("member_master_id").String() != "" {
			memberMasterId = value.Get("member_master_id").Int()
			return false
		}
		return true
	})

	userDetail := []any{}
	userDetail = append(userDetail, memberMasterId)
	userDetail = append(userDetail, model.UserCommunicationMemberDetailBadgeByID{
		MemberMasterID: int(memberMasterId),
	})

	signBody := GetUserData("updateUserCommunicationMemberDetailBadge.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_communication_member_detail_badge_by_id", userDetail)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsUpdateUserLiveDifficultyNewFlag(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("updateUserLiveDifficultyNewFlag.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishUserStorySide(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishUserStorySide.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishUserStoryMember(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishUserStoryMember.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSetTheme(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	var memberMasterId, suitMasterId, backgroundMasterId int64
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("member_master_id").String() != "" {
			memberMasterId = value.Get("member_master_id").Int()
			suitMasterId = value.Get("suit_master_id").Int()
			backgroundMasterId = value.Get("custom_background_master_id").Int()

			gjson.Parse(GetUserData("memberSettings.json")).Get("user_member_by_member_id").
				ForEach(func(kk, vv gjson.Result) bool {
					if vv.IsObject() {
						if vv.Get("member_master_id").Int() == memberMasterId {
							SetUserData("memberSettings.json", "user_member_by_member_id."+
								kk.String()+".custom_background_master_id", backgroundMasterId)
							SetUserData("memberSettings.json", "user_member_by_member_id."+
								kk.String()+".suit_master_id", suitMasterId)
							return false
						}
					}
					return true
				})
			return false
		}
		return true
	})

	userMemberRes := []any{}
	userMemberRes = append(userMemberRes, memberMasterId)
	userMemberRes = append(userMemberRes, model.UserMemberInfo{
		MemberMasterID:           int(memberMasterId),
		CustomBackgroundMasterID: int(backgroundMasterId),
		SuitMasterID:             int(suitMasterId),
		LovePoint:                10271130,
		LovePointLimit:           10271130,
		LoveLevel:                450,
		ViewStatus:               1,
		IsNew:                    false,
	})

	userSuitRes := []any{}
	userSuitRes = append(userSuitRes, suitMasterId)
	userSuitRes = append(userSuitRes, model.AsSuitInfo{
		SuitMasterID: int(suitMasterId),
		IsNew:        false,
	})

	signBody := GetUserData("setTheme.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_member_by_member_id", userMemberRes)
	signBody, _ = sjson.Set(signBody, "user_model.user_suit_by_suit_id", userSuitRes)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchProfile(ctx *gin.Context) {
	userInfo := gjson.Parse(GetUserData("userStatus.json"))
	signBody := GetUserData("fetchProfile.json")
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.name.dot_under_text",
		userInfo.Get("name.dot_under_text").String())
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.introduction_message.dot_under_text",
		userInfo.Get("message.dot_under_text").String())
	signBody, _ = sjson.Set(signBody, "profile_info.basic_info.emblem_id",
		userInfo.Get("emblem_id").Int())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSetProfile(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	if req.Get("name").String() != "" {
		SetUserData("userStatus.json", "name.dot_under_text",
			gjson.Parse(reqBody).Array()[0].Get("name").String())
	} else if req.Get("nickname").String() != "" {
		SetUserData("userStatus.json", "nickname.dot_under_text",
			gjson.Parse(reqBody).Array()[0].Get("nickname").String())
	} else if req.Get("message").String() != "" {
		SetUserData("userStatus.json", "message.dot_under_text",
			gjson.Parse(reqBody).Array()[0].Get("message").String())
	}

	signBody, _ := sjson.Set(GetUserData("setProfile.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchEmblem(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("fetchEmblem.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsActivateEmblem(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var emblemId int64
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("emblem_master_id").String() != "" {
			emblemId = value.Get("emblem_master_id").Int()

			SetUserData("userStatus.json", "emblem_id", emblemId)

			return false
		}
		return true
	})

	signBody := GetUserData("activateEmblem.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_status.emblem_id", emblemId)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSaveUserNaviVoice(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("saveUserNaviVoice.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSaveDeckAll(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	req := model.AsSaveDeckReq{}
	decoder := json.NewDecoder(strings.NewReader(reqBody.String()))
	decoder.UseNumber()
	err := decoder.Decode(&req)
	CheckErr(err)
	// fmt.Println("Raw:", req.SquadDict)

	liveDeckInfo := GetUserData("liveDeck.json")
	keyDeckName := fmt.Sprintf("user_live_deck_by_id.%d.name.dot_under_text", req.DeckID*2-1)
	// fmt.Println(keyDeckName)
	deckName := gjson.Parse(liveDeckInfo).Get(keyDeckName).String()
	// fmt.Println("deckName:", deckName)

	if req.CardWithSuit[1] == 0 {
		req.CardWithSuit[1] = req.CardWithSuit[0]
	}
	if req.CardWithSuit[3] == 0 {
		req.CardWithSuit[3] = req.CardWithSuit[2]
	}
	if req.CardWithSuit[5] == 0 {
		req.CardWithSuit[5] = req.CardWithSuit[4]
	}
	if req.CardWithSuit[7] == 0 {
		req.CardWithSuit[7] = req.CardWithSuit[6]
	}
	if req.CardWithSuit[9] == 0 {
		req.CardWithSuit[9] = req.CardWithSuit[8]
	}
	if req.CardWithSuit[11] == 0 {
		req.CardWithSuit[11] = req.CardWithSuit[10]
	}
	if req.CardWithSuit[13] == 0 {
		req.CardWithSuit[13] = req.CardWithSuit[12]
	}
	if req.CardWithSuit[15] == 0 {
		req.CardWithSuit[15] = req.CardWithSuit[14]
	}
	if req.CardWithSuit[17] == 0 {
		req.CardWithSuit[17] = req.CardWithSuit[16]
	}

	deckInfo := model.AsDeckInfo{
		UserLiveDeckID: req.DeckID,
		Name: model.AsDeckName{
			DotUnderText: deckName,
		},
		CardMasterID1: req.CardWithSuit[0],
		CardMasterID2: req.CardWithSuit[2],
		CardMasterID3: req.CardWithSuit[4],
		CardMasterID4: req.CardWithSuit[6],
		CardMasterID5: req.CardWithSuit[8],
		CardMasterID6: req.CardWithSuit[10],
		CardMasterID7: req.CardWithSuit[12],
		CardMasterID8: req.CardWithSuit[14],
		CardMasterID9: req.CardWithSuit[16],
		SuitMasterID1: req.CardWithSuit[1],
		SuitMasterID2: req.CardWithSuit[3],
		SuitMasterID3: req.CardWithSuit[5],
		SuitMasterID4: req.CardWithSuit[7],
		SuitMasterID5: req.CardWithSuit[9],
		SuitMasterID6: req.CardWithSuit[11],
		SuitMasterID7: req.CardWithSuit[13],
		SuitMasterID8: req.CardWithSuit[15],
		SuitMasterID9: req.CardWithSuit[17],
	}
	// fmt.Println(deckInfo)

	keyLiveDeck := fmt.Sprintf("user_live_deck_by_id.%d", req.DeckID*2-1)
	SetUserData("liveDeck.json", keyLiveDeck, deckInfo)

	deckInfoRes := []model.AsResp{}
	deckInfoRes = append(deckInfoRes, req.DeckID)
	deckInfoRes = append(deckInfoRes, deckInfo)

	partyInfoRes := []model.AsResp{}
	for k, v := range req.SquadDict {
		if k%2 == 0 {
			partyId, err := v.(json.Number).Int64()
			if err != nil {
				panic(err)
			}
			// fmt.Println("Party ID:", partyId)

			rDictInfo, err := json.Marshal(req.SquadDict[k+1])
			CheckErr(err)

			dictInfo := model.AsDeckSquadDict{}
			decoder := json.NewDecoder(bytes.NewReader(rDictInfo))
			decoder.UseNumber()
			err = decoder.Decode(&dictInfo)
			CheckErr(err)
			// fmt.Println("Party Info:", dictInfo)

			roleIds := []int{}
			err = MainEng.Table("m_card").
				Where("id IN (?,?,?)", dictInfo.CardMasterIds[0], dictInfo.CardMasterIds[1], dictInfo.CardMasterIds[2]).
				Cols("role").Find(&roleIds)
			CheckErr(err)
			// fmt.Println("roleIds:", roleIds)

			partyIcon, partyName := GetPartyInfoByRoleIds(roleIds)
			realPartyName := GetRealPartyName(partyName)
			partyInfo := model.AsPartyInfo{
				PartyID:        int(partyId),
				UserLiveDeckID: req.DeckID,
				Name: model.AsPartyName{
					DotUnderText: realPartyName,
				},
				IconMasterID:     partyIcon,
				CardMasterID1:    dictInfo.CardMasterIds[0],
				CardMasterID2:    dictInfo.CardMasterIds[1],
				CardMasterID3:    dictInfo.CardMasterIds[2],
				UserAccessoryID1: dictInfo.UserAccessoryIds[0],
				UserAccessoryID2: dictInfo.UserAccessoryIds[1],
				UserAccessoryID3: dictInfo.UserAccessoryIds[2],
			}
			// fmt.Println(partyInfo)

			gjson.Parse(liveDeckInfo).Get("user_live_party_by_id").ForEach(func(key, value gjson.Result) bool {
				if value.IsObject() && value.Get("party_id").Int() == partyId {
					SetUserData("liveDeck.json", "user_live_party_by_id."+key.String(), partyInfo)
					return false
				}
				return true
			})

			partyInfoRes = append(partyInfoRes, partyId)
			partyInfoRes = append(partyInfoRes, partyInfo)
		}
	}

	respBody := GetUserData("saveDeckAll.json")
	respBody, _ = sjson.Set(respBody, "user_model.user_status", GetUserStatus())
	respBody, _ = sjson.Set(respBody, "user_model.user_live_deck_by_id", deckInfoRes)
	respBody, _ = sjson.Set(respBody, "user_model.user_live_party_by_id", partyInfoRes)
	resp := SignResp(ctx.GetString("ep"), respBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchLivePartners(ctx *gin.Context) {
	signBody := GetUserData("fetchLivePartners.json")
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchLiveDeckSelect(ctx *gin.Context) {
	signBody := GetUserData("fetchLiveDeckSelect.json")
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsLiveStart(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	liveStartReq := model.AsLiveStartReq{}
	if err := json.Unmarshal([]byte(reqBody.String()), &liveStartReq); err != nil {
		panic(err)
	}
	// fmt.Println(liveStartReq)

	var cardInfo string
	partnerResp := gjson.Parse(GetUserData("fetchLivePartners.json")).Get("partner_select_state.live_partners")
	partnerResp.ForEach(func(k, v gjson.Result) bool {
		userId := v.Get("user_id").Int()
		if userId == int64(liveStartReq.PartnerUserID) {
			v.Get("card_by_category").ForEach(func(kk, vv gjson.Result) bool {
				if vv.IsObject() {
					cardId := vv.Get("card_master_id").Int()
					if cardId == int64(liveStartReq.PartnerCardMasterID) {
						cardInfo = vv.String()
						// fmt.Println(cardInfo)
						return false
					}
				}
				return true
			})
			return false
		}
		return true
	})

	// 保存请求包因为 /live/finish 接口的响应包里有部分字段不在该接口的请求包里
	liveId := time.Now().UnixNano()
	liveIdStr := strconv.Itoa(int(liveId))
	err := database.LevelDb.Put([]byte("live_"+liveIdStr), []byte(reqBody.String()))
	CheckErr(err)

	liveDifficultyId := strconv.Itoa(liveStartReq.LiveDifficultyID)
	liveNotes := utils.ReadAllText("assets/as/notes/" + liveDifficultyId + ".json")
	if liveNotes == "" {
		panic("歌曲情报信息不存在！")
	}

	var liveNotesRes model.AsLiveStageInfo
	if err := json.Unmarshal([]byte(liveNotes), &liveNotesRes); err != nil {
		panic(err)
	}

	if liveStartReq.IsAutoPlay {
		for k := range liveNotesRes.LiveNotes {
			liveNotesRes.LiveNotes[k].AutoJudgeType = 30
		}
	}

	var partnerInfo any
	if cardInfo != "" {
		var info map[string]any
		if err = json.Unmarshal([]byte(cardInfo), &info); err != nil {
			panic(err)
		}
		partnerInfo = info
	} else {
		partnerInfo = nil
	}

	liveStartResp := GetUserData("liveStart.json")
	liveStartResp, _ = sjson.Set(liveStartResp, "live.live_id", liveId)
	liveStartResp, _ = sjson.Set(liveStartResp, "live.deck_id", liveStartReq.DeckID)
	liveStartResp, _ = sjson.Set(liveStartResp, "live.live_stage", liveNotesRes)
	liveStartResp, _ = sjson.Set(liveStartResp, "live.live_partner_card", partnerInfo)
	liveStartResp, _ = sjson.Set(liveStartResp, "user_model_diff.user_status", GetUserStatus())
	liveStartResp, _ = sjson.Set(liveStartResp, "user_model_diff.user_status.latest_live_deck_id", liveStartReq.DeckID)
	liveStartResp, _ = sjson.Set(liveStartResp, "user_model_diff.user_status.last_live_difficulty_id", liveStartReq.LiveDifficultyID)
	resp := SignResp(ctx.GetString("ep"), liveStartResp, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsLiveFinish(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	var cardMasterId, maxVolt, skillCount, appealCount int64
	liveFinishReq := gjson.Parse(reqBody.String())
	liveFinishReq.Get("live_score.card_stat_dict").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			volt := value.Get("got_voltage").Int()
			if volt > maxVolt {
				maxVolt = volt

				cardMasterId = value.Get("card_master_id").Int()
				skillCount = value.Get("skill_triggered_count").Int()
				appealCount = value.Get("appeal_count").Int()
			}
		}
		return true
	})

	mvpInfo := model.AsMvpInfo{
		CardMasterID:        cardMasterId,
		GetVoltage:          maxVolt,
		SkillTriggeredCount: skillCount,
		AppealCount:         appealCount,
	}

	liveId := liveFinishReq.Get("live_id").String()
	res, err := database.LevelDb.Get([]byte("live_" + liveId))
	CheckErr(err)

	liveStartReq := model.AsLiveStartReq{}
	if err := json.Unmarshal(res, &liveStartReq); err != nil {
		panic(err)
	}
	// fmt.Println("liveStartReq:", liveStartReq)

	var partnerInfo any
	if liveStartReq.PartnerUserID != 0 {
		info := model.AsLivePartnerInfo{
			LastPlayedAt:                        time.Now().Unix(),
			RecommendCardMasterID:               liveStartReq.PartnerCardMasterID,
			RecommendCardLevel:                  1,
			IsRecommendCardImageAwaken:          true,
			IsRecommendCardAllTrainingActivated: true,
			IsNew:                               false,
			FriendApprovedAt:                    nil,
			RequestStatus:                       3,
			IsRequestPending:                    false,
		}
		partnerResp := gjson.Parse(GetUserData("fetchLivePartners.json")).Get("partner_select_state.live_partners")
		partnerResp.ForEach(func(k, v gjson.Result) bool {
			userId := v.Get("user_id").Int()
			if userId == int64(liveStartReq.PartnerUserID) {
				info.UserID = int(userId)
				info.Name.DotUnderText = v.Get("name.dot_under_text").String()
				info.Rank = int(v.Get("rank").Int())
				info.EmblemID = int(v.Get("emblem_id").Int())
				info.IntroductionMessage.DotUnderText = v.Get("introduction_message.dot_under_text").String()
			}
			return true
		})
		partnerInfo = info
	} else {
		partnerInfo = nil
	}

	liveResult := model.AsLiveResultAchievementStatus{
		ClearCount:       1,
		GotVoltage:       liveFinishReq.Get("live_score.current_score").Int(),
		RemainingStamina: liveFinishReq.Get("live_score.remaining_stamina").Int(),
	}

	liveFinishResp := GetUserData("liveFinish.json")
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.live_difficulty_master_id", liveStartReq.LiveDifficultyID)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.live_deck_id", liveStartReq.DeckID)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.mvp", mvpInfo)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.partner", partnerInfo)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.live_result_achievement_status", liveResult)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.voltage", liveFinishReq.Get("live_score.current_score").Int())
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.last_best_voltage", liveFinishReq.Get("live_score.current_score").Int())
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.before_user_exp", GetUserStatus()["exp"].(float64))
	liveFinishResp, _ = sjson.Set(liveFinishResp, "live_result.gain_user_exp", 0)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "user_model_diff.user_status", GetUserStatus())
	liveFinishResp, _ = sjson.Set(liveFinishResp, "user_model_diff.user_status.latest_live_deck_id", liveStartReq.DeckID)
	liveFinishResp, _ = sjson.Set(liveFinishResp, "user_model_diff.user_status.last_live_difficulty_id", liveStartReq.LiveDifficultyID)
	resp := SignResp(ctx.GetString("ep"), liveFinishResp, sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsGetOtherUserCard(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	userCardReq := model.AsUserCardReq{}
	if err := json.Unmarshal([]byte(reqBody.String()), &userCardReq); err != nil {
		panic(err)
	}
	// fmt.Println(liveStartReq)

	var cardInfo string
	partnerResp := gjson.Parse(GetUserData("fetchLivePartners.json")).Get("partner_select_state.live_partners")
	partnerResp.ForEach(func(k, v gjson.Result) bool {
		userId := v.Get("user_id").Int()
		if userId == userCardReq.UserID {
			v.Get("card_by_category").ForEach(func(kk, vv gjson.Result) bool {
				if vv.IsObject() {
					cardId := vv.Get("card_master_id").Int()
					if cardId == userCardReq.CardMasterID {
						cardInfo = vv.String()
						// fmt.Println(cardInfo)
						return false
					}
				}
				return true
			})
			return false
		}
		return true
	})

	var userCardInfo map[string]any
	if err := json.Unmarshal([]byte(cardInfo), &userCardInfo); err != nil {
		panic(err)
	}

	userCardResp := GetUserData("getOtherUserCard.json")
	userCardResp, _ = sjson.Set(userCardResp, "other_user_card", userCardInfo)
	resp := SignResp(ctx.GetString("ep"), userCardResp, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchNoticeDetail(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	noticeId := reqBody.Get("notice_id").String()

	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/notices/"+noticeId+".json"), sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsChangeIsAwakeningImage(ctx *gin.Context) {
	reqBody := gjson.Parse(ctx.GetString("reqBody")).Array()[0]
	// fmt.Println(reqBody.String())

	req := model.AsCardAwakeningReq{}
	if err := json.Unmarshal([]byte(reqBody.String()), &req); err != nil {
		panic(err)
	}

	loginData := GetUserData("userCard.json")
	cardInfo := model.AsCardInfo{}
	gjson.Parse(loginData).Get("user_card_by_card_id").
		ForEach(func(key, value gjson.Result) bool {
			if value.IsObject() {
				if err := json.Unmarshal([]byte(value.String()), &cardInfo); err != nil {
					panic(err)
				}

				if cardInfo.CardMasterID == req.CardMasterID {
					cardInfo.IsAwakeningImage = req.IsAwakeningImage

					k := "user_card_by_card_id." + key.String() + ".is_awakening_image"
					SetUserData("userCard.json", k, req.IsAwakeningImage)

					return false
				}
			}
			return true
		})

	userCardInfo := []any{}
	userCardInfo = append(userCardInfo, cardInfo.CardMasterID)
	userCardInfo = append(userCardInfo, cardInfo)

	cardResp := GetUserData("changeIsAwakeningImage.json")
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_status", GetUserStatus())
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_card_by_card_id", userCardInfo)
	resp := SignResp(ctx.GetString("ep"), cardResp, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishStory(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishStory.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishStoryMain(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishUserStoryMain.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishStoryLinkage(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("finishStoryLinkage.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchTrainingTree(ctx *gin.Context) {
	signBody := GetUserData("fetchTrainingTree.json")
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsUpdatePushNotificationSettings(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), "{}", sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsExecuteLesson(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	deckId := req.Get("selected_deck_id").Int()

	var deckInfo string
	var actionList []model.AsLessonMenuAction
	gjson.Parse(GetUserData("lessonDeck.json")).Get("user_lesson_deck_by_id").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() && value.Get("user_lesson_deck_id").Int() == deckId {
			deckInfo = value.String()
			// fmt.Println("Deck Info:", deckInfo)

			gjson.Parse(deckInfo).ForEach(func(kk, vv gjson.Result) bool {
				// fmt.Printf("kk: %s, vv: %s\n", kk.String(), vv.String())
				if strings.Contains(kk.String(), "card_master_id") {
					actionList = append(actionList, model.AsLessonMenuAction{
						CardMasterID:                  vv.Int(),
						Position:                      0,
						IsAddedPassiveSkill:           true,
						IsAddedSpecialPassiveSkill:    true,
						IsRankupedPassiveSkill:        true,
						IsRankupedSpecialPassiveSkill: true,
						IsPromotedSkill:               true,
						MaxRarity:                     4,
						UpCount:                       1,
					})
				}
				return true
			})
			return false
		}
		return true
	})
	// fmt.Println(actionList)

	SetUserData("executeLesson.json", "lesson_menu_actions.1", actionList)
	SetUserData("executeLesson.json", "lesson_menu_actions.3", actionList)
	SetUserData("executeLesson.json", "lesson_menu_actions.5", actionList)
	signBody := SetUserData("executeLesson.json", "lesson_menu_actions.7", actionList)
	SetUserData("userStatus.json", "main_lesson_deck_id", deckId)
	signBody, _ = sjson.Set(signBody, "user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsResultLesson(ctx *gin.Context) {
	userData := GetUserStatus()
	signBody, _ := sjson.Set(GetUserData("resultLesson.json"),
		"user_model_diff.user_status", userData)
	signBody, _ = sjson.Set(signBody, "selected_deck_id", userData["main_lesson_deck_id"])
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSkillEditResult(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]

	var cardList []any
	index := 1
	cardData := GetUserData("userCard.json")
	cardInfo := gjson.Parse(cardData).Get("user_card_by_card_id")
	cardInfo.ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() {
			if index > 9 {
				return false
			}
			// fmt.Println("cardInfo:", value.String())

			skillList := req.Get("selected_skill_ids")
			skillList.ForEach(func(kk, vv gjson.Result) bool {
				if kk.Int()%2 == 0 && vv.Int() == value.Get("card_master_id").Int() {
					skill := skillList.Get(fmt.Sprintf("%d", kk.Int()+1))
					skill.ForEach(func(kkk, vvv gjson.Result) bool {
						skillIdKey := fmt.Sprintf("user_card_by_card_id.%s.additional_passive_skill_%d_id", key.String(), kkk.Int()+1)
						cardData = SetUserData("userCard.json", skillIdKey, vvv.Int())
						return true
					})

					card := gjson.Parse(cardData).Get("user_card_by_card_id." + key.String())
					cardList = append(cardList, card.Get("card_master_id").Int())
					cardList = append(cardList, card.Value())

					index++
				}
				return true
			})
		}
		return true
	})

	signBody := GetUserData("skillEditResult.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_card_by_card_id", cardList)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSaveDeckLesson(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	deckId := req.Get("deck_id").Int()
	lessonDeck := GetUserData("lessonDeck.json")

	var deckInfo string
	var deckIndex string
	gjson.Parse(lessonDeck).Get("user_lesson_deck_by_id").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() && value.Get("user_lesson_deck_id").Int() == deckId {
			deckInfo = value.String()
			deckIndex = key.String()
			// fmt.Println("Lesson Deck:", deckInfo)
			return false
		}
		return true
	})

	cardList := req.Get("card_master_ids")
	cardList.ForEach(func(key, value gjson.Result) bool {
		if key.Int()%2 == 0 {
			position := value.String()
			// fmt.Println("Position:", position)

			cardMasterId := cardList.Get(fmt.Sprintf("%d", key.Int()+1)).Int()
			// fmt.Println("Card:", cardMasterId)

			deckInfo, _ = sjson.Set(deckInfo, "card_master_id_"+position, cardMasterId)
			// fmt.Println("New Lesson Deck:", deckInfo)

			SetUserData("lessonDeck.json", "user_lesson_deck_by_id."+deckIndex, gjson.Parse(deckInfo).Value())
			// lessonDeck, _ = sjson.Set(lessonDeck, "user_lesson_deck_by_id."+deckIndex, gjson.Parse(deckInfo).Value())
		}
		return true
	})

	signBody := GetUserData("saveDeckLesson.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_lesson_deck_by_id.0", deckId)
	signBody, _ = sjson.Set(signBody, "user_model.user_lesson_deck_by_id.1", gjson.Parse(deckInfo).Value())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSaveSuit(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	deckId := req.Get("deck_id").Int()
	cardId := req.Get("card_index").Int()
	suitId := req.Get("suit_master_id").Int()

	deckIndex := deckId*2 - 1
	keyLiveDeck := fmt.Sprintf("user_live_deck_by_id.%d", deckIndex)
	// fmt.Println("keyLiveDeck:", keyLiveDeck)
	liveDeck := gjson.Parse(GetUserData("liveDeck.json")).Get(keyLiveDeck).String()
	// fmt.Println(liveDeck)
	keyLiveDeckInfo := fmt.Sprintf("suit_master_id_%d", cardId)
	liveDeck, _ = sjson.Set(liveDeck, keyLiveDeckInfo, suitId)
	// fmt.Println(liveDeck)

	var deckInfo model.AsDeckInfo
	if err := json.Unmarshal([]byte(liveDeck), &deckInfo); err != nil {
		panic(err)
	}

	SetUserData("liveDeck.json", keyLiveDeck, deckInfo)

	signBody, _ := sjson.Set(GetUserData("saveSuit.json"),
		"user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_live_deck_by_id.0", deckId)
	signBody, _ = sjson.Set(signBody, "user_model.user_live_deck_by_id.1", deckInfo)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSaveDeck(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	// fmt.Println(reqBody)

	req := gjson.Parse(reqBody).Array()[0]
	deckId := req.Get("deck_id")
	// fmt.Println("deckId:", deckId)

	position := req.Get("card_master_ids.0")
	cardMasterId := req.Get("card_master_ids.1")
	// fmt.Println("cardMasterId:", cardMasterId)

	var deckInfo, partyInfo string
	var oldCardMasterId int64
	var partyId int64
	var savePartyInfo model.AsPartyInfo
	deckList := GetUserData("liveDeck.json")
	gjson.Parse(deckList).Get("user_live_deck_by_id").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() && value.Get("user_live_deck_id").String() == deckId.String() {
			deckInfo = value.String()
			// fmt.Println("deckInfo:", deckInfo)

			oldCardMasterId = gjson.Parse(deckInfo).Get("card_master_id_" + position.String()).Int()
			deckInfo, _ = sjson.Set(deckInfo, "card_master_id_"+position.String(), cardMasterId.Int())
			deckInfo, _ = sjson.Set(deckInfo, "suit_master_id_"+position.String(), cardMasterId.Int())
			// fmt.Println("New deckInfo:", deckInfo)

			SetUserData("liveDeck.json", "user_live_deck_by_id."+key.String(), gjson.Parse(deckInfo).Value())

			return false
		}
		return true
	})
	gjson.Parse(deckList).Get("user_live_party_by_id").ForEach(func(key, value gjson.Result) bool {
		if value.IsObject() && (value.Get("party_id").String() == deckId.String()+"01" ||
			value.Get("party_id").String() == deckId.String()+"02" ||
			value.Get("party_id").String() == deckId.String()+"03") {
			value.ForEach(func(kk, vv gjson.Result) bool {
				if vv.Int() == oldCardMasterId {
					partyInfo = value.String()
					// fmt.Println("partyInfo:", partyInfo)

					partyInfo, _ = sjson.Set(partyInfo, kk.String(), cardMasterId.Int())
					// fmt.Println("New partyInfo:", partyInfo)

					newPartyInfo := gjson.Parse(partyInfo)
					partyId = newPartyInfo.Get("party_id").Int()

					roleIds := []int{}
					err := MainEng.Table("m_card").
						Where("id IN (?,?,?)", newPartyInfo.Get("card_master_id_1").Int(),
							newPartyInfo.Get("card_master_id_2").Int(),
							newPartyInfo.Get("card_master_id_3").Int()).
						Cols("role").Find(&roleIds)
					CheckErr(err)
					// fmt.Println("roleIds:", roleIds)

					partyIcon, partyName := GetPartyInfoByRoleIds(roleIds)
					realPartyName := GetRealPartyName(partyName)
					partyInfo, _ = sjson.Set(partyInfo, "name.dot_under_text", realPartyName)
					partyInfo, _ = sjson.Set(partyInfo, "icon_master_id", partyIcon)
					// fmt.Println("New partyInfo 2:", partyInfo)

					decoder := json.NewDecoder(strings.NewReader(partyInfo))
					decoder.UseNumber()
					err = decoder.Decode(&savePartyInfo)
					CheckErr(err)
					SetUserData("liveDeck.json", "user_live_party_by_id."+key.String(), savePartyInfo)

					return false
				}
				return true
			})
		}
		return true
	})

	signBody := GetUserData("SaveDeck.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_live_deck_by_id.0", deckId.Int())
	signBody, _ = sjson.Set(signBody, "user_model.user_live_deck_by_id.1", gjson.Parse(deckInfo).Value())
	signBody, _ = sjson.Set(signBody, "user_model.user_live_party_by_id.0", partyId)
	signBody, _ = sjson.Set(signBody, "user_model.user_live_party_by_id.1", savePartyInfo)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)
	// fmt.Println(resp)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSetFavoriteMember(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	SetUserData("userStatus.json", "favorite_member_id",
		gjson.Parse(reqBody).Array()[0].Get("member_master_id").Int())
	signBody, _ := sjson.Set(GetUserData("setFavoriteMember.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchMission(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("fetchMission.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsClearMissionBadge(ctx *gin.Context) {
	signBody, _ := sjson.Set(GetUserData("clearMissionBadge.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchPresent(ctx *gin.Context) {
	signBody := GetUserData("fetchPresent.json")
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

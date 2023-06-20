package handler

import (
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

	loginBody := utils.ReadAllText("assets/as/login.json")
	loginBody, _ = sjson.Set(loginBody, "session_key", newKey64)
	loginBody, _ = sjson.Set(loginBody, "user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), loginBody, config.StartUpKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchBootstrap(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/fetchBootstrap.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchBillingHistory(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchBillingHistory.json"), sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchNotice(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchNotice.json"), sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsGetPackUrl(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var req []model.AsReq
	err := json.Unmarshal([]byte(reqBody), &req)
	if err != nil {
		panic(err)
	}
	// fmt.Println(req)

	packBody, ok := req[0].(map[string]any)
	if !ok {
		panic("Assertion failed!")
	}
	// fmt.Println(packBody)

	packNames, ok := packBody["pack_names"].([]any)
	if !ok {
		panic("Assertion failed!")
	}

	// 生成更新包 map
	var packageList []string
	var urlList []string

	err = json.Unmarshal([]byte(utils.ReadAllText("assets/as/packages.json")), &packageList)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(utils.ReadAllText("assets/as/urls.json")), &urlList)
	if err != nil {
		panic(err)
	}

	if len(packageList) != len(urlList) {
		fmt.Println("File size not match!")
		return
	}

	packageUrls := map[string]string{}
	for k, p := range packageList {
		packageUrls[p] = urlList[k]
	}

	// Response
	var respUrls []string
	for _, pack := range packNames {
		packName, ok := pack.(string)
		if !ok {
			panic("Assertion failed!")
		}
		// fmt.Println(packageUrls[packName])
		respUrls = append(respUrls, packageUrls[packName])
	}

	urlResp := model.PackUrlRespBody{
		UrlList: respUrls,
	}

	var resp []model.AsResp
	resp = append(resp, time.Now().UnixMilli()) // 时间戳
	resp = append(resp, config.MasterVersion)   // 版本号
	resp = append(resp, 0)                      // 固定值
	resp = append(resp, urlResp)                // 数据体

	mm, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(mm))

	signBody := mm[1 : len(mm)-1]
	// fmt.Println(string(signBody))

	ep := strings.ReplaceAll(ctx.Request.URL.String(), "/ep3071", "")
	// fmt.Println(ep)

	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+string(signBody)), []byte(sessionKey))

	resp = append(resp, sign)
	mm, err = json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(mm))

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, string(mm))
}

func AsUpdateCardNewFlag(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/updateCardNewFlag.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsGetClearedPlatformAchievement(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/getClearedPlatformAchievement.json"), sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchLiveMusicSelect(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/fetchLiveMusicSelect.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsLiveMvStart(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/liveMvStart.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsLiveMvSaveDeck(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")

	var req []model.AsReq
	err := json.Unmarshal([]byte(reqBody), &req)
	if err != nil {
		panic(err)
	}

	body, ok := req[0].(map[string]any)
	if !ok {
		panic("Assertion failed!")
	}

	reqB, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(reqB))

	saveReq := model.LiveSaveDeckReq{}
	err = json.Unmarshal(reqB, &saveReq)
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

	var userLiveMvDeckCustomByID []model.UserLiveMvDeckCustomByID
	userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, saveReq.LiveMasterID)
	userLiveMvDeckCustomByID = append(userLiveMvDeckCustomByID, userLiveMvDeckInfo)
	// fmt.Println(userLiveMvDeckCustomByID)

	viewStatusIds := map[int]int{}
	for k, v := range saveReq.ViewStatusByPos {
		if k%2 == 0 {
			viewStatusId := saveReq.ViewStatusByPos[k+1]
			viewStatusIds[v] = viewStatusId
		}
	}
	// fmt.Println(viewStatusIds)

	var userMemberByMemberID []model.UserMemberByMemberID
	for k, v := range memberIds {
		userMemberByMemberID = append(userMemberByMemberID, v)
		userMemberByMemberID = append(userMemberByMemberID, model.UserMemberInfo{
			MemberMasterID:           v,
			CustomBackgroundMasterID: 103506600,
			SuitMasterID:             suitIds[k],
			LovePoint:                0,
			LovePointLimit:           999999,
			LoveLevel:                1,
			ViewStatus:               viewStatusIds[k],
			IsNew:                    true,
		})
	}
	// fmt.Println(userMemberByMemberID)

	saveResp := model.LiveSaveDeckResp{
		UserModel: model.UserModel{
			UserStatus:                                              CommonUserStatus(),
			UserMemberByMemberID:                                    userMemberByMemberID,
			UserCardByCardID:                                        []any{},
			UserSuitBySuitID:                                        []any{},
			UserLiveDeckByID:                                        []any{},
			UserLivePartyByID:                                       []any{},
			UserLessonDeckByID:                                      []any{},
			UserLiveMvDeckByID:                                      []any{},
			UserLiveMvDeckCustomByID:                                userLiveMvDeckCustomByID,
			UserLiveDifficultyByDifficultyID:                        []any{},
			UserStoryMainByStoryMainID:                              []any{},
			UserStoryMainSelectedByStoryMainCellID:                  []any{},
			UserVoiceByVoiceID:                                      []any{},
			UserEmblemByEmblemID:                                    []any{},
			UserGachaTicketByTicketID:                               []any{},
			UserGachaPointByPointID:                                 []any{},
			UserLessonEnhancingItemByItemID:                         []any{},
			UserTrainingMaterialByItemID:                            []any{},
			UserGradeUpItemByItemID:                                 []any{},
			UserCustomBackgroundByID:                                []any{},
			UserStorySideByID:                                       []any{},
			UserStoryMemberByID:                                     []any{},
			UserCommunicationMemberDetailBadgeByID:                  []any{},
			UserStoryEventHistoryByID:                               []any{},
			UserRecoveryLpByID:                                      []any{},
			UserRecoveryApByID:                                      []any{},
			UserMissionByMissionID:                                  []any{},
			UserDailyMissionByMissionID:                             []any{},
			UserWeeklyMissionByMissionID:                            []any{},
			UserInfoTriggerBasicByTriggerID:                         []any{},
			UserInfoTriggerCardGradeUpByTriggerID:                   []any{},
			UserInfoTriggerMemberGuildSupportItemExpiredByTriggerID: []any{},
			UserInfoTriggerMemberLoveLevelUpByTriggerID:             []any{},
			UserAccessoryByUserAccessoryID:                          []any{},
			UserAccessoryLevelUpItemByID:                            []any{},
			UserAccessoryRarityUpItemByID:                           []any{},
			UserUnlockScenesByEnum:                                  []any{},
			UserSceneTipsByEnum:                                     []any{},
			UserRuleDescriptionByID:                                 []any{},
			UserExchangeEventPointByID:                              []any{},
			UserSchoolIdolFestivalIDRewardMissionByID:               []any{},
			UserGpsPresentReceivedByID:                              []any{},
			UserEventMarathonByEventMasterID:                        []any{},
			UserEventMiningByEventMasterID:                          []any{},
			UserEventCoopByEventMasterID:                            []any{},
			UserLiveSkipTicketByID:                                  []any{},
			UserStoryEventUnlockItemByID:                            []any{},
			UserEventMarathonBoosterByID:                            []any{},
			UserReferenceBookByID:                                   []any{},
			UserReviewRequestProcessFlowByID:                        []any{},
			UserRankExpByID:                                         []any{},
			UserShareByID:                                           []any{},
			UserTowerByTowerID:                                      []any{},
			UserRecoveryTowerCardUsedCountItemByRecoveryTowerCardUsedCountItemMasterID: []any{},
			UserStoryLinkageByID:             []any{},
			UserSubscriptionStatusByID:       []any{},
			UserStoryMainPartDigestMovieByID: []any{},
			UserMemberGuildByID:              []any{},
			UserMemberGuildSupportItemByID:   []any{},
			UserDailyTheaterByDailyTheaterID: []any{},
			UserPlayListByID:                 []any{},
		},
	}
	respB, err := json.Marshal(saveResp)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(b))

	resp := SignResp(ctx.GetString("ep"), string(respB), sessionKey)
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

	signBody := utils.ReadAllText("assets/as/fetchCommunicationMemberDetail.json")
	signBody, _ = sjson.Set(signBody, "member_love_panels.0.member_id", memberId)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsTapLovePoint(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/tapLovePoint.json"),
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

	signBody := utils.ReadAllText("assets/as/updateUserCommunicationMemberDetailBadge.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_communication_member_detail_badge_by_id", userDetail)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsUpdateUserLiveDifficultyNewFlag(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/updateUserLiveDifficultyNewFlag.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishUserStorySide(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/finishUserStorySide.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishUserStoryMember(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/finishUserStoryMember.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSetTheme(ctx *gin.Context) {
	reqBody := ctx.GetString("reqBody")
	var memberMasterId, suitMasterId, backgroundMasterId int64
	gjson.Parse(reqBody).ForEach(func(key, value gjson.Result) bool {
		if value.Get("member_master_id").String() != "" {
			memberMasterId = value.Get("member_master_id").Int()
			suitMasterId = value.Get("suit_master_id").Int()
			backgroundMasterId = value.Get("custom_background_master_id").Int()
			return false
		}
		return true
	})

	userMemberRes := []model.SetThemeRes{}
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

	userSuitRes := []model.SetThemeRes{}
	userSuitRes = append(userSuitRes, suitMasterId)
	userSuitRes = append(userSuitRes, model.AsSuitInfo{
		SuitMasterID: int(suitMasterId),
		IsNew:        false,
	})

	signBody := utils.ReadAllText("assets/as/setTheme.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_member_by_member_id", userMemberRes)
	signBody, _ = sjson.Set(signBody, "user_model.user_suit_by_suit_id", userSuitRes)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchProfile(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchProfile.json"), sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchEmblem(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/fetchEmblem.json"),
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
			return false
		}
		return true
	})

	signBody := utils.ReadAllText("assets/as/activateEmblem.json")
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_status", GetUserStatus())
	signBody, _ = sjson.Set(signBody, "user_model.user_status.emblem_id", emblemId)
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsSaveUserNaviVoice(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/saveUserNaviVoice.json"),
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
			DotUnderText: "队伍名称变了?",
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

	deckInfoRes := []model.AsResp{}
	deckInfoRes = append(deckInfoRes, req.DeckID)
	deckInfoRes = append(deckInfoRes, deckInfo)

	partyInfoRes := []model.AsResp{}
	for k, v := range req.SquadDict {
		if k%2 == 0 {
			partyId, err := v.(json.Number).Int64()
			CheckErr(err)
			// fmt.Println("Party ID:", partyId)

			rDictInfo, err := json.Marshal(req.SquadDict[k+1])
			CheckErr(err)

			dictInfo := model.AsDeckSquadDict{}
			if err = json.Unmarshal(rDictInfo, &dictInfo); err != nil {
				panic(err)
			}
			// fmt.Println("Party Info:", dictInfo)

			roleIds := []int{}
			err = MainEng.Table("m_card").
				Where("id IN (?,?,?)", dictInfo.CardMasterIds[0], dictInfo.CardMasterIds[1], dictInfo.CardMasterIds[2]).
				Cols("role").Find(&roleIds)
			CheckErr(err)
			// fmt.Println("roleIds:", roleIds)

			var partyIcon int
			var partyName string
			// 脑残逻辑部分
			exists, err := MainEng.Table("m_live_party_name").
				Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[0], roleIds[1], roleIds[2]).
				Cols("name,live_party_icon").Get(&partyName, &partyIcon)
			CheckErr(err)
			if !exists {
				exists, err = MainEng.Table("m_live_party_name").
					Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[0], roleIds[2], roleIds[1]).
					Cols("name,live_party_icon").Get(&partyName, &partyIcon)
				CheckErr(err)
				if !exists {
					exists, err = MainEng.Table("m_live_party_name").
						Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[1], roleIds[0], roleIds[2]).
						Cols("name,live_party_icon").Get(&partyName, &partyIcon)
					CheckErr(err)
					if !exists {
						exists, err = MainEng.Table("m_live_party_name").
							Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[1], roleIds[2], roleIds[0]).
							Cols("name,live_party_icon").Get(&partyName, &partyIcon)
						CheckErr(err)
						if !exists {
							exists, err = MainEng.Table("m_live_party_name").
								Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[2], roleIds[0], roleIds[1]).
								Cols("name,live_party_icon").Get(&partyName, &partyIcon)
							CheckErr(err)
							if !exists {
								exists, err = MainEng.Table("m_live_party_name").
									Where("role_1 = ? AND role_2 = ? AND role_3 = ?", roleIds[2], roleIds[1], roleIds[0]).
									Cols("name,live_party_icon").Get(&partyName, &partyIcon)
								CheckErr(err)
								if !exists {
									panic("Fuck you!")
								}
							}
						}
					}
				}
			}

			var realPartyName string
			_, err = MainEng.Table("m_dictionary").Where("id = ?", strings.ReplaceAll(partyName, "k.", "")).Cols("message").Get(&realPartyName)
			CheckErr(err)

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

			partyInfoRes = append(partyInfoRes, partyId)
			partyInfoRes = append(partyInfoRes, partyInfo)
		}
	}

	respBody := utils.ReadAllText("assets/as/saveDeckAll.json")
	respBody, _ = sjson.Set(respBody, "user_model.user_status", GetUserStatus())
	respBody, _ = sjson.Set(respBody, "user_model.user_live_deck_by_id", deckInfoRes)
	respBody, _ = sjson.Set(respBody, "user_model.user_live_party_by_id", partyInfoRes)
	resp := SignResp(ctx.GetString("ep"), respBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchLivePartners(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchLivePartners.json"), sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchLiveDeckSelect(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchLiveDeckSelect.json"), sessionKey)

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
	partnerResp := gjson.Parse(utils.ReadAllText("assets/as/fetchLivePartners.json")).Get("partner_select_state.live_partners")
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

	var liveNotesRes map[string]any
	if err = json.Unmarshal([]byte(liveNotes), &liveNotesRes); err != nil {
		panic(err)
	}

	var partnerInfo map[string]any
	if err = json.Unmarshal([]byte(cardInfo), &partnerInfo); err != nil {
		panic(err)
	}

	liveStartResp := utils.ReadAllText("assets/as/liveStart.json")
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

	partnerInfo := model.AsLivePartnerInfo{
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
	partnerResp := gjson.Parse(utils.ReadAllText("assets/as/fetchLivePartners.json")).Get("partner_select_state.live_partners")
	partnerResp.ForEach(func(k, v gjson.Result) bool {
		userId := v.Get("user_id").Int()
		if userId == int64(liveStartReq.PartnerUserID) {
			partnerInfo.UserID = int(userId)
			partnerInfo.Name.DotUnderText = v.Get("name.dot_under_text").String()
			partnerInfo.Rank = int(v.Get("rank").Int())
			partnerInfo.EmblemID = int(v.Get("emblem_id").Int())
			partnerInfo.IntroductionMessage.DotUnderText = v.Get("introduction_message.dot_under_text").String()
		}
		return true
	})

	liveResult := model.AsLiveResultAchievementStatus{
		ClearCount:       1,
		GotVoltage:       liveFinishReq.Get("live_score.current_score").Int(),
		RemainingStamina: liveFinishReq.Get("live_score.remaining_stamina").Int(),
	}

	liveFinishResp := utils.ReadAllText("assets/as/liveFinish.json")
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
	partnerResp := gjson.Parse(utils.ReadAllText("assets/as/fetchLivePartners.json")).Get("partner_select_state.live_partners")
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

	userCardResp := utils.ReadAllText("assets/as/getOtherUserCard.json")
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

	cardInfo := model.AsCardInfo{}
	gjson.Parse(utils.ReadAllText("assets/as/login.json")).Get("user_model.user_card_by_card_id").
		ForEach(func(key, value gjson.Result) bool {
			if value.IsObject() {
				if err := json.Unmarshal([]byte(value.String()), &cardInfo); err != nil {
					panic(err)
				}

				if cardInfo.CardMasterID == req.CardMasterID {
					cardInfo.IsAwakeningImage = req.IsAwakeningImage

					return false
				}
			}
			return true
		})

	userCardInfo := []any{}
	userCardInfo = append(userCardInfo, cardInfo.CardMasterID)
	userCardInfo = append(userCardInfo, cardInfo)

	cardResp := utils.ReadAllText("assets/as/changeIsAwakeningImage.json")
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_status", GetUserStatus())
	cardResp, _ = sjson.Set(cardResp, "user_model_diff.user_card_by_card_id", userCardInfo)
	resp := SignResp(ctx.GetString("ep"), cardResp, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishStory(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/finishStory.json"),
		"user_model.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishStoryMain(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/finishUserStoryMain.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFinishStoryLinkage(ctx *gin.Context) {
	signBody, _ := sjson.Set(utils.ReadAllText("assets/as/finishStoryLinkage.json"),
		"user_model_diff.user_status", GetUserStatus())
	resp := SignResp(ctx.GetString("ep"), signBody, sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

func AsFetchTrainingTree(ctx *gin.Context) {
	resp := SignResp(ctx.GetString("ep"), utils.ReadAllText("assets/as/fetchTrainingTree.json"), sessionKey)

	ctx.Header("Content-Type", "application/json")
	ctx.String(http.StatusOK, resp)
}

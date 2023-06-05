package tools

func init() {
	InitUserData(0)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// cardIds := []int{}
	// err = eng.Table("m_card").Cols("id").OrderBy("id ASC").Find(&cardIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, card := range cardIds {
	// 	cardInfo := model.AsCardInfo{
	// 		CardMasterID:               card,
	// 		Level:                      1,
	// 		Exp:                        0,
	// 		LovePoint:                  0,
	// 		IsFavorite:                 false,
	// 		IsAwakening:                true,
	// 		IsAwakeningImage:           true,
	// 		IsAllTrainingActivated:     false,
	// 		TrainingActivatedCellCount: 0,
	// 		MaxFreePassiveSkill:        1,
	// 		Grade:                      0,
	// 		TrainingLife:               0,
	// 		TrainingAttack:             0,
	// 		TrainingDexterity:          0,
	// 		ActiveSkillLevel:           1,
	// 		PassiveSkillALevel:         1,
	// 		PassiveSkillBLevel:         1,
	// 		PassiveSkillCLevel:         1,
	// 		AdditionalPassiveSkill1ID:  0,
	// 		AdditionalPassiveSkill2ID:  0,
	// 		AdditionalPassiveSkill3ID:  0,
	// 		AdditionalPassiveSkill4ID:  0,
	// 		AcquiredAt:                 int(time.Now().Unix()),
	// 		IsNew:                      false,
	// 	}
	// 	m, err := json.Marshal(cardInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", card, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// packageList := []string{}
	// urlList := []string{}

	// err := json.Unmarshal([]byte(utils.ReadAllText("data/packages.json")), &packageList)
	// CheckErr(err)

	// err = json.Unmarshal([]byte(utils.ReadAllText("data/urls.json")), &urlList)
	// CheckErr(err)

	// if len(packageList) != len(urlList) {
	// 	fmt.Println("File size not match!")
	// 	return
	// }

	// packageUrls := map[string]string{}
	// for k, p := range packageList {
	// 	packageUrls[p] = urlList[k]
	// }

	// // fmt.Println(packageUrls)

	// // packUrlBody := model.PackUrlBody{
	// // 	PackNames: []string{
	// // 		"cxh9rj",
	// // 		"ib3f4p",
	// // 		"iwjqao",
	// // 	},
	// // }

	// // hash := "asdfgh"

	// // asReq := []model.AsReq{}
	// // asReq = append(asReq, packUrlBody)
	// // asReq = append(asReq, hash)
	// // mm, err := json.Marshal(asReq)
	// // CheckErr(err)
	// // fmt.Println(string(mm))

	// jsonStr := utils.ReadAllText("data/api1.json")
	// req := []model.AsReq{}
	// err = json.Unmarshal([]byte(jsonStr), &req)
	// CheckErr(err)
	// // fmt.Println(req)

	// packBody, ok := req[0].(map[string]interface{})
	// if !ok {
	// 	panic("Assertion failed!")
	// }
	// // fmt.Println(packBody)

	// packNames, ok := packBody["pack_names"].([]interface{})
	// if !ok {
	// 	panic("Assertion failed!")
	// }

	// respUrls := []string{}
	// for _, pack := range packNames {
	// 	packName, ok := pack.(string)
	// 	if !ok {
	// 		panic("Assertion failed!")
	// 	}
	// 	fmt.Println(packageUrls[packName])
	// 	respUrls = append(respUrls, packageUrls[packName])
	// }

	// urlResp := model.PackUrlRespBody{
	// 	UrlList: respUrls,
	// }

	// resp := []model.AsResp{}
	// resp = append(resp, time.Now().UnixMilli()) // 时间戳
	// resp = append(resp, config.MasterVersion)   // 版本号
	// resp = append(resp, 0)                      // 固定值
	// resp = append(resp, urlResp)                // 数据体

	// mm, err := json.Marshal(resp)
	// CheckErr(err)
	// // fmt.Println(string(mm))

	// signBody := mm[1 : len(mm)-1]
	// fmt.Println(string(signBody))

	// sessionKey := "12345678123456781234567812345678"
	// sign := encrypt.HMAC_SHA1_Encrypt(signBody, []byte(sessionKey))

	// resp = append(resp, sign)
	// mm, err = json.Marshal(resp)
	// CheckErr(err)
	// fmt.Println(string(mm))

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// suitIds := []int{}
	// err = eng.Table("m_suit").Cols("id").OrderBy("id ASC").Find(&suitIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, suit := range suitIds {
	// 	suitInfo := model.AsSuitInfo{
	// 		SuitMasterID: suit,
	// 		IsNew:        false,
	// 	}
	// 	m, err := json.Marshal(suitInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", suit, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// emblemIds := []int{}
	// err = eng.Table("m_emblem").Cols("id").OrderBy("id ASC").Find(&emblemIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, emblem := range emblemIds {
	// 	emblemInfo := model.AsEmblemInfo{
	// 		EmblemMID:  emblem,
	// 		AcquiredAt: time.Now().Unix(),
	// 	}
	// 	m, err := json.Marshal(emblemInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", emblem, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// emblemIds := []int{}
	// err = eng.Table("m_emblem").Cols("id").OrderBy("id ASC").Find(&emblemIds)
	// CheckErr(err)

	// ids := []model.AsEmblemId{}
	// for _, id := range emblemIds {
	// 	ids = append(ids, model.AsEmblemId{
	// 		EmblemMasterID: id,
	// 		IsNew:          false,
	// 	})
	// }
	// m, err := json.Marshal(ids)
	// CheckErr(err)
	// fmt.Println(string(m))

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// difficultyIds := []int{}
	// err = eng.Table("m_live_difficulty").Cols("live_difficulty_id").OrderBy("live_difficulty_id ASC").Find(&difficultyIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, id := range difficultyIds {
	// 	difficultyInfo := model.AsLiveDifficultyInfo{
	// 		LiveDifficultyID:              id,
	// 		MaxScore:                      0,
	// 		MaxCombo:                      0,
	// 		PlayCount:                     0,
	// 		ClearCount:                    0,
	// 		CancelCount:                   0,
	// 		NotClearedCount:               0,
	// 		IsFullCombo:                   false,
	// 		ClearedDifficultyAchievement1: nil,
	// 		ClearedDifficultyAchievement2: nil,
	// 		ClearedDifficultyAchievement3: nil,
	// 		EnableAutoplay:                false,
	// 		IsAutoplay:                    false,
	// 		IsNew:                         false,
	// 	}
	// 	m, err := json.Marshal(difficultyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", id, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// storyIds := []int{}
	// err = eng.Table("m_story_main_cell").Cols("id").OrderBy("id ASC").Find(&storyIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, id := range storyIds {
	// 	storyInfo := model.AsMainStoryInfo{
	// 		StoryMainMasterID: id,
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", id, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// voiceIds := []int{}
	// err = eng.Table("m_navi_voice").Cols("id").OrderBy("id ASC").Find(&voiceIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, id := range voiceIds {
	// 	storyInfo := model.AsNaviVoiceInfo{
	// 		NaviVoiceMasterID: id,
	// 		IsNew:             false,
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", id, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// backgroundIds := []int{}
	// err = eng.Table("m_custom_background").Cols("id").OrderBy("id ASC").Find(&backgroundIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, id := range backgroundIds {
	// 	storyInfo := model.AsCustomBackgroundInfo{
	// 		CustomBackgroundMasterID: id,
	// 		IsNew:                    false,
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", id, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// storyIds := []int{}
	// err = eng.Table("m_story_side").Cols("id").OrderBy("id ASC").Find(&storyIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, id := range storyIds {
	// 	storyInfo := model.AsStorySideInfo{
	// 		StorySideMasterID: id,
	// 		IsNew:             false,
	// 		AcquiredAt:        time.Now().Unix(),
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", id, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// storyIds := []int{}
	// err = eng.Table("m_story_member").Cols("id").OrderBy("id ASC").Find(&storyIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, id := range storyIds {
	// 	storyInfo := model.AsStoryMemberInfo{
	// 		StoryMemberMasterID: id,
	// 		IsNew:               false,
	// 		AcquiredAt:          time.Now().Unix(),
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", id, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// eventIds := []int{}
	// err = eng.Table("m_story_event_history_detail").Cols("story_event_id").OrderBy("story_event_id ASC").Find(&eventIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, id := range eventIds {
	// 	storyInfo := model.AsStoryEventInfo{
	// 		StoryEventID: id,
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", id, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// missionRes := []model.AsMissionRes{}
	// err = eng.Table("m_mission").Cols("id,mission_clear_condition_count").Where("term = 3").OrderBy("id ASC").Find(&missionRes)
	// CheckErr(err)

	// jsonStr := "["
	// for _, res := range missionRes {
	// 	storyInfo := model.AsFreeMissionInfo{
	// 		MissionMID:       res.ID,
	// 		IsNew:            false,
	// 		MissionCount:     res.Count,
	// 		IsCleared:        true,
	// 		IsReceivedReward: true,
	// 		NewExpiredAt:     time.Now().Unix(),
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", res.ID, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// missionRes := []model.AsMissionRes{}
	// err = eng.Table("m_mission").Cols("id,mission_clear_condition_count").
	// 	Where("term = 1 AND end_at > ?", time.Now().Unix()).OrderBy("id ASC").Find(&missionRes)
	// CheckErr(err)

	// jsonStr := "["
	// for _, res := range missionRes {
	// 	storyInfo := model.AsDailyMissionInfo{
	// 		MissionMID:        res.ID,
	// 		IsNew:             false,
	// 		MissionStartCount: res.Count,
	// 		MissionCount:      res.Count,
	// 		IsCleared:         true,
	// 		IsReceivedReward:  true,
	// 		ClearedExpiredAt:  time.Now().Unix(),
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", res.ID, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)

	// missionRes := []model.AsMissionRes{}
	// err = eng.Table("m_mission").Cols("id,mission_clear_condition_count").
	// 	Where("term = 2 AND end_at > ?", time.Now().Unix()).OrderBy("id ASC").Find(&missionRes)
	// CheckErr(err)

	// jsonStr := "["
	// for _, res := range missionRes {
	// 	storyInfo := model.AsWeeklyMissionInfo{
	// 		MissionMID:        res.ID,
	// 		IsNew:             false,
	// 		MissionStartCount: res.Count,
	// 		MissionCount:      res.Count,
	// 		IsCleared:         true,
	// 		IsReceivedReward:  true,
	// 		ClearedExpiredAt:  time.Now().Unix(),
	// 		NewExpiredAt:      time.Now().Unix(),
	// 	}
	// 	m, err := json.Marshal(storyInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", res.ID, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// eng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = eng.Ping()
	// CheckErr(err)
	// eng.ShowSQL(true)

	// memberIds := []int{}
	// err = eng.Table("m_member").Cols("id").OrderBy("id ASC").Find(&memberIds)
	// CheckErr(err)

	// jsonStr := "["
	// for _, memberId := range memberIds {
	// 	cellIds := []int{}
	// 	err = eng.Table("m_member_love_panel_cell").
	// 		Join("LEFT", "m_member_love_panel", "m_member_love_panel_cell.member_love_panel_master_id = m_member_love_panel.id").
	// 		Cols("m_member_love_panel_cell.id").Where("m_member_love_panel.member_master_id = ?", memberId).
	// 		OrderBy("m_member_love_panel_cell.id ASC").Find(&cellIds)
	// 	CheckErr(err)

	// 	panelInfo := model.AsMemberLovePanelInfo{
	// 		MemberID:               memberId,
	// 		MemberLovePanelCellIds: cellIds,
	// 	}

	// 	m, err := json.Marshal(panelInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%s,", string(m))
	// }

	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)

	// mEng, err := xorm.NewEngine("sqlite", "assets/masterdata.db")
	// CheckErr(err)
	// err = mEng.Ping()
	// CheckErr(err)
	// defer mEng.Close()

	// dEng, err := xorm.NewEngine("sqlite", "assets/dictionary_zh_k.db")
	// CheckErr(err)
	// err = dEng.Ping()
	// CheckErr(err)
	// defer dEng.Close()

	// cardRes := []model.AsCardRes{}
	// err = mEng.Table("m_card").Cols("id,card_rarity_type,max_passive_skill_slot").OrderBy("id ASC").Find(&cardRes)
	// CheckErr(err)

	// jsonStr := "["
	// for _, card := range cardRes {
	// 	// 绊板等级加成
	// 	cardLevel := 0
	// 	cardCellCount := 0
	// 	if card.CardRarityType == 10 {
	// 		cardLevel = 40 + 42
	// 		cardCellCount = 61
	// 	} else if card.CardRarityType == 20 {
	// 		cardLevel = 60 + 24
	// 		cardCellCount = 75
	// 	} else if card.CardRarityType == 30 {
	// 		cardLevel = 80 + 12
	// 		cardCellCount = 87
	// 	}

	// 	var apBuff, stBuff, teBuff int
	// 	_, err := mEng.Table("m_training_tree_card_param").Where("id = ? AND training_content_type = ?", card.ID, 2).Select("SUM(value)").Get(&stBuff)
	// 	CheckErr(err)
	// 	// fmt.Println(stBuff)

	// 	_, err = mEng.Table("m_training_tree_card_param").Where("id = ? AND training_content_type = ?", card.ID, 3).Select("SUM(value)").Get(&apBuff)
	// 	CheckErr(err)
	// 	// fmt.Println(apBuff)

	// 	_, err = mEng.Table("m_training_tree_card_param").Where("id = ? AND training_content_type = ?", card.ID, 4).Select("SUM(value)").Get(&teBuff)
	// 	CheckErr(err)
	// 	// fmt.Println(teBuff)

	// 	var skillName string
	// 	var skillId int
	// 	_, err = mEng.Table("m_card_passive_skill_original").Where("card_master_id = ? AND skill_level = 5", card.ID).Cols("name").Get(&skillName)
	// 	CheckErr(err)
	// 	skillName = strings.ReplaceAll(skillName, "k.", "")
	// 	// fmt.Println(skillName)

	// 	// dEng.ShowSQL(true)
	// 	condition := "id LIKE '%" + skillName + "%' AND (message LIKE '%表现%同策略%' OR message LIKE '%表现%同属性%') AND message NOT LIKE '%时%'"
	// 	count, err := dEng.Table("m_dictionary").
	// 		Where(condition).
	// 		Count()
	// 	CheckErr(err)
	// 	if count > 0 {
	// 		skillId = 30000507
	// 	} else {
	// 		skillId = 30000482
	// 	}

	// 	var passiveSkillLevel int
	// 	_, err = mEng.Table("m_card_passive_skill_original").Where("card_master_id = ?", card.ID).
	// 		Cols("skill_level").OrderBy("skill_level DESC").Limit(1).Get(&passiveSkillLevel)
	// 	CheckErr(err)

	// 	cardInfo := model.AsCardInfo{
	// 		CardMasterID:               card.ID,
	// 		Level:                      cardLevel,
	// 		Exp:                        0,
	// 		LovePoint:                  0,
	// 		IsFavorite:                 false,
	// 		IsAwakening:                true,
	// 		IsAwakeningImage:           true,
	// 		IsAllTrainingActivated:     true,
	// 		TrainingActivatedCellCount: cardCellCount,
	// 		MaxFreePassiveSkill:        card.MaxPassiveSkillSlot,
	// 		Grade:                      5,
	// 		TrainingLife:               stBuff,
	// 		TrainingAttack:             apBuff,
	// 		TrainingDexterity:          teBuff,
	// 		ActiveSkillLevel:           5,
	// 		PassiveSkillALevel:         passiveSkillLevel,
	// 		PassiveSkillBLevel:         1,
	// 		PassiveSkillCLevel:         1,
	// 		AdditionalPassiveSkill1ID:  skillId,
	// 		AdditionalPassiveSkill2ID:  skillId,
	// 		AdditionalPassiveSkill3ID:  skillId,
	// 		AdditionalPassiveSkill4ID:  skillId,
	// 		AcquiredAt:                 time.Now().Unix(),
	// 		IsNew:                      false,
	// 	}
	// 	m, err := json.Marshal(cardInfo)
	// 	CheckErr(err)

	// 	jsonStr += fmt.Sprintf("%d,%s,", card.ID, string(m))
	// }
	// jsonStr = strings.TrimRight(jsonStr, ",")
	// jsonStr += "]"
	// fmt.Println(jsonStr)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

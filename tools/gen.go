package tools

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/model"
	"honoka-chan/utils"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func GenApi1Data() {
	db, err := sql.Open("sqlite3", "assets/main.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// global
	var centerId int
	var respAll []interface{}

	// live_status_result
	var liveDifficultyId int
	normalLives := []model.NormalLiveStatusList{}
	sql := `SELECT live_difficulty_id FROM normal_live_m ORDER BY live_difficulty_id ASC`
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&liveDifficultyId)
		if err != nil {
			panic(err)
		}

		normalLive := model.NormalLiveStatusList{
			LiveDifficultyID:   liveDifficultyId,
			Status:             1,
			HiScore:            0,
			HiComboCount:       0,
			ClearCnt:           0,
			AchievedGoalIDList: []int{},
		}

		normalLives = append(normalLives, normalLive)
	}

	specialLives := []model.SpecialLiveStatusList{}
	sql = `SELECT live_difficulty_id FROM special_live_m ORDER BY live_difficulty_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err = rows.Scan(&liveDifficultyId)
		if err != nil {
			panic(err)
		}

		specialLive := model.SpecialLiveStatusList{
			LiveDifficultyID:   liveDifficultyId,
			Status:             1,
			HiScore:            0,
			HiComboCount:       0,
			ClearCnt:           0,
			AchievedGoalIDList: []int{},
		}

		specialLives = append(specialLives, specialLive)
	}

	LiveStatusResp := model.LiveStatusResp{
		// _ = model.LiveStatusResp{
		Result: model.LiveStatusResult{
			NormalLiveStatusList:   normalLives,
			SpecialLiveStatusList:  specialLives,
			TrainingLiveStatusList: []model.TrainingLiveStatusList{},
			MarathonLiveStatusList: []interface{}{},
			FreeLiveStatusList:     []interface{}{},
			CanResumeLive:          false,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, LiveStatusResp)

	// live_list_result
	livesList := []model.LiveList{}
	for _, v := range specialLives {
		livesList = append(livesList, model.LiveList{
			LiveDifficultyID: v.LiveDifficultyID,
			StartDate:        "2023-01-01 00:00:00",
			EndDate:          "2037-01-01 00:00:00",
			IsRandom:         false,
		})
	}
	liveListResp := model.LiveScheduleResp{
		// _ = model.LiveScheduleResp{
		Result: model.LiveScheduleResult{
			EventList:              []interface{}{},
			LiveList:               livesList,
			LimitedBonusList:       []interface{}{},
			LimitedBonusCommonList: []model.LimitedBonusCommonList{}, // 特效道具
			RandomLiveList:         []model.RandomLiveList{},         // 随机歌曲
			FreeLiveList:           []interface{}{},
			TrainingLiveList:       []model.TrainingLiveList{}, // 挑战歌曲
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, liveListResp)

	// unit_list_result
	sql = `SELECT unit_id,album_series_id,rarity,rank_min,rank_max,hp_max,default_removable_skill_capacity FROM unit_m WHERE unit_id NOT IN (SELECT unit_id FROM unit_m WHERE unit_type_id IN (SELECT unit_type_id FROM unit_type_m WHERE image_button_asset IS NULL AND (background_color = 'dcdbe3' AND original_attribute_id IS NULL) OR (unit_type_id IN (10,110,127,128,129) OR unit_type_id BETWEEN 131 AND 140))) ORDER BY unit_number ASC;`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}

	unitsData := []model.Active{}
	oId := 3000000000 // 起始卡片 ID，随意设置
	sdt := time.Now().Add(-time.Hour * 24 * 30)

	defaultDeckUnitList := []model.PlayRewardUnitList{}

	for rows.Next() {
		unitData := model.Active{}
		var uid, rit, rank, max_rank, hp_max, skill_capacity int
		var album_series_id interface{}
		err = rows.Scan(&uid, &album_series_id, &rit, &rank, &max_rank, &hp_max, &skill_capacity)
		if err != nil {
			fmt.Println(err)
			continue
		}
		oId++
		sdt = sdt.Add(time.Second * 3)

		unitData.UnitOwningUserID = oId
		unitData.UnitID = uid
		unitData.NextExp = 0
		unitData.Rank = rank
		unitData.MaxRank = max_rank
		unitData.MaxHp = hp_max
		unitData.DisplayRank = 2
		unitData.UnitSkillLevel = 8
		unitData.FavoriteFlag = true
		unitData.IsRankMax = true
		unitData.IsLevelMax = true
		unitData.IsLoveMax = true
		unitData.IsSkillLevelMax = true
		unitData.IsRemovableSkillCapacityMax = true
		unitData.InsertDate = sdt.Local().Format("2006-01-02 03:04:05")
		if rit != 4 {
			unitData.LevelLimitID = 0
			unitData.UnitRemovableSkillCapacity = skill_capacity
			unitData.IsSigned = false
			switch rit {
			case 1:
				// N
				unitData.Exp = 8000
				unitData.Level = 40
				unitData.MaxLevel = 40
				unitData.Love = 50
				unitData.MaxLove = 50
				unitData.UnitSkillExp = 0
				unitData.UnitSkillLevel = 0
				unitData.FavoriteFlag = false
				unitData.IsRankMax = false
				unitData.IsLevelMax = false
				unitData.IsLoveMax = false
				unitData.IsSkillLevelMax = false
				unitData.IsRemovableSkillCapacityMax = false
			case 2:
				// R
				unitData.Exp = 13500
				unitData.Level = 60
				unitData.MaxLevel = 60
				unitData.Love = 200
				unitData.MaxLove = 200
				unitData.UnitSkillExp = 490
			case 3:
				// SR
				unitData.Exp = 36800
				unitData.Level = 80
				unitData.MaxLevel = 80
				unitData.Love = 500
				unitData.MaxLove = 500
				unitData.UnitSkillExp = 4900
			case 5:
				// SSR
				unitData.Exp = 56657
				unitData.Level = 90
				unitData.MaxLevel = 90
				unitData.Love = 750
				unitData.MaxLove = 750
				unitData.UnitSkillExp = 12700
			}
		} else {
			// UR
			unitData.Exp = 79700
			unitData.Level = 100
			unitData.MaxLevel = 100
			unitData.LevelLimitID = 2
			unitData.Love = 1000
			unitData.MaxLove = 1000
			unitData.UnitSkillExp = 29900
			unitData.UnitRemovableSkillCapacity = 8

			rs, err := db.Query("SELECT COUNT(*) AS ct FROM unit_sign_asset_m WHERE unit_id = ?", uid)
			if err != nil {
				panic(err)
			}

			ct := 0
			for rs.Next() {
				err = rs.Scan(&ct)
				if err != nil {
					panic(err)
				}
			}
			if ct > 0 {
				unitData.IsSigned = true
			} else {
				unitData.IsSigned = false
			}
		}

		unitsData = append(unitsData, unitData)

		// 中心位成员
		if uid == 3927 { // AC15 果
			centerId = oId
		}

		switch album_series_id := album_series_id.(type) {
		case int64:
			if album_series_id == 615 {
				// 仆光套
				defaultDeckUnitList = append(defaultDeckUnitList, model.PlayRewardUnitList{
					UnitOwningUserID: oId,
					UnitID:           uid,
					Level:            100,
					LevelLimitID:     2,
					DisplayRank:      2,
					Love:             1000,
					UnitSkillLevel:   8,
					IsRankMax:        true,
					IsLoveMax:        true,
					IsLevelMax:       true,
					IsSigned:         unitData.IsSigned,
					BeforeLove:       1000,
					MaxLove:          1000,
				})
			}
		}
	}

	unitListResp := model.UnitAllResp{
		// _ = model.UnitAllResp{
		Result: model.UnitAllResult{
			Active:  unitsData,
			Waiting: []model.Waiting{},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, unitListResp)

	// unit_deck_result
	oUids := []model.UnitOwningUserIds{}
	for k, v := range defaultDeckUnitList[0:9] {
		position := k + 1
		defaultDeckUnitList[k].Position = position
		oUids = append(oUids, model.UnitOwningUserIds{
			Position:         position,
			UnitOwningUserID: v.UnitOwningUserID,
		})
	}
	unitDeckResp := model.UnitDeckInfoResp{
		// _ = model.UnitDeckInfoResp{
		Result: []model.UnitDeckInfo{
			{
				UnitDeckID:        1,
				MainFlag:          true,
				DeckName:          "Future style",
				UnitOwningUserIds: oUids,
			},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, unitDeckResp)

	// HACK 保存卡组信息
	deck, err := json.Marshal(defaultDeckUnitList)
	if err != nil {
		panic(err)
	}
	err = database.LevelDb.Put([]byte("deck_info"), deck)
	if err != nil {
		panic(err)
	}

	// unit_support_result
	unitSupportResp := model.UnitSupportResp{
		// _ = model.UnitSupportResp{
		Result: model.UnitSupportResult{
			UnitSupportList: []model.UnitSupportList{},
		}, // 练习道具
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, unitSupportResp)

	// owning_equip_result
	rmSkillResp := model.RemovableSkillResp{
		// _ = model.RemovableSkillResp{
		Result: model.RemovableSkillResult{
			OwningInfo:    []model.OwningInfo{},
			EquipmentInfo: []interface{}{},
		}, // 宝石
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, rmSkillResp)

	// costume_list_result
	costumeListResp := model.CostumeListResp{
		// _ = model.CostumeListResp{
		Result: model.CostumeListResult{
			CostumeList: []model.CostumeList{},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, costumeListResp)

	// album_unit_result
	albumLists := []model.AlbumResult{}
	sql = `SELECT unit_id,rarity FROM unit_m ORDER BY unit_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		albumList := model.AlbumResult{
			RankMaxFlag:      true,
			LoveMaxFlag:      true,
			RankLevelMaxFlag: true,
			AllMaxFlag:       true,
			FavoritePoint:    1000,
		}
		var uid, rit int
		err = rows.Scan(&uid, &rit)
		if err != nil {
			panic(err)
		}
		albumList.UnitID = uid
		if rit != 4 {
			albumList.SignFlag = false
			if rit == 1 {
				albumList.HighestLovePerUnit = 50
				albumList.TotalLove = 50
			} else if rit == 2 {
				albumList.HighestLovePerUnit = 200
				albumList.TotalLove = 200
			} else if rit == 3 {
				albumList.HighestLovePerUnit = 500
				albumList.TotalLove = 500
			} else if rit == 5 {
				albumList.HighestLovePerUnit = 750
				albumList.TotalLove = 750
			}
		} else {
			albumList.HighestLovePerUnit = 1000
			albumList.TotalLove = 1000

			rs, err := db.Query("SELECT COUNT(*) AS ct FROM unit_sign_asset_m WHERE unit_id = ?", uid)
			if err != nil {
				panic(err)
			}

			ct := 0
			for rs.Next() {
				err = rs.Scan(&ct)
				if err != nil {
					panic(err)
				}
			}
			if ct > 0 {
				albumList.SignFlag = true
			} else {
				albumList.SignFlag = false
			}
		}

		albumLists = append(albumLists, albumList)
	}

	albumResp := model.AlbumResp{
		// _ = model.AlbumResp{
		Result:     albumLists,
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, albumResp)

	// scenario_status_result
	sql = `SELECT scenario_id FROM scenario_m ORDER BY scenario_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	scenarioLists := []model.ScenarioStatusList{}
	for rows.Next() {
		var sid int
		err = rows.Scan(&sid)
		if err != nil {
			panic(err)
		}
		scenarioLists = append(scenarioLists, model.ScenarioStatusList{
			ScenarioID: sid,
			Status:     2,
		})
	}
	scenarioResp := model.ScenarioStatusResp{
		// _ = model.ScenarioStatusResp{
		Result: model.ScenarioStatusResult{
			ScenarioStatusList: scenarioLists,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, scenarioResp)

	// subscenario_status_result
	sql = `SELECT subscenario_id FROM subscenario_m ORDER BY subscenario_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	subScenarioLists := []model.SubscenarioStatusList{}
	for rows.Next() {
		var sid int
		err = rows.Scan(&sid)
		if err != nil {
			panic(err)
		}
		subScenarioLists = append(subScenarioLists, model.SubscenarioStatusList{
			SubscenarioID: sid,
			Status:        2,
		})
	}
	subScenarioResp := model.SubscenarioStatusResp{
		// _ = model.SubscenarioStatusResp{
		Result: model.SubscenarioStatusResult{
			SubscenarioStatusList:  subScenarioLists,
			UnlockedSubscenarioIds: []interface{}{},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, subScenarioResp)

	// event_scenario_result
	eventsList := []model.EventScenarioList{}
	sql = `SELECT event_id FROM event_scenario_m GROUP BY event_id ORDER BY event_id DESC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var eventId int
		err = rows.Scan(&eventId)
		if err != nil {
			panic(err)
		}

		sql = `SELECT event_scenario_id,chapter,chapter_asset,open_date FROM event_scenario_m WHERE event_id = ? ORDER BY chapter DESC`
		chaps, err := db.Query(sql, eventId)
		if err != nil {
			panic(err)
		}
		chapsList := []model.EventScenarioChapterList{}
		var open_date string
		for chaps.Next() {
			var event_scenario_id, chapter int
			var chapter_asset interface{}
			err = chaps.Scan(&event_scenario_id, &chapter, &chapter_asset, &open_date)
			if err != nil {
				panic(err)
			}
			chapList := model.EventScenarioChapterList{
				EventScenarioID: event_scenario_id,
				Chapter:         chapter,
				Status:          2,
				OpenFlashFlag:   0,
				IsReward:        false,
				CostType:        1000,
				ItemID:          1200,
				Amount:          1,
			}
			if chapter_asset != nil {
				chapList.ChapterAsset = chapter_asset.(string)
			}

			chapsList = append(chapsList, chapList)
		}

		eventList := model.EventScenarioList{
			EventID:     eventId,
			OpenDate:    strings.ReplaceAll(open_date, "/", "-"),
			ChapterList: chapsList,
		}

		// HACK event_scenario_btn_asset
		if eventId == 10001 {
			eventList.EventScenarioBtnAsset = "assets/image/ui/eventscenario/38_se_ba_t.png"
		} else if eventId == 221 {
			eventList.EventScenarioBtnAsset = "assets/image/ui/eventscenario/215_se_ba_t.png"
		} else {
			eventList.EventScenarioBtnAsset = fmt.Sprintf("assets/image/ui/eventscenario/%d_se_ba_t.png", eventId)
		}

		eventsList = append(eventsList, eventList)
	}
	eventScenarioResp := model.EventScenarioStatusResp{
		// _ = model.EventScenarioStatusResp{
		Result: model.EventScenarioStatusResult{
			EventScenarioList: eventsList, //
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, eventScenarioResp)

	// multi_unit_scenario_result
	sql = `SELECT multi_unit_id FROM multi_unit_scenario_m GROUP BY multi_unit_id ORDER BY multi_unit_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	var mId int
	multiUnitsList := []model.MultiUnitScenarioStatusList{}
	for rows.Next() {
		err = rows.Scan(&mId)
		if err != nil {
			panic(err)
		}

		sql = `SELECT multi_unit_scenario_btn_asset,open_date,multi_unit_scenario_id,chapter FROM multi_unit_scenario_m a LEFT JOIN multi_unit_scenario_open_m b ON a.multi_unit_id = b.multi_unit_id WHERE a.multi_unit_id = ?`
		units, err := db.Query(sql, mId)
		if err != nil {
			panic(err)
		}
		var multi_unit_scenario_id, chapter int
		var multi_unit_scenario_btn_asset, open_date string
		for units.Next() {
			err = units.Scan(&multi_unit_scenario_btn_asset, &open_date, &multi_unit_scenario_id, &chapter)
			if err != nil {
				panic(err)
			}
		}

		multiUnitsList = append(multiUnitsList, model.MultiUnitScenarioStatusList{
			MultiUnitID:               mId,
			Status:                    2,
			MultiUnitScenarioBtnAsset: multi_unit_scenario_btn_asset,
			OpenDate:                  strings.ReplaceAll(open_date, "/", "-"),
			ChapterList: []model.MultiUnitScenarioChapterList{
				{
					MultiUnitScenarioID: multi_unit_scenario_id,
					Chapter:             chapter,
					Status:              2,
				},
			},
		})
	}
	unitsResp := model.MultiUnitScenarioStatusResp{
		// _ = model.MultiUnitScenarioStatusResp{
		Result: model.MultiUnitScenarioStatusResult{
			MultiUnitScenarioStatusList:  multiUnitsList,
			UnlockedMultiUnitScenarioIds: []interface{}{},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, unitsResp)

	// product_result
	productResp := model.ProductListResp{
		// _ = model.ProductListResp{
		Result: model.ProductListResult{
			RestrictionInfo: model.RestrictionInfo{
				Restricted: false,
			},
			UnderAgeInfo: model.UnderAgeInfo{
				BirthSet:    false,
				HasLimit:    false,
				LimitAmount: nil,
				MonthUsed:   0,
			},
			SnsProductList:   []model.SnsProductList{},
			ProductList:      []model.ProductList{},
			SubscriptionList: []model.SubscriptionList{},
			ShowPointShop:    false,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, productResp)

	// banner_result
	bannerResp := model.BannerListResp{
		// _ = model.BannerListResp{
		Result: model.BannerListResult{
			TimeLimit: "2037-12-31 23:59:59",
			BannerList: []model.BannerList{
				{
					BannerType:       1,
					TargetID:         1743,
					AssetPath:        "assets/image/secretbox/icon/s_ba_1743_1.png",
					FixedFlag:        false,
					BackSide:         false,
					BannerID:         101151,
					StartDate:        "2013-04-15 00:00:00",
					EndDate:          "2037-12-31 23:59:59",
					AddUnitStartDate: "2022-01-01 00:00:00",
				},
				{
					BannerType: 2,
					TargetID:   1,
					AssetPath:  "assets/image/webview/wv_ba_01.png",
					WebviewURL: "/manga",
					FixedFlag:  false,
					BackSide:   true,
					BannerID:   200001,
					StartDate:  "2016-10-15 15:00:00",
					EndDate:    "2037-12-31 23:59:59",
				},
			},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, bannerResp)

	// item_marquee_result
	marqueeResp := model.NoticeMarqueeResp{
		// _ = model.NoticeMarqueeResp{
		Result: model.NoticeMarqueeResult{
			ItemCount:   0,
			MarqueeList: []interface{}{},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, marqueeResp)

	// user_intro_result
	userIntroResp := model.UserNaviResp{
		// _ = model.UserNaviResp{
		Result: model.UserNaviResult{
			User: model.User{
				UserID:           9999999,
				UnitOwningUserID: centerId,
			},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, userIntroResp)

	// special_cutin_result
	cutinResp := model.SpecialCutinResp{
		// _ = model.SpecialCutinResp{
		Result: model.SpecialCutinResult{
			SpecialCutinList: []interface{}{},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, cutinResp)

	// award_result
	sql = `SELECT award_id FROM award_m ORDER BY award_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	awardsList := []model.AwardInfo{}
	for rows.Next() {
		var aId int
		err = rows.Scan(&aId)
		if err != nil {
			panic(err)
		}
		isSet := false
		if aId == 113 { // 极推穗乃果
			isSet = true
		}
		awardsList = append(awardsList, model.AwardInfo{
			AwardID:    aId,
			IsSet:      isSet,
			InsertDate: time.Now().Format("2006-01-02 03:04:05"),
		})
	}
	awardResp := model.AwardInfoResp{
		// _ = model.AwardInfoResp{
		Result: model.AwardInfoResult{
			AwardInfo: awardsList,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, awardResp)

	// background_result
	sql = `SELECT background_id FROM background_m ORDER BY background_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	backgroundsList := []model.BackgroundInfo{}
	for rows.Next() {
		var bId int
		err = rows.Scan(&bId)
		if err != nil {
			panic(err)
		}
		isSet := false
		if bId == 143 { // 穗乃果的房间[情人节]
			isSet = true
		}
		backgroundsList = append(backgroundsList, model.BackgroundInfo{
			BackgroundID: bId,
			IsSet:        isSet,
			InsertDate:   time.Now().Format("2006-01-02 03:04:05"),
		})
	}
	backgroundResp := model.BackgroundInfoResp{
		// _ = model.BackgroundInfoResp{
		Result: model.BackgroundInfoResult{
			BackgroundInfo: backgroundsList,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, backgroundResp)

	// stamp_result 使用不到的功能就不弄了
	stampResp := utils.ReadAllText("assets/stamp.json")
	var mStampResp interface{}
	err = json.Unmarshal([]byte(stampResp), &mStampResp)
	if err != nil {
		panic(err)
	}
	// _ = utils.ReadAllText("assets/stamp.json")
	respAll = append(respAll, mStampResp)

	// exchange_point_result
	sql = `SELECT exchange_point_id FROM exchange_point_m ORDER BY exchange_point_id ASC`
	rows, err = db.Query(sql)
	if err != nil {
		panic(err)
	}
	exPointsList := []model.ExchangePointList{}
	for rows.Next() {
		var eId int
		err = rows.Scan(&eId)
		if err != nil {
			panic(err)
		}
		exPointsList = append(exPointsList, model.ExchangePointList{
			Rarity:        eId,
			ExchangePoint: 9999,
		})
	}
	exPointsResp := model.ExchangePointResp{
		// _ = model.ExchangePointResp{
		Result: model.ExchangePointResult{
			ExchangePointList: exPointsList,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, exPointsResp)

	// live_se_result
	liveSeResp := model.LiveSeInfoResp{
		// _ = model.LiveSeInfoResp{
		Result: model.LiveSeInfoResult{
			LiveSeList: []int{1, 2, 3},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, liveSeResp)

	// live_icon_result
	liveIconResp := model.LiveIconInfoResp{
		// _ = model.LiveIconInfoResp{
		Result: model.LiveIconInfoResult{
			LiveNotesIconList: []int{1, 2, 3},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, liveIconResp)

	// item_list_result 暂时不知道部分字段啥逻辑
	itemResp := utils.ReadAllText("assets/item.json")
	// _ = utils.ReadAllText("assets/item.json")
	var mItemResp interface{}
	err = json.Unmarshal([]byte(itemResp), &mItemResp)
	if err != nil {
		panic(err)
	}
	respAll = append(respAll, mItemResp)

	// marathon_result
	marathonResp := model.MarathonInfoResp{
		// _ = model.MarathonInfoResp{
		Result:     []interface{}{},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, marathonResp)

	// challenge_result
	challengeResp := model.ChallengeInfoResp{
		// _ = model.ChallengeInfoResp{
		Result:     []interface{}{},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, challengeResp)

	// Final
	k, err := json.Marshal(respAll)
	if err != nil {
		panic(err)
	}
	resp := model.Response{
		ResponseData: k,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	k, err = json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(k))

	utils.WriteAllText("assets/api1.json", string(k))
}

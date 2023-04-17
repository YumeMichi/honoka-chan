package tools

import (
	"encoding/json"
	"fmt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DifficultyResult struct {
	Difficulty int `json:"difficulty"`
	ClearCnt   int `json:"clear_cnt"`
}

type DifficultyResp struct {
	Result     []DifficultyResult `json:"result"`
	Status     int                `json:"status"`
	CommandNum bool               `json:"commandNum"`
	TimeStamp  int64              `json:"timeStamp"`
}

type LoveResp struct {
	Result     []interface{} `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

type UserInfo struct {
	UserID               int    `json:"user_id"`
	Name                 string `json:"name"`
	Level                int    `json:"level"`
	CostMax              int    `json:"cost_max"`
	UnitMax              int    `json:"unit_max"`
	EnergyMax            int    `json:"energy_max"`
	FriendMax            int    `json:"friend_max"`
	UnitCnt              int    `json:"unit_cnt"`
	InviteCode           string `json:"invite_code"`
	ElapsedTimeFromLogin string `json:"elapsed_time_from_login"`
	Introduction         string `json:"introduction"`
}

type AccessoryInfo struct {
	AccessoryOwningUserID int  `json:"accessory_owning_user_id"`
	AccessoryID           int  `json:"accessory_id"`
	Exp                   int  `json:"exp"`
	NextExp               int  `json:"next_exp"`
	Level                 int  `json:"level"`
	MaxLevel              int  `json:"max_level"`
	RankUpCount           int  `json:"rank_up_count"`
	FavoriteFlag          bool `json:"favorite_flag"`
}

type Costume struct {
	UnitID    int  `json:"unit_id"`
	IsRankMax bool `json:"is_rank_max"`
	IsSigned  bool `json:"is_signed"`
}

type CenterUnitInfo struct {
	UnitOwningUserID           int           `json:"unit_owning_user_id"`
	UnitID                     int           `json:"unit_id"`
	Exp                        int           `json:"exp"`
	NextExp                    int           `json:"next_exp"`
	Level                      int           `json:"level"`
	LevelLimitID               int           `json:"level_limit_id"`
	MaxLevel                   int           `json:"max_level"`
	Rank                       int           `json:"rank"`
	MaxRank                    int           `json:"max_rank"`
	Love                       int           `json:"love"`
	MaxLove                    int           `json:"max_love"`
	UnitSkillLevel             int           `json:"unit_skill_level"`
	MaxHp                      int           `json:"max_hp"`
	FavoriteFlag               bool          `json:"favorite_flag"`
	DisplayRank                int           `json:"display_rank"`
	UnitSkillExp               int           `json:"unit_skill_exp"`
	UnitRemovableSkillCapacity int           `json:"unit_removable_skill_capacity"`
	Attribute                  int           `json:"attribute"`
	Smile                      int           `json:"smile"`
	Cute                       int           `json:"cute"`
	Cool                       int           `json:"cool"`
	IsLoveMax                  bool          `json:"is_love_max"`
	IsLevelMax                 bool          `json:"is_level_max"`
	IsRankMax                  bool          `json:"is_rank_max"`
	IsSigned                   bool          `json:"is_signed"`
	IsSkillLevelMax            bool          `json:"is_skill_level_max"`
	SettingAwardID             int           `json:"setting_award_id"`
	RemovableSkillIds          []int         `json:"removable_skill_ids"`
	AccessoryInfo              AccessoryInfo `json:"accessory_info"`
	Costume                    Costume       `json:"costume"`
	TotalSmile                 int           `json:"total_smile"`
	TotalCute                  int           `json:"total_cute"`
	TotalCool                  int           `json:"total_cool"`
	TotalHp                    int           `json:"total_hp"`
}

type NaviUnitInfo struct {
	UnitOwningUserID            int    `json:"unit_owning_user_id"`
	UnitID                      int    `json:"unit_id"`
	Exp                         int    `json:"exp"`
	NextExp                     int    `json:"next_exp"`
	Level                       int    `json:"level"`
	MaxLevel                    int    `json:"max_level"`
	LevelLimitID                int    `json:"level_limit_id"`
	Rank                        int    `json:"rank"`
	MaxRank                     int    `json:"max_rank"`
	Love                        int    `json:"love"`
	MaxLove                     int    `json:"max_love"`
	UnitSkillExp                int    `json:"unit_skill_exp"`
	UnitSkillLevel              int    `json:"unit_skill_level"`
	MaxHp                       int    `json:"max_hp"`
	UnitRemovableSkillCapacity  int    `json:"unit_removable_skill_capacity"`
	FavoriteFlag                bool   `json:"favorite_flag"`
	DisplayRank                 int    `json:"display_rank"`
	IsRankMax                   bool   `json:"is_rank_max"`
	IsLoveMax                   bool   `json:"is_love_max"`
	IsLevelMax                  bool   `json:"is_level_max"`
	IsSigned                    bool   `json:"is_signed"`
	IsSkillLevelMax             bool   `json:"is_skill_level_max"`
	IsRemovableSkillCapacityMax bool   `json:"is_removable_skill_capacity_max"`
	InsertDate                  string `json:"insert_date"`
	TotalSmile                  int    `json:"total_smile"`
	TotalCute                   int    `json:"total_cute"`
	TotalCool                   int    `json:"total_cool"`
	TotalHp                     int    `json:"total_hp"`
	RemovableSkillIds           []int  `json:"removable_skill_ids"`
}

type ProfileResult struct {
	UserInfo            UserInfo       `json:"user_info"`
	CenterUnitInfo      CenterUnitInfo `json:"center_unit_info"`
	NaviUnitInfo        NaviUnitInfo   `json:"navi_unit_info"`
	IsAlliance          bool           `json:"is_alliance"`
	FriendStatus        int            `json:"friend_status"`
	SettingAwardID      int            `json:"setting_award_id"`
	SettingBackgroundID int            `json:"setting_background_id"`
}

type ProfileResp struct {
	Result     ProfileResult `json:"result"`
	Status     int           `json:"status"`
	CommandNum bool          `json:"commandNum"`
	TimeStamp  int64         `json:"timeStamp"`
}

var (
	CenterId = 41674
)

func GenApi1Data() {
	// global
	var respAll []interface{}

	// live_status_result
	var liveDifficultyId int
	normalLives := []model.NormalLiveStatusList{}
	sql := `SELECT live_difficulty_id FROM normal_live_m ORDER BY live_difficulty_id ASC`
	rows, err := MainEng.DB().Query(sql)
	CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&liveDifficultyId)
		CheckErr(err)

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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	for rows.Next() {
		err = rows.Scan(&liveDifficultyId)
		CheckErr(err)

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
	unitsData := []model.Active{}
	err = MainEng.Table("common_unit_m").Select("*").Find(&unitsData)
	if err != nil {
		panic(err)
	}

	userUnits := []model.Active{}
	err = UserEng.Table("user_unit_m").Select("*").Find(&userUnits)
	if err != nil {
		panic(err)
	}

	unitsData = append(unitsData, userUnits...)

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
	uId := 1681205441 // Hardcode for now
	userDeck := []UserDeckData{}
	err = UserEng.Table("user_deck_m").Where("user_id = ?", uId).Asc("deck_id").Find(&userDeck)
	CheckErr(err)

	unitDeckInfo := []model.UnitDeckInfo{}
	for _, deck := range userDeck {
		deckUnit := []UnitDeckData{}
		err = UserEng.Table("deck_unit_m").Where("user_deck_id = ?", deck.ID).Asc("position").Find(&deckUnit)
		CheckErr(err)

		oUids := []model.UnitOwningUserIds{}
		for _, unit := range deckUnit {
			oUids = append(oUids, model.UnitOwningUserIds{
				Position:         unit.Position,
				UnitOwningUserID: unit.UnitOwningUserID,
			})
		}

		mainFlag := false
		if deck.MainFlag == 1 {
			mainFlag = true
		}
		unitDeckInfo = append(unitDeckInfo, model.UnitDeckInfo{
			UnitDeckID:        deck.DeckID,
			MainFlag:          mainFlag,
			DeckName:          deck.DeckName,
			UnitOwningUserIds: oUids,
		})
	}
	unitDeckResp := model.UnitDeckInfoResp{
		// _ = model.UnitDeckInfoResp{
		Result:     unitDeckInfo,
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, unitDeckResp)

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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
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
		CheckErr(err)
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

			stmt, err := MainEng.DB().Prepare("SELECT COUNT(*) AS ct FROM unit_sign_asset_m WHERE unit_id = ?")
			CheckErr(err)
			defer stmt.Close()

			var count int
			err = stmt.QueryRow(uid).Scan(&count)
			CheckErr(err)

			if count > 0 {
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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	scenarioLists := []model.ScenarioStatusList{}
	for rows.Next() {
		var sid int
		err = rows.Scan(&sid)
		CheckErr(err)
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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	subScenarioLists := []model.SubscenarioStatusList{}
	for rows.Next() {
		var sid int
		err = rows.Scan(&sid)
		CheckErr(err)
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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		var eventId int
		err = rows.Scan(&eventId)
		CheckErr(err)

		sql = `SELECT event_scenario_id,chapter,chapter_asset,open_date FROM event_scenario_m WHERE event_id = ? ORDER BY chapter DESC`
		chaps, err := MainEng.DB().Query(sql, eventId)
		CheckErr(err)
		defer chaps.Close()
		chapsList := []model.EventScenarioChapterList{}
		var open_date string
		for chaps.Next() {
			var event_scenario_id, chapter int
			var chapter_asset interface{}
			err = chaps.Scan(&event_scenario_id, &chapter, &chapter_asset, &open_date)
			CheckErr(err)
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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	var mId int
	multiUnitsList := []model.MultiUnitScenarioStatusList{}
	for rows.Next() {
		err = rows.Scan(&mId)
		CheckErr(err)

		sql = `SELECT multi_unit_scenario_btn_asset,open_date,multi_unit_scenario_id,chapter FROM multi_unit_scenario_m a LEFT JOIN multi_unit_scenario_open_m b ON a.multi_unit_id = b.multi_unit_id WHERE a.multi_unit_id = ?`
		units, err := MainEng.DB().Query(sql, mId)
		CheckErr(err)
		defer units.Close()
		var multi_unit_scenario_id, chapter int
		var multi_unit_scenario_btn_asset, open_date string
		for units.Next() {
			err = units.Scan(&multi_unit_scenario_btn_asset, &open_date, &multi_unit_scenario_id, &chapter)
			CheckErr(err)
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
				UnitOwningUserID: CenterId,
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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	awardsList := []model.AwardInfo{}
	for rows.Next() {
		var aId int
		err = rows.Scan(&aId)
		CheckErr(err)
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
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	backgroundsList := []model.BackgroundInfo{}
	for rows.Next() {
		var bId int
		err = rows.Scan(&bId)
		CheckErr(err)
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
	CheckErr(err)
	// _ = utils.ReadAllText("assets/stamp.json")
	respAll = append(respAll, mStampResp)

	// exchange_point_result
	sql = `SELECT exchange_point_id FROM exchange_point_m ORDER BY exchange_point_id ASC`
	rows, err = MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	exPointsList := []model.ExchangePointList{}
	for rows.Next() {
		var eId int
		err = rows.Scan(&eId)
		CheckErr(err)
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
	CheckErr(err)
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
	CheckErr(err)
	resp := model.Response{
		ResponseData: k,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	k, err = json.Marshal(resp)
	CheckErr(err)
	// fmt.Println(string(k))

	utils.WriteAllText("assets/api1.json", string(k))
}

func GenApi2Data() {
	// global
	var respAll []interface{}

	// login_topinfo_result
	topInfoResp := model.TopInfoResp{
		// _ = model.TopInfoResp{
		Result: model.TopInfoResult{
			FriendActionCnt:        0,
			FriendGreetCnt:         0,
			FriendVarietyCnt:       0,
			FriendNewCnt:           0,
			PresentCnt:             0,
			SecretBoxBadgeFlag:     false,
			ServerDatetime:         time.Now().Format("2006-01-02 15:04:05"),
			ServerTimestamp:        time.Now().Unix(),
			NoticeFriendDatetime:   time.Now().Format("2006-01-02 15:04:05"),
			NoticeMailDatetime:     "2000-01-01 12:00:00",
			FriendsApprovalWaitCnt: 0,
			FriendsRequestCnt:      0,
			IsTodayBirthday:        false,
			LicenseInfo: model.TopInfoLicenseInfo{
				LicenseList:  []interface{}{},
				LicensedInfo: []interface{}{},
				ExpiredInfo:  []interface{}{},
				BadgeFlag:    false,
			},
			UsingBuffInfo:     []interface{}{},
			IsKlabIDTaskFlag:  false,
			KlabIDTaskCanSync: false,
			HasUnreadAnnounce: false,
			ExchangeBadgeCnt:  []int{0, 0, 0},
			AdFlag:            false,
			HasAdReward:       false,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, topInfoResp)

	// login_topinfo_once_result
	topInfoOnceResp := model.TopInfoOnceResp{
		// _ = model.TopInfoOnceResp{
		Result: model.TopInfoOnceResult{
			NewAchievementCnt:            0,
			UnaccomplishedAchievementCnt: 0,
			LiveDailyRewardExist:         false,
			TrainingEnergy:               10,
			TrainingEnergyMax:            10,
			Notification: model.TopInfoOnceNotification{
				Push:       false,
				Lp:         false,
				UpdateInfo: false,
				Campaign:   false,
				Live:       false,
				Lbonus:     false,
				Event:      false,
				Secretbox:  false,
				Birthday:   true,
			},
			OpenArena:               true,
			CostumeStatus:           true,
			OpenAccessory:           true,
			ArenaSiSkillUniqueCheck: true,
			OpenV98:                 true,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, topInfoOnceResp)

	// unit_accessory_result
	unitAccResp := model.UnitAccessoryAllResp{
		// _ = model.UnitAccessoryAllResp{
		Result: model.UnitAccessoryAllResult{
			AccessoryList:      []model.AccessoryList{},
			WearingInfo:        []model.WearingInfo{},
			EspecialCreateFlag: false,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, unitAccResp)

	// museum_result
	sql := `SELECT museum_contents_id,smile_buff,pure_buff,cool_buff FROM museum_contents_m ORDER BY museum_contents_id ASC`
	rows, err := MainEng.DB().Query(sql)
	CheckErr(err)
	defer rows.Close()
	var smileBuf, pureBuf, coolBuf int
	var mIds []int
	for rows.Next() {
		var museum_contents_id, smile_buff, pure_buff, cool_buff int
		err = rows.Scan(&museum_contents_id, &smile_buff, &pure_buff, &cool_buff)
		CheckErr(err)
		smileBuf += smile_buff
		pureBuf += pure_buff
		coolBuf += cool_buff
		mIds = append(mIds, museum_contents_id)
	}

	museumInfoResp := model.MuseumInfoResp{
		Result: model.MuseumInfoResult{
			MuseumInfo: model.MuseumInfo{
				Parameter: model.MuseumInfoParameter{
					Smile: smileBuf,
					Pure:  pureBuf,
					Cool:  coolBuf,
				},
				ContentsIDList: mIds,
			},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, museumInfoResp)

	// Final
	k, err := json.Marshal(respAll)
	CheckErr(err)
	resp := model.Response{
		ResponseData: k,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	k, err = json.Marshal(resp)
	CheckErr(err)
	// fmt.Println(string(k))

	utils.WriteAllText("assets/api2.json", string(k))
}

func GenApi3Data() {
	// global
	var respAll []interface{}

	// profile_livecnt_result
	difficultyResp := DifficultyResp{
		// _ = DifficultyResp{
		Result: []DifficultyResult{
			{
				Difficulty: 1,
				ClearCnt:   315,
			},
			{
				Difficulty: 2,
				ClearCnt:   310,
			},
			{
				Difficulty: 3,
				ClearCnt:   314,
			},
			{
				Difficulty: 4,
				ClearCnt:   455,
			},
			{
				Difficulty: 6,
				ClearCnt:   233,
			},
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, difficultyResp)

	// profile_card_ranking_result
	var result []interface{}
	love := utils.ReadAllText("assets/love.json")
	err := json.Unmarshal([]byte(love), &result)
	CheckErr(err)
	loveResp := LoveResp{
		// _ = LoveResp{
		Result:     result,
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, loveResp)

	// profile_info_result
	profileResp := ProfileResp{
		Result: ProfileResult{
			UserInfo: UserInfo{
				UserID:               9999999, // 3241988
				Name:                 "\u68a6\u8def @\u65c5\u7acb\u3061\u306e\u65e5\u306b",
				Level:                1028,
				CostMax:              100,
				UnitMax:              5000,
				EnergyMax:            417,
				FriendMax:            99,
				UnitCnt:              3898,
				InviteCode:           "377385143",
				ElapsedTimeFromLogin: "14\u5c0f\u65f6\u524d",
				Introduction:         "\u5728\u4e0d\u77e5\u4e0d\u89c9\u4e2d \u65f6\u5149\u4ece\u4e0d\u505c\u7559\\n\u5982\u4eca\u7684\u6211\u4eec \u6bd5\u4e1a\u518d\u56de\u9996\\n\u8ffd\u68a6\u7684\u670b\u53cb\u4eec\u4e00\u540c \u8e0f\u4e0a\u4e86\u65b0\u7684\u5f81\u9014\\n\u5728\u4e0d\u4e45\u7684\u4eca\u540e \u5728\u672a\u77e5\u7684\u67d0\u5904\\n    \u6211\u4eec\u4e00\u5b9a\u4f1a\u518d\u4e00\u6b21\u9082\u9005\\n    \u8bf7\u4e0d\u8981\u5fd8\u8bb0\u6211\u4eec \u66fe\u7ecf\u7684\u7b11\u5bb9\r",
			},
			CenterUnitInfo: CenterUnitInfo{
				UnitOwningUserID:           CenterId,
				UnitID:                     3927,
				Exp:                        79700,
				NextExp:                    0,
				Level:                      100,
				LevelLimitID:               2,
				MaxLevel:                   100,
				Rank:                       2,
				MaxRank:                    2,
				Love:                       1000,
				MaxLove:                    1000,
				UnitSkillLevel:             8,
				MaxHp:                      6,
				FavoriteFlag:               true,
				DisplayRank:                2,
				UnitSkillExp:               29900,
				UnitRemovableSkillCapacity: 8,
				Attribute:                  1,
				Smile:                      1,
				Cute:                       1,
				Cool:                       1,
				IsLoveMax:                  true,
				IsLevelMax:                 true,
				IsRankMax:                  true,
				IsSigned:                   true,
				IsSkillLevelMax:            true,
				SettingAwardID:             113,
				RemovableSkillIds:          []int{},
				AccessoryInfo:              AccessoryInfo{},
				Costume:                    Costume{},
				TotalSmile:                 1,
				TotalCute:                  1,
				TotalCool:                  1,
				TotalHp:                    6,
			},
			NaviUnitInfo: NaviUnitInfo{
				UnitOwningUserID:            CenterId,
				UnitID:                      3927,
				Exp:                         79700,
				NextExp:                     0,
				Level:                       100,
				MaxLevel:                    100,
				LevelLimitID:                2,
				Rank:                        2,
				MaxRank:                     2,
				Love:                        1000,
				MaxLove:                     1000,
				UnitSkillExp:                29900,
				UnitSkillLevel:              8,
				MaxHp:                       6,
				UnitRemovableSkillCapacity:  8,
				FavoriteFlag:                true,
				DisplayRank:                 2,
				IsLoveMax:                   true,
				IsLevelMax:                  true,
				IsRankMax:                   true,
				IsSigned:                    true,
				IsSkillLevelMax:             true,
				IsRemovableSkillCapacityMax: true,
				InsertDate:                  "2016-10-11 10:33:03",
				TotalSmile:                  1,
				TotalCute:                   1,
				TotalCool:                   1,
				TotalHp:                     6,
				RemovableSkillIds:           []int{},
			},
			IsAlliance:          false,
			FriendStatus:        0,
			SettingAwardID:      113,
			SettingBackgroundID: 143,
		},
		Status:     200,
		CommandNum: false,
		TimeStamp:  time.Now().Unix(),
	}
	respAll = append(respAll, profileResp)

	// Final
	k, err := json.Marshal(respAll)
	CheckErr(err)
	resp := model.Response{
		ResponseData: k,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	k, err = json.Marshal(resp)
	CheckErr(err)
	// fmt.Println(string(k))

	utils.WriteAllText("assets/api3.json", string(k))
}

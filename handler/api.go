package handler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/tools"
	"honoka-chan/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func ApiHandler(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.GetString("userid"))
	CheckErr(err)

	var formdata []model.SifApi
	err = json.Unmarshal([]byte(ctx.PostForm("request_data")), &formdata)
	if err != nil {
		fmt.Println(err)
		return
	}
	var results []interface{}
	for _, v := range formdata {
		var res []byte
		var err error
		// fmt.Println(v)
		fmt.Println(v.Module, v.Action)

		switch v.Module {
		case "login":
			if v.Action == "topInfo" {
				// key = "login_topinfo_result"
				topInfoResp := model.TopInfoResp{
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
				res, err = json.Marshal(topInfoResp)
				CheckErr(err)
			} else if v.Action == "topInfoOnce" {
				// key = "login_topinfo_once_result"
				topInfoOnceResp := model.TopInfoOnceResp{
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
				res, err = json.Marshal(topInfoOnceResp)
				CheckErr(err)
			}
		case "live":
			if v.Action == "liveStatus" {
				// key = "live_status_result"
				var liveDifficultyId int
				normalLives := []model.NormalLiveStatusList{}
				sql := `SELECT live_difficulty_id FROM normal_live_m ORDER BY live_difficulty_id ASC`
				rows, err := MainEng.DB().Query(sql)
				CheckErr(err)
				defer rows.Close()
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
				defer rows.Close()
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
				res, err = json.Marshal(LiveStatusResp)
				CheckErr(err)
			} else if v.Action == "schedule" {
				// key = "live_list_result"
				var liveDifficultyId int
				specialLives := []model.SpecialLiveStatusList{}
				sql := `SELECT live_difficulty_id FROM special_live_m ORDER BY live_difficulty_id ASC`
				rows, err := MainEng.DB().Query(sql)
				CheckErr(err)
				defer rows.Close()
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
				res, err = json.Marshal(liveListResp)
				CheckErr(err)
			}
		case "unit":
			switch v.Action {
			case "unitAll":
				// key = "unit_list_result"
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
					Result: model.UnitAllResult{
						Active:  unitsData,
						Waiting: []model.Waiting{},
					},
					Status:     200,
					CommandNum: false,
					TimeStamp:  time.Now().Unix(),
				}
				res, err = json.Marshal(unitListResp)
				CheckErr(err)
			case "deckInfo":
				// key = "unit_deck_result"
				userDeck := []tools.UserDeckData{}
				err = UserEng.Table("user_deck_m").Where("user_id = ?", userId).Asc("deck_id").Find(&userDeck)
				CheckErr(err)

				unitDeckInfo := []model.UnitDeckInfo{}
				for _, deck := range userDeck {
					deckUnit := []tools.UnitDeckData{}
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
					Result:     unitDeckInfo,
					Status:     200,
					CommandNum: false,
					TimeStamp:  time.Now().Unix(),
				}
				res, err = json.Marshal(unitDeckResp)
				CheckErr(err)
			case "supporterAll":
				// key = "unit_support_result"
				unitSupportResp := model.UnitSupportResp{
					Result: model.UnitSupportResult{
						UnitSupportList: []model.UnitSupportList{},
					}, // 练习道具
					Status:     200,
					CommandNum: false,
					TimeStamp:  time.Now().Unix(),
				}
				res, err = json.Marshal(unitSupportResp)
				CheckErr(err)
			case "removableSkillInfo":
				// key = "owning_equip_result"
				rmSkillResp := model.RemovableSkillResp{
					Result: model.RemovableSkillResult{
						OwningInfo:    []model.OwningInfo{},
						EquipmentInfo: []interface{}{},
					}, // 宝石
					Status:     200,
					CommandNum: false,
					TimeStamp:  time.Now().Unix(),
				}
				res, err = json.Marshal(rmSkillResp)
				CheckErr(err)
			case "accessoryAll":
				// key = "unit_accessory_result"
				accessoryList := []model.AccessoryList{}
				err := MainEng.Table("common_accessory_m").Find(&accessoryList)
				CheckErr(err)
				for k, _ := range accessoryList {
					accessoryList[k].NextExp = 0
					accessoryList[k].Level = 8
					accessoryList[k].MaxLevel = 8
					accessoryList[k].RankUpCount = 4
					accessoryList[k].FavoriteFlag = true
				}
				wearingInfo := []model.WearingInfo{}
				err = UserEng.Table("accessory_wear_m").Where("user_id = ?", ctx.GetString("userid")).Find(&wearingInfo)
				CheckErr(err)
				unitAccResp := model.UnitAccessoryAllResp{
					Result: model.UnitAccessoryAllResult{
						AccessoryList:      accessoryList,
						WearingInfo:        wearingInfo,
						EspecialCreateFlag: false,
					},
					Status:     200,
					CommandNum: false,
					TimeStamp:  time.Now().Unix(),
				}
				res, err = json.Marshal(unitAccResp)
				CheckErr(err)
			// case "accessoryTab":
			// case "accessoryMaterialAll":
			default:
				err = errors.New("invalid option")
				CheckErr(err)
			}
		case "costume":
			// key = "costume_list_result"
			costumeListResp := model.CostumeListResp{
				Result: model.CostumeListResult{
					CostumeList: []model.CostumeList{},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(costumeListResp)
			CheckErr(err)
		case "album":
			// key = "album_unit_result"
			albumLists := []model.AlbumResult{}
			sql := `SELECT unit_id,rarity FROM unit_m ORDER BY unit_id ASC`
			rows, err := MainEng.DB().Query(sql)
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
				Result:     albumLists,
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(albumResp)
			CheckErr(err)
		case "scenario":
			// key = "scenario_status_result"
			sql := `SELECT scenario_id FROM scenario_m ORDER BY scenario_id ASC`
			rows, err := MainEng.DB().Query(sql)
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
				Result: model.ScenarioStatusResult{
					ScenarioStatusList: scenarioLists,
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(scenarioResp)
			CheckErr(err)
		case "subscenario":
			// key = "subscenario_status_result"
			sql := `SELECT subscenario_id FROM subscenario_m ORDER BY subscenario_id ASC`
			rows, err := MainEng.DB().Query(sql)
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
				Result: model.SubscenarioStatusResult{
					SubscenarioStatusList:  subScenarioLists,
					UnlockedSubscenarioIds: []interface{}{},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(subScenarioResp)
			CheckErr(err)
		case "eventscenario":
			// key = "event_scenario_result"
			eventsList := []model.EventScenarioList{}
			sql := `SELECT event_id FROM event_scenario_m GROUP BY event_id ORDER BY event_id DESC`
			rows, err := MainEng.DB().Query(sql)
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
				Result: model.EventScenarioStatusResult{
					EventScenarioList: eventsList, //
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(eventScenarioResp)
			CheckErr(err)
		case "multiunit":
			// key = "multi_unit_scenario_result"
			sql := `SELECT multi_unit_id FROM multi_unit_scenario_m GROUP BY multi_unit_id ORDER BY multi_unit_id ASC`
			rows, err := MainEng.DB().Query(sql)
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
				Result: model.MultiUnitScenarioStatusResult{
					MultiUnitScenarioStatusList:  multiUnitsList,
					UnlockedMultiUnitScenarioIds: []interface{}{},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(unitsResp)
			CheckErr(err)
		case "payment":
			// key = "product_result"
			productResp := model.ProductListResp{
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
			res, err = json.Marshal(productResp)
			CheckErr(err)
		case "banner":
			// key = "banner_result"
			bannerResp := model.BannerListResp{
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
			res, err = json.Marshal(bannerResp)
			CheckErr(err)
		case "notice":
			// key = "item_marquee_result"
			marqueeResp := model.NoticeMarqueeResp{
				Result: model.NoticeMarqueeResult{
					ItemCount:   0,
					MarqueeList: []interface{}{},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(marqueeResp)
			CheckErr(err)
		case "user":
			// key = "user_intro_result"
			userIntroResp := model.UserNaviResp{
				Result: model.UserNaviResult{
					User: model.User{
						UserID:           9999999,
						UnitOwningUserID: 41674,
					},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(userIntroResp)
			CheckErr(err)
		case "navigation":
			// key = "special_cutin_result"
			cutinResp := model.SpecialCutinResp{
				Result: model.SpecialCutinResult{
					SpecialCutinList: []interface{}{},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(cutinResp)
			CheckErr(err)
		case "award":
			// key = "award_result"
			sql := `SELECT award_id FROM award_m ORDER BY award_id ASC`
			rows, err := MainEng.DB().Query(sql)
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
				Result: model.AwardInfoResult{
					AwardInfo: awardsList,
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(awardResp)
			CheckErr(err)
		case "background":
			// key = "background_result"
			sql := `SELECT background_id FROM background_m ORDER BY background_id ASC`
			rows, err := MainEng.DB().Query(sql)
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
				Result: model.BackgroundInfoResult{
					BackgroundInfo: backgroundsList,
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(backgroundResp)
			CheckErr(err)
		case "stamp":
			// key = "stamp_result"
			stampResp := utils.ReadAllText("assets/stamp.json")
			var mStampResp interface{}
			err = json.Unmarshal([]byte(stampResp), &mStampResp)
			CheckErr(err)
			res, err = json.Marshal(mStampResp)
			CheckErr(err)
		case "exchange":
			// key = "exchange_point_result"
			sql := `SELECT exchange_point_id FROM exchange_point_m ORDER BY exchange_point_id ASC`
			rows, err := MainEng.DB().Query(sql)
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
				Result: model.ExchangePointResult{
					ExchangePointList: exPointsList,
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(exPointsResp)
			CheckErr(err)
		case "livese":
			// key = "live_se_result"
			liveSeResp := model.LiveSeInfoResp{
				Result: model.LiveSeInfoResult{
					LiveSeList: []int{1, 2, 3},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(liveSeResp)
			CheckErr(err)
		case "liveicon":
			// key = "live_icon_result"
			liveIconResp := model.LiveIconInfoResp{
				Result: model.LiveIconInfoResult{
					LiveNotesIconList: []int{1, 2, 3},
				},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(liveIconResp)
			CheckErr(err)
		case "item":
			// key = "item_list_result"
			itemResp := utils.ReadAllText("assets/item.json")
			var mItemResp interface{}
			err = json.Unmarshal([]byte(itemResp), &mItemResp)
			CheckErr(err)
			res, err = json.Marshal(mItemResp)
			CheckErr(err)
		case "marathon":
			// key = "marathon_result"
			marathonResp := model.MarathonInfoResp{
				Result:     []interface{}{},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(marathonResp)
			CheckErr(err)
		case "challenge":
			// key = "challenge_result"
			challengeResp := model.ChallengeInfoResp{
				Result:     []interface{}{},
				Status:     200,
				CommandNum: false,
				TimeStamp:  time.Now().Unix(),
			}
			res, err = json.Marshal(challengeResp)
			CheckErr(err)
		case "museum":
			// key = "museum_result"
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
			res, err = json.Marshal(museumInfoResp)
			CheckErr(err)
		case "profile":
			if v.Action == "liveCnt" {
				// key = "profile_livecnt_result"
				difficultyResp := tools.DifficultyResp{
					Result: []tools.DifficultyResult{
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
				res, err = json.Marshal(difficultyResp)
				CheckErr(err)
			} else if v.Action == "cardRanking" {
				// key = "profile_card_ranking_result"
				var result []interface{}
				love := utils.ReadAllText("assets/love.json")
				err := json.Unmarshal([]byte(love), &result)
				CheckErr(err)
				loveResp := tools.LoveResp{
					Result:     result,
					Status:     200,
					CommandNum: false,
					TimeStamp:  time.Now().Unix(),
				}
				res, err = json.Marshal(loveResp)
				CheckErr(err)
			} else if v.Action == "profileInfo" {
				// key = "profile_info_result"
				pref := tools.UserPref{}
				exists, err := UserEng.Table("user_preference_m").Where("user_id = ?", ctx.GetString("userid")).Get(&pref)
				CheckErr(err)
				if !exists {
					ctx.String(http.StatusForbidden, ErrorMsg)
					return
				}
				profileResp := tools.ProfileResp{
					Result: tools.ProfileResult{
						UserInfo: tools.UserInfo{
							UserID:               pref.UserID,
							Name:                 pref.UserName,
							Level:                pref.UserLevel,
							CostMax:              100,
							UnitMax:              5000,
							EnergyMax:            1000,
							FriendMax:            99,
							UnitCnt:              3898,
							InviteCode:           strconv.Itoa(pref.UserID),
							ElapsedTimeFromLogin: "14\u5c0f\u65f6\u524d",
							Introduction:         pref.UserDesc,
						},
						CenterUnitInfo: tools.CenterUnitInfo{
							UnitOwningUserID:           41674,
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
							AccessoryInfo:              tools.AccessoryInfo{},
							Costume:                    tools.Costume{},
							TotalSmile:                 1,
							TotalCute:                  1,
							TotalCool:                  1,
							TotalHp:                    6,
						},
						NaviUnitInfo: tools.NaviUnitInfo{
							UnitOwningUserID:            41674,
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
				res, err = json.Marshal(profileResp)
				CheckErr(err)
			}
		default:
			// fmt.Println(ErrorMsg)
			fmt.Println(v)
			err = errors.New("invalid option")
			CheckErr(err)
		}

		var result interface{}
		err = json.Unmarshal([]byte(res), &result)
		CheckErr(err)
		results = append(results, result)
	}
	// fmt.Println(results)
	b, err := json.Marshal(results)
	CheckErr(err)
	rp := model.Response{
		ResponseData: b,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	b, err = json.Marshal(rp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(b, "privatekey.pem")))

	ctx.String(http.StatusOK, string(b))
}

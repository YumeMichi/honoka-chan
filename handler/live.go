package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
)

type GameOverResp struct {
	ResponseData []interface{} `json:"response_data"`
	ReleaseInfo  []interface{} `json:"release_info"`
	StatusCode   int           `json:"status_code"`
}

func PartyListHandler(ctx *gin.Context) {
	resp := utils.ReadAllText("assets/partylist.json")

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1([]byte(resp), "privatekey.pem")))

	ctx.String(http.StatusOK, resp)
}

func PlayLiveHandler(ctx *gin.Context) {
	playReq := model.PlayReq{}
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &playReq)
	CheckErr(err)

	tDifficultyId := playReq.LiveDifficultyID
	difficultyId, err := strconv.Atoi(tDifficultyId)
	CheckErr(err)
	deckId := playReq.UnitDeckID

	// Save Deck Id for /live/reward
	key := "live_deck_" + ctx.GetString("userid")
	err = database.LevelDb.Put([]byte(key), []byte(strconv.Itoa(deckId)))
	CheckErr(err)

	// Song type: normal / special
	// sqlite3 doesn't support FULL OUTER JOIN so use UNION ALL here.
	sql := `SELECT notes_setting_asset,c_rank_score,b_rank_score,a_rank_score,s_rank_score,ac_flag,swing_flag FROM live_setting_m WHERE live_setting_id IN (SELECT live_setting_id FROM normal_live_m WHERE live_difficulty_id = ? UNION ALL SELECT live_setting_id FROM special_live_m WHERE live_difficulty_id = ?)`
	var notes_setting_asset string
	var c_rank_score, b_rank_score, a_rank_score, s_rank_score, ac_flag, swing_flag int
	err = MainEng.DB().QueryRow(sql, difficultyId, difficultyId).Scan(&notes_setting_asset, &c_rank_score, &b_rank_score, &a_rank_score, &s_rank_score, &ac_flag, &swing_flag)
	CheckErr(err)

	// fmt.Println(notes_setting_asset)
	// fmt.Println(c_rank_score, b_rank_score, a_rank_score, s_rank_score)

	notes := []model.NotesList{}
	// fmt.Println("./assets/notes/" + notes_setting_asset)
	notes_list := utils.ReadAllText("./assets/notes/" + notes_setting_asset)
	err = json.Unmarshal([]byte(notes_list), &notes)
	CheckErr(err)

	ranks := []model.RankInfo{}
	ranks = append(ranks, model.RankInfo{
		Rank:    5,
		RankMin: 0,
		RankMax: c_rank_score,
	}, model.RankInfo{
		Rank:    4,
		RankMin: c_rank_score + 1,
		RankMax: b_rank_score,
	}, model.RankInfo{
		Rank:    3,
		RankMin: b_rank_score + 1,
		RankMax: a_rank_score,
	}, model.RankInfo{
		Rank:    2,
		RankMin: a_rank_score + 1,
		RankMax: s_rank_score,
	}, model.RankInfo{
		Rank:    1,
		RankMin: s_rank_score + 1,
		RankMax: 0,
	})

	// UserEng.ShowSQL(true)
	// MainEng.ShowSQL(true)
	owningIdList := []int{}
	err = UserEng.Table("deck_unit_m").Join("LEFT", "user_deck_m", "deck_unit_m.user_deck_id = user_deck_m.id").
		Where("user_id = ? AND deck_id = ?", ctx.GetString("userid"), deckId).Cols("unit_owning_user_id").
		OrderBy("deck_unit_m.position ASC").Find(&owningIdList)
	CheckErr(err)

	unitList := []model.UnitList{}
	var totalSmile, totalPure, totalCool, maxLove float64
	var totalHp int
	for _, owningId := range owningIdList {
		var uId int
		exists, err := MainEng.Table("common_unit_m").Where("unit_owning_user_id = ?", owningId).Cols("unit_id").Get(&uId)
		CheckErr(err)

		var maxHp, attrId, unitTypeId int
		var baseSmile, basePure, baseCool, smileMax, pureMax, coolMax float64
		if exists {
			// 公共卡片仅为100级属性
			_, err = MainEng.Table("unit_m").Where("unit_id = ?", uId).
				Join("LEFT", "unit_rarity_m", "unit_m.rarity = unit_rarity_m.rarity").
				Cols("attribute_id,hp_max,smile_max,pure_max,cool_max,after_love_max,unit_type_id").
				Get(&attrId, &maxHp, &baseSmile, &basePure, &baseCool, &maxLove, &unitTypeId)
			CheckErr(err)
		} else {
			// 用户卡片暂时固定为满级350级
			exists, err := UserEng.Table("user_unit_m").Where("unit_owning_user_id = ?", owningId).Cols("unit_id").Get(&uId)
			CheckErr(err)

			if exists {
				// 卡片100级基础属性
				_, err = MainEng.Table("unit_m").Where("unit_id = ?", uId).
					Join("LEFT", "unit_rarity_m", "unit_m.rarity = unit_rarity_m.rarity").
					Cols("attribute_id,hp_max,smile_max,pure_max,cool_max,after_love_max,unit_type_id").
					Get(&attrId, &maxHp, &baseSmile, &basePure, &baseCool, &maxLove, &unitTypeId)
				CheckErr(err)

				// 增量属性
				var diffSmile, diffPure, diffCool float64
				_, err = MainEng.Table("unit_level_limit_pattern_m").Where("unit_level_limit_id = 1 AND unit_level = 350").
					Cols("smile_diff,pure_diff,cool_diff").Get(&diffSmile, &diffPure, &diffCool)
				CheckErr(err)

				// 更新卡片属性（注意这里是负数，要用减号）
				baseSmile -= diffSmile
				basePure -= diffPure
				baseCool -= diffCool
			} else {
				panic("no such unit")
			}
		}

		// 饰品属性加成（满级）
		var accessoryOwningId int
		_, err = UserEng.Table("accessory_wear_m").Where("unit_owning_user_id = ?", owningId).
			Cols("accessory_owning_user_id").Get(&accessoryOwningId)
		CheckErr(err)
		var smileAccessory, pureAccessory, coolAccessory float64
		_, err = MainEng.Table("common_accessory_m").Join("LEFT", "accessory_m", "common_accessory_m.accessory_id = accessory_m.accessory_id").
			Where("accessory_owning_user_id = ?", accessoryOwningId).Cols("smile_max,pure_max,cool_max").
			Get(&smileAccessory, &pureAccessory, &coolAccessory)
		CheckErr(err)
		// fmt.Println("基础属性:", baseSmile, basePure, baseCool)

		// 饰品属性加成（该加成会影响个宝等百分比宝石属性加成的计算，故先计算。）
		baseSmile += smileAccessory
		basePure += pureAccessory
		baseCool += coolAccessory
		// fmt.Println("饰品属性加成:", smileAccessory, pureAccessory, coolAccessory)
		// fmt.Println("饰品属性加成后的基础属性:", baseSmile, basePure, baseCool)

		// 回忆画廊属性加成（该加成会影响个宝等百分比宝石属性加成的计算，故先计算。）
		var smileBuff, pureBuff, coolBuff float64
		_, err = MainEng.Table("museum_contents_m").Select("SUM(smile_buff),SUM(pure_buff),SUM(cool_buff)").Get(&smileBuff, &pureBuff, &coolBuff)
		CheckErr(err)
		baseSmile += smileBuff
		basePure += pureBuff
		baseCool += coolBuff
		// fmt.Println("回忆画廊属性加成:", smileBuff, pureBuff, coolBuff)
		// fmt.Println("回忆画廊属性加成后的基础属性:", baseSmile, basePure, baseCool)

		// 绊属性加成（该加成会影响个宝等百分比宝石属性加成的计算，故先计算。）
		if attrId == 1 {
			baseSmile += maxLove
		} else if attrId == 2 {
			basePure += maxLove
		} else if attrId == 3 {
			baseCool += maxLove
		}
		// fmt.Println("绊属性加成:", maxLove)
		// fmt.Println("绊属性加成后的基础属性:", baseSmile, basePure, baseCool)

		// 宝石属性加成
		var kissSmile, kissPure, kissCool float64
		var skillSmile, skillPure, skillCool float64
		// var mainCenterSmile, mainCenterPure, mainCenterCool float64
		// var secCenterSmile, secCenterPure, secCenterCool float64

		// 宝石加成（满级）
		removableSkillIds := []int{}
		err = UserEng.Table("skill_equip_m").Where("unit_owning_user_id = ? AND user_id = ?", owningId, ctx.GetString("userid")).
			Cols("unit_removable_skill_id").Find(&removableSkillIds)
		CheckErr(err)

		for _, sk := range removableSkillIds {
			// 判断宝石效果类型（效果范围、效果类型、效果值、是否固定数值）
			var effectRange, effectType, fixedValueFlag, refType int
			var effectValue float64
			_, err = MainEng.Table("unit_removable_skill_m").Where("unit_removable_skill_id = ?", sk).
				Cols("effect_range,effect_type,effect_value,fixed_value_flag,target_reference_type").
				Get(&effectRange, &effectType, &effectValue, &fixedValueFlag, &refType)
			CheckErr(err)

			if fixedValueFlag == 1 {
				// 吻、眼神属性加成（固定数值）
				if effectType == 1 {
					kissSmile += effectValue
				} else if effectType == 2 {
					kissPure += effectValue
				} else if effectType == 3 {
					kissCool += effectValue
				}
				// fmt.Println("吻、眼神属性加成:", kissSmile, kissPure, kissCool)
			} else {
				// 仅效果类型为1、2、3的有属性加成
				if effectType == 1 || effectType == 2 || effectType == 3 {
					// 加成范围：2:全员 1:非全员
					if effectRange == 2 {
						if effectType == 1 {
							skillSmile += math.Ceil(baseSmile * (effectValue / 100))
						} else if effectType == 2 {
							skillPure += math.Ceil(basePure * (effectValue / 100))
						} else if effectType == 3 {
							skillCool += math.Ceil(baseCool * (effectValue / 100))
						}
						// fmt.Println("全员类宝石属性加成:", skillSmile, skillPure, skillCool)
					} else {
						// refType: 1 -> 年级类加成, target_type -> 指定年级（这里不需要使用，因为能装上宝石肯定是符合的）
						// refType: 2 -> 个宝
						// refType: 3 -> 爆分、奶、判宝石, 0 -> 竞技场宝石
						if refType == 1 || refType == 2 { // 年级类和个宝都是百分比加成
							if effectType == 1 {
								skillSmile += math.Ceil(baseSmile * (effectValue / 100))
							} else if effectType == 2 {
								skillPure += math.Ceil(basePure * (effectValue / 100))
							} else if effectValue == 3 {
								skillCool += math.Ceil(baseCool * (effectValue / 100))
							}
							// fmt.Println("年级类宝石、个宝属性加成:", skillSmile, skillPure, skillCool)
						}
					}
				}
			}
		}

		// 单卡属性
		smileMax = baseSmile + kissSmile + skillSmile
		pureMax = basePure + kissPure + skillPure
		coolMax = baseCool + kissCool + skillCool

		// 主唱技能加成
		var myCenterUnitId int
		_, err = UserEng.Table("deck_unit_m").Join("LEFT", "user_deck_m", "deck_unit_m.user_deck_id = user_deck_m.id").
			Where("user_deck_m.deck_id = ? AND user_deck_m.user_id = ? AND deck_unit_m.position = 5", playReq.UnitDeckID, ctx.GetString("userid")).
			Cols("deck_unit_m.unit_id").Get(&myCenterUnitId)
		CheckErr(err)

		// 主唱技能加成：主C技能（这里不使用新C技能，即以某属性的百分比提升另一属性）
		var myAttrId int
		var myEffectValue float64
		_, err = MainEng.Table("unit_m").
			Join("LEFT", "unit_leader_skill_m", "unit_m.default_leader_skill_id = unit_leader_skill_m.unit_leader_skill_id").
			Where("unit_m.unit_id = ?", myCenterUnitId).
			Cols("unit_m.attribute_id,unit_leader_skill_m.effect_value").Get(&myAttrId, &myEffectValue)
		CheckErr(err)
		var myCenterSmile, myCenterPure, myCenterCool float64
		if myAttrId == 1 {
			myCenterSmile = math.Ceil(smileMax * (myEffectValue / 100))
		} else if myAttrId == 2 {
			myCenterPure = math.Ceil(pureMax * (myEffectValue / 100))
		} else if myAttrId == 3 {
			myCenterCool = math.Ceil(coolMax * (myEffectValue / 100))
		}
		// fmt.Println("主C技能属性加成:", myCenterSmile, myCenterPure, myCenterCool)

		// 主唱技能加成：副C技能
		var mySubEffectValue float64
		var myMemberTagId int
		_, err = MainEng.Table("unit_m").
			Join("LEFT", "unit_leader_skill_extra_m", "unit_m.default_leader_skill_id = unit_leader_skill_extra_m.unit_leader_skill_id").
			Where("unit_m.unit_id = ?", myCenterUnitId).
			Cols("unit_leader_skill_extra_m.effect_value,unit_leader_skill_extra_m.member_tag_id").Get(&mySubEffectValue, &myMemberTagId)
		CheckErr(err)

		exists, err = MainEng.Table("unit_type_member_tag_m").
			Where("unit_type_id = ? AND member_tag_id = ?", unitTypeId, myMemberTagId).Exist()
		CheckErr(err)
		var mySubSmile, mySubPure, mySubCool float64
		if exists {
			if myAttrId == 1 {
				mySubSmile = math.Ceil(smileMax * (mySubEffectValue / 100))
			} else if myAttrId == 2 {
				mySubPure = math.Ceil(pureMax * (mySubEffectValue / 100))
			} else if myAttrId == 3 {
				mySubCool = math.Ceil(coolMax * (mySubEffectValue / 100))
			}
			// fmt.Println("副C技能属性加成:", mySubSmile, mySubPure, mySubCool)
		}

		// 好友主唱技能加成
		// TODO 好友支援存入数据库
		var tomoUnitId int64
		partyList := gjson.Parse(utils.ReadAllText("assets/partylist.json")).Get("response_data.party_list")
		partyList.ForEach(func(key, value gjson.Result) bool {
			if value.Get("user_info.user_id").Int() == playReq.PartyUserID {
				tomoUnitId = value.Get("center_unit_info.unit_id").Int()
				return false
			}
			return true
		})
		// fmt.Println("好友UnitID:", tomoUnitId)

		// 好友主唱技能加成：主C技能（这里不使用新C技能，即以某属性的百分比提升另一属性）
		var tomoAttrId int
		var tomoEffectValue float64
		_, err = MainEng.Table("unit_m").
			Join("LEFT", "unit_leader_skill_m", "unit_m.default_leader_skill_id = unit_leader_skill_m.unit_leader_skill_id").
			Where("unit_m.unit_id = ?", tomoUnitId).
			Cols("unit_m.attribute_id,unit_leader_skill_m.effect_value").Get(&tomoAttrId, &tomoEffectValue)
		CheckErr(err)
		var tomoCenterSmile, tomoCenterPure, tomoCenterCool float64
		if myAttrId == 1 {
			tomoCenterSmile = math.Ceil(smileMax * (tomoEffectValue / 100))
		} else if myAttrId == 2 {
			tomoCenterPure = math.Ceil(pureMax * (tomoEffectValue / 100))
		} else if myAttrId == 3 {
			tomoCenterCool = math.Ceil(coolMax * (tomoEffectValue / 100))
		}
		// fmt.Println("好友主C技能属性加成:", tomoCenterSmile, tomoCenterPure, tomoCenterCool)

		// 好友主唱技能加成：副C技能
		var tomoSubEffectValue float64
		var tomoMemberTagId int
		_, err = MainEng.Table("unit_m").
			Join("LEFT", "unit_leader_skill_extra_m", "unit_m.default_leader_skill_id = unit_leader_skill_extra_m.unit_leader_skill_id").
			Where("unit_m.unit_id = ?", tomoUnitId).
			Cols("unit_leader_skill_extra_m.effect_value,unit_leader_skill_extra_m.member_tag_id").Get(&tomoSubEffectValue, &tomoMemberTagId)
		CheckErr(err)

		exists, err = MainEng.Table("unit_type_member_tag_m").
			Where("unit_type_id = ? AND member_tag_id = ?", unitTypeId, tomoMemberTagId).Exist()
		CheckErr(err)
		var tomoSubSmile, tomoSubPure, tomoSubCool float64
		if exists {
			if myAttrId == 1 {
				tomoSubSmile = math.Ceil(smileMax * (tomoSubEffectValue / 100))
			} else if myAttrId == 2 {
				tomoSubPure = math.Ceil(pureMax * (tomoSubEffectValue / 100))
			} else if myAttrId == 3 {
				tomoSubCool = math.Ceil(coolMax * (tomoSubEffectValue / 100))
			}
			// fmt.Println("好友副C技能属性加成:", tomoSubSmile, tomoSubPure, tomoSubCool)
		}

		// 全部卡属性
		totalSmile += smileMax + myCenterSmile + mySubSmile + tomoCenterSmile + tomoSubSmile
		totalPure += pureMax + myCenterPure + mySubPure + tomoCenterPure + tomoSubPure
		totalCool += coolMax + myCenterCool + mySubCool + tomoCenterCool + tomoSubCool
		totalHp += maxHp

		// 单卡属性计算结果取上取整
		fixedSmileMax := int(smileMax)
		fixedPureMax := int(pureMax)
		fixedCoolMax := int(coolMax)
		// fmt.Println("单卡属性:", fixedSmileMax, fixedPureMax, fixedCoolMax)

		unitList = append(unitList, model.UnitList{
			Smile: fixedSmileMax,
			Cute:  fixedPureMax,
			Cool:  fixedCoolMax,
		})
	}

	// 全部卡属性计算结果取上取整
	fixedTotalSmile := int(math.Ceil(totalSmile))
	fixedTotalPure := int(math.Ceil(totalPure))
	fixedTotalCool := int(math.Ceil(totalCool))
	// fmt.Println("全卡组属性:", fixedTotalSmile, fixedTotalPure, fixedTotalCool)

	lives := []model.PlayLiveList{}
	lives = append(lives, model.PlayLiveList{
		LiveInfo: model.LiveInfo{
			LiveDifficultyID: difficultyId,
			IsRandom:         false,
			AcFlag:           ac_flag,
			SwingFlag:        swing_flag,
			NotesList:        notes,
		},
		DeckInfo: model.DeckInfo{
			UnitDeckID:       deckId,
			TotalSmile:       fixedTotalSmile,
			TotalCute:        fixedTotalPure,
			TotalCool:        fixedTotalCool,
			TotalHp:          totalHp,
			PreparedHpDamage: 0,
			UnitList:         unitList,
		},
	})

	resp := model.PlayResponseData{
		RankInfo:            ranks,
		EnergyFullTime:      "2023-03-20 01:28:55",
		OverMaxEnergy:       0,
		AvailableLiveResume: false,
		LiveList:            lives,
		IsMarathonEvent:     false,
		MarathonEventID:     nil,
		NoSkill:             false,
		CanActivateEffect:   true,
		ServerTimestamp:     time.Now().Unix(),
	}

	m, err := json.Marshal(resp)
	CheckErr(err)

	res := model.Response{
		ResponseData: m,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}

	mm, err := json.Marshal(res)
	CheckErr(err)
	// fmt.Println(string(mm))

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(mm, "privatekey.pem")))

	ctx.String(http.StatusOK, string(mm))
}

func GameOverHandler(ctx *gin.Context) {
	overResp := GameOverResp{
		ResponseData: []interface{}{},
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(overResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func PlayScoreHandler(ctx *gin.Context) {
	playScoreReq := model.PlayScoreReq{}
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &playScoreReq)
	CheckErr(err)

	tDifficultyId := playScoreReq.LiveDifficultyID
	difficultyId, err := strconv.Atoi(tDifficultyId)
	CheckErr(err)

	// Song type: normal / special
	// sqlite3 doesn't support FULL OUTER JOIN so use UNION ALL here.
	sql := `SELECT notes_setting_asset,c_rank_score,b_rank_score,a_rank_score,s_rank_score,ac_flag,swing_flag FROM live_setting_m WHERE live_setting_id IN (SELECT live_setting_id FROM normal_live_m WHERE live_difficulty_id = ? UNION ALL SELECT live_setting_id FROM special_live_m WHERE live_difficulty_id = ?)`
	var notes_setting_asset string
	var c_rank_score, b_rank_score, a_rank_score, s_rank_score, ac_flag, swing_flag int
	err = MainEng.DB().QueryRow(sql, difficultyId, difficultyId).Scan(&notes_setting_asset, &c_rank_score, &b_rank_score, &a_rank_score, &s_rank_score, &ac_flag, &swing_flag)
	CheckErr(err)

	// fmt.Println(notes_setting_asset)
	// fmt.Println(c_rank_score, b_rank_score, a_rank_score, s_rank_score)

	notes := []model.NotesList{}
	// fmt.Println("./assets/notes/" + notes_setting_asset)
	notes_list := utils.ReadAllText("./assets/notes/" + notes_setting_asset)
	err = json.Unmarshal([]byte(notes_list), &notes)
	CheckErr(err)

	ranks := []model.RankInfo{}
	ranks = append(ranks, model.RankInfo{
		Rank:    5,
		RankMin: 0,
		RankMax: c_rank_score,
	}, model.RankInfo{
		Rank:    4,
		RankMin: c_rank_score + 1,
		RankMax: b_rank_score,
	}, model.RankInfo{
		Rank:    3,
		RankMin: b_rank_score + 1,
		RankMax: a_rank_score,
	}, model.RankInfo{
		Rank:    2,
		RankMin: a_rank_score + 1,
		RankMax: s_rank_score,
	}, model.RankInfo{
		Rank:    1,
		RankMin: s_rank_score + 1,
		RankMax: 0,
	})

	resp := model.PlayScoreResponseData{
		On: model.On{
			HasRecord: false,
			LiveInfo: model.LiveInfo{
				LiveDifficultyID: difficultyId,
				IsRandom:         false,
				AcFlag:           ac_flag,
				SwingFlag:        swing_flag,
				NotesList:        notes,
			},
		},
		Off: model.Off{
			HasRecord: false,
			LiveInfo: model.LiveInfo{
				LiveDifficultyID: difficultyId,
				IsRandom:         false,
				AcFlag:           ac_flag,
				SwingFlag:        swing_flag,
				NotesList:        notes,
			},
		},
		RankInfo:          ranks,
		CanActivateEffect: true,
		ServerTimestamp:   int(time.Now().Unix()),
	}

	m, err := json.Marshal(resp)
	CheckErr(err)

	res := model.Response{
		ResponseData: m,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}

	mm, err := json.Marshal(res)
	CheckErr(err)
	// fmt.Println(string(mm))

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(mm, "privatekey.pem")))

	ctx.String(http.StatusOK, string(mm))
}

func PlayRewardHandler(ctx *gin.Context) {
	playRewardReq := model.PlayRewardReq{}
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &playRewardReq)
	CheckErr(err)

	difficultyId := playRewardReq.LiveDifficultyID

	// Song type: normal / special
	// sqlite3 doesn't support FULL OUTER JOIN so use UNION ALL here.
	sql := `SELECT c_rank_score,b_rank_score,a_rank_score,s_rank_score,c_rank_combo,b_rank_combo,a_rank_combo,s_rank_combo,ac_flag,swing_flag FROM live_setting_m WHERE live_setting_id IN (SELECT live_setting_id FROM normal_live_m WHERE live_difficulty_id = ? UNION ALL SELECT live_setting_id FROM special_live_m WHERE live_difficulty_id = ?)`
	var c_rank_score, b_rank_score, a_rank_score, s_rank_score, c_rank_combo, b_rank_combo, a_rank_combo, s_rank_combo, ac_flag, swing_flag int
	err = MainEng.DB().QueryRow(sql, difficultyId, difficultyId).Scan(&c_rank_score, &b_rank_score, &a_rank_score, &s_rank_score, &c_rank_combo, &b_rank_combo, &a_rank_combo, &s_rank_combo, &ac_flag, &swing_flag)
	CheckErr(err)

	key := "live_deck_" + ctx.GetString("userid")
	deckId, err := database.LevelDb.Get([]byte(key))
	CheckErr(err)
	unitsList := []model.PlayRewardUnitList{}
	err = UserEng.Table("deck_unit_m").Join("LEFT", "user_deck_m", "deck_unit_m.user_deck_id = user_deck_m.id").
		Where("user_id = ? AND deck_id = ?", ctx.GetString("userid"), string(deckId)).Find(&unitsList)
	CheckErr(err)

	totalScore := playRewardReq.ScoreSmile + playRewardReq.ScoreCool + playRewardReq.ScoreCute
	resp := model.RewardResponseData{
		LiveInfo: []model.RewardLiveInfo{
			{
				LiveDifficultyID: difficultyId,
				IsRandom:         false,
				AcFlag:           ac_flag,
				SwingFlag:        swing_flag,
			},
		},
		TotalLove:   0,
		IsHighScore: true,
		HiScore:     totalScore,
		BaseRewardInfo: model.BaseRewardInfo{
			PlayerExp: 830,
			PlayerExpUnitMax: model.PlayerExpUnitMax{
				Before: 900,
				After:  900,
			},
			PlayerExpFriendMax: model.PlayerExpFriendMax{
				Before: 99,
				After:  99,
			},
			PlayerExpLpMax: model.PlayerExpLpMax{
				Before: 417,
				After:  417,
			},
			GameCoin:              4500,
			GameCoinRewardBoxFlag: false,
			SocialPoint:           10,
		},
		RewardUnitList: model.RewardUnitList{
			LiveClear: []model.LiveClear{},
			LiveRank:  []model.LiveRank{},
			LiveCombo: []interface{}{},
		},
		UnlockedSubscenarioIds:       []interface{}{},
		UnlockedMultiUnitScenarioIds: []interface{}{},
		EffortPoint:                  []model.EffortPoint{},
		IsEffortPointVisible:         false,
		LimitedEffortBox:             []interface{}{},
		UnitList:                     unitsList,
		BeforeUserInfo: model.BeforeUserInfo{
			Level:                          1028,
			Exp:                            28823566,
			PreviousExp:                    27734700,
			NextExp:                        28941885,
			GameCoin:                       86505544,
			SnsCoin:                        49,
			FreeSnsCoin:                    48,
			PaidSnsCoin:                    1,
			SocialPoint:                    1438165,
			UnitMax:                        5000,
			WaitingUnitMax:                 1000,
			CurrentEnergy:                  392,
			EnergyMax:                      417,
			TrainingEnergy:                 9,
			TrainingEnergyMax:              10,
			EnergyFullTime:                 "2023-03-20 01:28:55",
			LicenseLiveEnergyRecoverlyTime: 60,
			FriendMax:                      99,
			TutorialState:                  -1,
			OverMaxEnergy:                  0,
			UnlockRandomLiveMuse:           1,
			UnlockRandomLiveAqours:         1,
		},
		AfterUserInfo: model.AfterUserInfo{
			Level:                          1028,
			Exp:                            28824396,
			PreviousExp:                    27734700,
			NextExp:                        28941885,
			GameCoin:                       86520044,
			SnsCoin:                        50,
			FreeSnsCoin:                    49,
			PaidSnsCoin:                    1,
			SocialPoint:                    1438375,
			UnitMax:                        5000,
			WaitingUnitMax:                 1000,
			CurrentEnergy:                  392,
			EnergyMax:                      417,
			TrainingEnergy:                 9,
			TrainingEnergyMax:              10,
			EnergyFullTime:                 "2023-03-20 01:28:55",
			LicenseLiveEnergyRecoverlyTime: 60,
			FriendMax:                      99,
			TutorialState:                  -1,
			OverMaxEnergy:                  0,
			UnlockRandomLiveMuse:           1,
			UnlockRandomLiveAqours:         1,
		},
		NextLevelInfo: []model.NextLevelInfo{
			{
				Level:   1028,
				FromExp: 28823566,
			},
		},
		GoalAccompInfo: model.GoalAccompInfo{
			AchievedIds: []interface{}{},
			Rewards:     []interface{}{},
		},
		SpecialRewardInfo:    []interface{}{},
		EventInfo:            []interface{}{},
		DailyRewardInfo:      []interface{}{},
		CanSendFriendRequest: false,
		UsingBuffInfo:        []interface{}{},
		ClassSystem: model.ClassSystem{
			RankInfo: model.RewardRankInfo{
				BeforeClassRankID: 10,
				AfterClassRankID:  10,
				RankUpDate:        "2020-02-12 11:57:15",
			},
			CompleteFlag: false,
			IsOpened:     true,
			IsVisible:    true,
		},
		AccomplishedAchievementList:  []model.AccomplishedAchievementList{},
		UnaccomplishedAchievementCnt: 15,
		AddedAchievementList:         []interface{}{},
		MuseumInfo:                   model.RewardMuseumInfo{},
		UnitSupportList:              []model.RewardUnitSupportList{},
		ServerTimestamp:              1679238066,
		PresentCnt:                   2159,
	}

	if playRewardReq.MaxCombo > s_rank_combo {
		resp.ComboRank = 1
	} else if playRewardReq.MaxCombo > a_rank_combo {
		resp.ComboRank = 2
	} else if playRewardReq.MaxCombo > b_rank_combo {
		resp.ComboRank = 3
	} else if playRewardReq.MaxCombo > c_rank_combo {
		resp.ComboRank = 4
	} else {
		resp.ComboRank = 5
	}

	if totalScore > s_rank_score {
		resp.Rank = 1
	} else if totalScore > a_rank_score {
		resp.Rank = 2
	} else if totalScore > b_rank_score {
		resp.Rank = 3
	} else if totalScore > c_rank_score {
		resp.Rank = 4
	} else {
		resp.Rank = 5
	}

	m, err := json.Marshal(resp)
	CheckErr(err)

	res := model.Response{
		ResponseData: m,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}

	mm, err := json.Marshal(res)
	CheckErr(err)
	// fmt.Println(string(mm))

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(mm, "privatekey.pem")))

	ctx.String(http.StatusOK, string(mm))
}

package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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

	unitList := []model.UnitList{}
	for i := 0; i < 9; i++ {
		unitList = append(unitList, model.UnitList{
			Smile: 1,
			Cute:  1,
			Cool:  1,
		})
	}

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
			TotalSmile:       9,
			TotalCute:        9,
			TotalCool:        9,
			TotalHp:          999,
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
		ServerTimestamp:     1679237935,
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
	err = UserEng.Table("deck_unit_m").Select("*").
		Join("LEFT", "user_deck_m", "deck_unit_m.user_deck_id = user_deck_m.id").
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

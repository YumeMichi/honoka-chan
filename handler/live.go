package handler

import (
	"database/sql"
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
	"time"

	"github.com/gin-gonic/gin"
)

func PartyListHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	resp := utils.ReadAllText("assets/partylist.json")
	xms := encrypt.RSA_Sign_SHA1([]byte(resp), "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, resp)
}

func PlayLiveHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	playReq := model.PlayReq{}
	err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &playReq)
	if err != nil {
		panic(err)
	}

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	// resp := utils.ReadAllText("assets/playsong.json")
	db, err := sql.Open("sqlite3", "assets/live.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	tDifficultyId := playReq.LiveDifficultyID
	difficultyId, err := strconv.Atoi(tDifficultyId)
	if err != nil {
		panic(err)
	}
	deckId := playReq.UnitDeckID

	// Song type: normal / special
	// sqlite3 doesn't support FULL OUTER JOIN so use UNION ALL here.
	sql := `SELECT notes_list,c_rank_score,b_rank_score,a_rank_score,s_rank_score,ac_flag,swing_flag FROM live_setting_m WHERE live_setting_id IN (SELECT live_setting_id FROM normal_live_m WHERE live_difficulty_id = ? UNION ALL SELECT live_setting_id FROM special_live_m WHERE live_difficulty_id = ?)`
	rows, err := db.Query(sql, difficultyId, difficultyId)
	if err != nil {
		panic(err)
	}

	var notes_list string
	var c_rank_score, b_rank_score, a_rank_score, s_rank_score, ac_flag, swing_flag int
	for rows.Next() {
		err = rows.Scan(&notes_list, &c_rank_score, &b_rank_score, &a_rank_score, &s_rank_score, &ac_flag, &swing_flag)
		if err != nil {
			panic(err)
		}
	}

	// fmt.Println(len(notes_list))
	// fmt.Println(c_rank_score, b_rank_score, a_rank_score, s_rank_score)

	notes := []model.NotesList{}
	err = json.Unmarshal([]byte(notes_list), &notes)
	if err != nil {
		panic(err)
	}

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
			TotalHp:          99,
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
	if err != nil {
		panic(err)
	}

	res := model.Response{
		ResponseData: m,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}

	mm, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	xms := encrypt.RSA_Sign_SHA1(mm, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(mm))
}

func GameOverHandler(ctx *gin.Context) {
	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, "Fuck you!")
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	resp := utils.ReadAllText("assets/gameover.json")
	xms := encrypt.RSA_Sign_SHA1([]byte(resp), "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, resp)
}

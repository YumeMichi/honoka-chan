package rlive

import (
	"crypto/rand"
	"errors"
	livehandler "honoka-chan/internal/handler/live"
	"honoka-chan/internal/middleware"
	usermodel "honoka-chan/internal/model/user"
	"honoka-chan/internal/router"
	rliveschema "honoka-chan/internal/schema/rlive"
	"honoka-chan/internal/session"
	honokautils "honoka-chan/internal/utils"
	pkgutils "honoka-chan/pkg/utils"
	"math/big"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type randomLiveCandidateRow struct {
	LiveDifficultyID  int    `xorm:"live_difficulty_id"`
	NotesSettingAsset string `xorm:"notes_setting_asset"`
	AcFlag            int    `xorm:"ac_flag"`
	SwingFlag         int    `xorm:"swing_flag"`
}

func lot(ctx *gin.Context) {
	ss := session.Get(ctx)
	defer ss.FinalizeOrRollback()

	req := rliveschema.LotReq{}
	err := honokautils.ParseRequestData(ctx, &req)
	if ss.CheckErr(err) {
		return
	}

	if err := validateLotReq(req); ss.CheckErr(err) {
		return
	}

	randomLive, err := loadOrCreateRandomLive(ss, req)
	if ss.CheckErr(err) {
		return
	}

	partyListData, err := livehandler.BuildPartyListData(ss)
	if ss.CheckErr(err) {
		return
	}

	ss.Respond(rliveschema.LotResp{
		ResponseData: rliveschema.LotData{
			LiveInfo: rliveschema.LiveInfo{
				LiveDifficultyID: randomLive.LiveDifficultyID,
				IsRandom:         true,
				AcFlag:           randomLive.AcFlag,
				SwingFlag:        randomLive.SwingFlag,
			},
			HasSlideNotes:     randomLive.SwingFlag,
			PartyList:         partyListData.PartyList,
			TrainingEnergy:    partyListData.TrainingEnergy,
			TrainingEnergyMax: partyListData.TrainingEnergyMax,
			Token:             randomLive.Token,
		},
		ReleaseInfo: []any{},
		StatusCode:  http.StatusOK,
	})
}

type randomLiveResult struct {
	Token            string
	LiveDifficultyID int
	AcFlag           int
	SwingFlag        int
}

func loadOrCreateRandomLive(ss *session.Session, req rliveschema.LotReq) (randomLiveResult, error) {
	sessionRow := usermodel.UserLiveRandom{}
	has, err := ss.UserEng.Table(new(usermodel.UserLiveRandom)).
		Where("user_id = ?", ss.UserID).
		Where("attribute = ?", req.Attribute).
		Where("difficulty = ?", req.Difficulty).
		Where("member_category = ?", req.MemberCategory).
		Get(&sessionRow)
	if err != nil {
		return randomLiveResult{}, err
	}

	if has {
		liveInfo, err := loadRandomLiveInfo(ss, sessionRow.LiveDifficultyID)
		if err != nil {
			return randomLiveResult{}, err
		}
		return randomLiveResult{
			Token:            sessionRow.Token,
			LiveDifficultyID: sessionRow.LiveDifficultyID,
			AcFlag:           liveInfo.AcFlag,
			SwingFlag:        liveInfo.SwingFlag,
		}, nil
	}

	candidates, err := listRandomLiveCandidates(ss, req.Difficulty, req.Attribute, req.MemberCategory)
	if err != nil {
		return randomLiveResult{}, err
	}
	if len(candidates) == 0 {
		return randomLiveResult{}, errors.New("no random live candidates available")
	}

	index, err := cryptoRandIndex(len(candidates))
	if err != nil {
		return randomLiveResult{}, err
	}
	picked := candidates[index]
	token := pkgutils.RandomStr(64)
	if token == "" {
		return randomLiveResult{}, errors.New("failed to generate random live token")
	}

	_, err = ss.UserEng.Insert(&usermodel.UserLiveRandom{
		UserID:           ss.UserID,
		Attribute:        req.Attribute,
		Difficulty:       req.Difficulty,
		MemberCategory:   req.MemberCategory,
		Token:            token,
		LiveDifficultyID: picked.LiveDifficultyID,
		InProgress:       false,
	})
	if err != nil {
		return randomLiveResult{}, err
	}

	return randomLiveResult{
		Token:            token,
		LiveDifficultyID: picked.LiveDifficultyID,
		AcFlag:           picked.AcFlag,
		SwingFlag:        picked.SwingFlag,
	}, nil
}

func validateLotReq(req rliveschema.LotReq) error {
	if req.Difficulty < 1 || req.Difficulty > 4 {
		return errors.New("invalid difficulty")
	}
	if req.Attribute < 1 || req.Attribute > 3 {
		return errors.New("invalid attribute")
	}
	if req.MemberCategory < 0 {
		return errors.New("invalid member category")
	}
	return nil
}

func listRandomLiveCandidates(ss *session.Session, difficulty, attribute, memberCategory int) ([]randomLiveCandidateRow, error) {
	rows := []randomLiveCandidateRow{}

	sql := `
SELECT
	difficulty.live_difficulty_id,
	setting.notes_setting_asset,
	setting.ac_flag,
	setting.swing_flag
FROM live_setting_m AS setting
INNER JOIN (
	SELECT live_setting_id, live_difficulty_id FROM special_live_m
	UNION
	SELECT live_setting_id, live_difficulty_id FROM normal_live_m
) AS difficulty ON setting.live_setting_id = difficulty.live_setting_id
INNER JOIN live_track_m AS track ON track.live_track_id = setting.live_track_id
WHERE setting.difficulty = ? AND setting.attribute_icon_id = ?
`
	args := []any{difficulty, attribute}
	if memberCategory > 0 {
		sql += ` AND track.member_category = ?`
		args = append(args, memberCategory)
	}
	sql += `
ORDER BY difficulty.live_difficulty_id ASC
`

	err := ss.MainEng.SQL(sql, args...).Find(&rows)
	if err != nil {
		return nil, err
	}

	result := make([]randomLiveCandidateRow, 0, len(rows))
	for _, row := range rows {
		if !hasRandomLiveBeatmap(row.NotesSettingAsset) {
			continue
		}
		result = append(result, row)
	}
	return result, nil
}

func loadRandomLiveInfo(ss *session.Session, liveDifficultyID int) (randomLiveCandidateRow, error) {
	rows := []randomLiveCandidateRow{}
	err := ss.MainEng.SQL(`
SELECT
	difficulty.live_difficulty_id,
	setting.notes_setting_asset,
	setting.ac_flag,
	setting.swing_flag
FROM live_setting_m AS setting
INNER JOIN (
	SELECT live_setting_id, live_difficulty_id FROM special_live_m
	UNION
	SELECT live_setting_id, live_difficulty_id FROM normal_live_m
) AS difficulty ON setting.live_setting_id = difficulty.live_setting_id
WHERE difficulty.live_difficulty_id = ?
LIMIT 1
`, liveDifficultyID).Find(&rows)
	if err != nil {
		return randomLiveCandidateRow{}, err
	}
	if len(rows) == 0 {
		return randomLiveCandidateRow{}, errors.New("random live difficulty not found")
	}
	return rows[0], nil
}

func hasRandomLiveBeatmap(notesSettingAsset string) bool {
	if notesSettingAsset == "" {
		return false
	}
	path := filepath.Join("assets/serverdata/beatmaps", notesSettingAsset)
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func cryptoRandIndex(length int) (int, error) {
	if length <= 0 {
		return 0, errors.New("invalid random candidate length")
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

func init() {
	router.AddHandler("main.php", "POST", "/rlive/lot", middleware.Common, lot)
}

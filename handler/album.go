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
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func AlbumSeriesAllHandler(ctx *gin.Context) {
	db, err := sql.Open("sqlite3", "assets/main.db")
	CheckErr(err)
	defer db.Close()

	reqTime := time.Now().Unix()

	authorizeStr := ctx.Request.Header["Authorize"]
	authToken, err := utils.GetAuthorizeToken(authorizeStr)
	if err != nil {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	userId := ctx.Request.Header[http.CanonicalHeaderKey("User-ID")]
	if len(userId) == 0 {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	if !database.MatchTokenUid(authToken, userId[0]) {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	nonce, err := utils.GetAuthorizeNonce(authorizeStr)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}
	nonce++

	respTime := time.Now().Unix()
	newAuthorizeStr := fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", respTime, authToken, nonce, userId[0], reqTime)
	// fmt.Println(newAuthorizeStr)

	//
	sql := `SELECT album_series_id FROM album_series_m`
	seriesRows, err := db.Query(sql)
	CheckErr(err)
	albumSeriesAllResp := []model.AlbumSeriesResp{}
	for seriesRows.Next() {
		var series int
		err = seriesRows.Scan(&series)
		CheckErr(err)

		albumSeriesAll := []model.AlbumResult{}
		stmt, err := db.Prepare("SELECT unit_id,rarity FROM unit_m WHERE album_series_id = ?")
		CheckErr(err)

		unitRows, err := stmt.Query(series)
		CheckErr(err)

		for unitRows.Next() {
			var unitId, rarity int
			err = unitRows.Scan(&unitId, &rarity)
			CheckErr(err)

			albumSeries := model.AlbumResult{
				UnitID:           unitId,
				RankMaxFlag:      true,
				LoveMaxFlag:      true,
				RankLevelMaxFlag: true,
				AllMaxFlag:       true,
				TotalLove:        10000,
				FavoritePoint:    1000,
			}

			if rarity != 4 {
				switch rarity {
				case 1:
					// N
					albumSeries.HighestLovePerUnit = 50
				case 2:
					// R
					albumSeries.HighestLovePerUnit = 200
				case 3:
					// SR
					albumSeries.HighestLovePerUnit = 500
				case 5:
					// SSR
					albumSeries.HighestLovePerUnit = 750
				}
			} else {
				// UR
				albumSeries.HighestLovePerUnit = 1000

				// IsSigned
				stmt, err = db.Prepare("SELECT COUNT(*) AS ct FROM unit_sign_asset_m WHERE unit_id = ?")
				CheckErr(err)
				signRows, err := stmt.Query(unitId)
				CheckErr(err)

				ct := 0
				for signRows.Next() {
					err = signRows.Scan(&ct)
					CheckErr(err)
				}
				if ct > 0 {
					albumSeries.SignFlag = true
				} else {
					albumSeries.SignFlag = false
				}
			}

			albumSeriesAll = append(albumSeriesAll, albumSeries)
		}
		albumSeriesAllResp = append(albumSeriesAllResp, model.AlbumSeriesResp{
			SeriesID: series,
			UnitList: albumSeriesAll,
		})
	}
	rb, err := json.Marshal(albumSeriesAllResp)
	CheckErr(err)
	resp := model.Response{
		ResponseData: rb,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	respb, err := json.Marshal(resp)
	CheckErr(err)
	// fmt.Println(string(respb))

	xms := encrypt.RSA_Sign_SHA1(respb, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("Server-Version", config.Conf.Server.VersionNumber)
	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)

	ctx.String(http.StatusOK, string(respb))
}

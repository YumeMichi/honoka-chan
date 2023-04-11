package handler

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"honoka-chan/utils"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func DownloadAdditionalHandler(ctx *gin.Context) {
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

	downloadReq := model.AdditionalReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	pkgList := []model.AdditionalResult{}
	if CdnUrl != "" {
		pkgType, pkgId := downloadReq.PackageType, downloadReq.PackageID
		stmt, err := db.Prepare("SELECT pkg_order,pkg_size FROM download_db WHERE pkg_type = ? AND pkg_id = ? ORDER BY pkg_order ASC")
		CheckErr(err)
		rows, err := stmt.Query(pkgType, pkgId)
		CheckErr(err)
		for rows.Next() {
			var pkgOrder, pkgSize int
			err = rows.Scan(&pkgOrder, &pkgSize)
			CheckErr(err)
			pkgList = append(pkgList, model.AdditionalResult{
				Size: pkgSize,
				URL:  fmt.Sprintf("%s/archives/%d_%d_%d.zip", CdnUrl, pkgType, pkgId, pkgOrder),
			})
		}
	}

	addResp := model.AdditionalResp{
		ResponseData: pkgList,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(addResp)
	CheckErr(err)
	// fmt.Println(string(resp))
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

func DownloadBatchHandler(ctx *gin.Context) {
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

	downloadReq := model.BatchReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	pkgList := []model.BatchResult{}
	if downloadReq.ClientVersion == PackageVersion && CdnUrl != "" {
		pkgType := downloadReq.PackageType
		exPkgId := ""
		if len(downloadReq.ExcludedPackageIds) > 0 {
			// https://stackoverflow.com/a/37533144
			exPkgId = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(downloadReq.ExcludedPackageIds)), ","), "[]")

			exPkgId = fmt.Sprintf("AND pkg_id NOT IN (%s)", exPkgId)
		}
		// fmt.Println(exPkgId)
		stmt, err := db.Prepare("SELECT pkg_id,pkg_order,pkg_size FROM download_db WHERE pkg_type = ? " + exPkgId + " ORDER BY pkg_id ASC, pkg_order ASC")
		CheckErr(err)
		rows, err := stmt.Query(pkgType)
		CheckErr(err)
		for rows.Next() {
			var pkgId, pkgOrder, pkgSize int
			err = rows.Scan(&pkgId, &pkgOrder, &pkgSize)
			CheckErr(err)
			pkgList = append(pkgList, model.BatchResult{
				Size: pkgSize,
				URL:  fmt.Sprintf("%s/archives/%d_%d_%d.zip", CdnUrl, pkgType, pkgId, pkgOrder),
			})
		}
	}

	batchResp := model.BatchResp{
		ResponseData: pkgList,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(batchResp)
	CheckErr(err)
	// fmt.Println(string(resp))
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

func DownloadUpdateHandler(ctx *gin.Context) {
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

	downloadReq := model.UpdateReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	pkgList := []model.UpdateResult{}
	if downloadReq.ExternalVersion != PackageVersion && CdnUrl != "" {
		pkgType := 99
		rows, err := db.Query("SELECT pkg_id,pkg_order,pkg_size FROM download_db WHERE pkg_type = 99 ORDER BY pkg_id ASC, pkg_order ASC")
		CheckErr(err)
		for rows.Next() {
			var pkgId, pkgOrder, pkgSize int
			err = rows.Scan(&pkgId, &pkgOrder, &pkgSize)
			CheckErr(err)
			pkgList = append(pkgList, model.UpdateResult{
				Size:    pkgSize,
				URL:     fmt.Sprintf("%s/archives/%d_%d_%d.zip", CdnUrl, pkgType, pkgId, pkgOrder),
				Version: PackageVersion,
			})
		}
	}

	updateResp := model.UpdateResp{
		ResponseData: pkgList,
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(updateResp)
	CheckErr(err)
	// fmt.Println(string(resp))
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

func DownloadUrlHandler(ctx *gin.Context) {
	// Extract SQL: SELECT CAST(pkg_type AS TEXT) || '_' || CAST(pkg_id AS TEXT) || '_' || CAST(pkg_order AS TEXT) || '.zip' AS zip_name FROM download_db ORDER BY pkg_type ASC,pkg_id ASC, pkg_order ASC;
	// Extract Cmd: cat list.txt | while read line; do; unzip -o $line; done
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

	downloadReq := model.UrlReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	urlList := []string{}
	for _, v := range downloadReq.PathList {
		urlList = append(urlList, fmt.Sprintf("%s/extracted/%s", CdnUrl, strings.ReplaceAll(v, "\\", "")))
	}
	urlResp := model.UrlResp{
		ResponseData: model.UrlResult{
			UrlList: urlList,
		},
		ReleaseInfo: []interface{}{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(urlResp)
	CheckErr(err)
	// fmt.Println(string(resp))
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

func DownloadEventHandler(ctx *gin.Context) {
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

	eventResp := model.EventResp{
		ResponseData: []interface{}{},
		ReleaseInfo:  []interface{}{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(eventResp)
	CheckErr(err)
	xms := encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")
	xms64 := base64.RawStdEncoding.EncodeToString(xms)

	ctx.Header("user_id", userId[0])
	ctx.Header("authorize", newAuthorizeStr)
	ctx.Header("X-Message-Sign", xms64)
	ctx.String(http.StatusOK, string(resp))
}

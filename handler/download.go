package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/config"
	"honoka-chan/encrypt"
	"honoka-chan/model"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"xorm.io/builder"
)

type PkgInfo struct {
	Id    int `xorm:"pkg_id"`
	Order int `xorm:"pkg_order"`
	Size  int `xorm:"pkg_size"`
}

func DownloadAdditional(ctx *gin.Context) {
	downloadReq := model.AdditionalReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	pkgList := []model.AdditionalRes{}
	if CdnUrl != "" {
		pkgType, pkgId := downloadReq.PackageType, downloadReq.PackageID
		var pkgInfo []PkgInfo
		err := MainEng.Table("download_m").Where("pkg_type = ? AND pkg_id = ? AND pkg_os = ?", pkgType, pkgId, downloadReq.TargetOs).
			Cols("pkg_id,pkg_order,pkg_size").
			OrderBy("pkg_id ASC, pkg_order ASC").Find(&pkgInfo)
		CheckErr(err)

		for _, pkg := range pkgInfo {
			pkgList = append(pkgList, model.AdditionalRes{
				Size: pkg.Size,
				URL:  fmt.Sprintf("%s/%s/archives/%d_%d_%d.zip", CdnUrl, downloadReq.TargetOs, pkgType, pkg.Id, pkg.Order),
			})
		}
	}

	addResp := model.AdditionalResp{
		ResponseData: pkgList,
		ReleaseInfo:  []any{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(addResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func DownloadBatch(ctx *gin.Context) {
	downloadReq := model.BatchReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	pkgList := []model.BatchRes{}
	if downloadReq.ClientVersion == config.PackageVersion && CdnUrl != "" {
		pkgType := downloadReq.PackageType
		var pkgInfo []PkgInfo
		err := MainEng.Table("download_m").Where(builder.NotIn("pkg_id", downloadReq.ExcludedPackageIds)).Where("pkg_type = ? AND pkg_os = ?", pkgType, downloadReq.Os).
			Cols("pkg_id,pkg_order,pkg_size").
			OrderBy("pkg_id ASC, pkg_order ASC").Find(&pkgInfo)
		CheckErr(err)

		for _, pkg := range pkgInfo {
			pkgList = append(pkgList, model.BatchRes{
				Size: pkg.Size,
				URL:  fmt.Sprintf("%s/%s/archives/%d_%d_%d.zip", CdnUrl, downloadReq.Os, pkgType, pkg.Id, pkg.Order),
			})
		}
	}

	batchResp := model.BatchResp{
		ResponseData: pkgList,
		ReleaseInfo:  []any{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(batchResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func DownloadUpdate(ctx *gin.Context) {
	downloadReq := model.UpdateReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	pkgList := []model.UpdateRes{}
	if downloadReq.ExternalVersion != config.PackageVersion && CdnUrl != "" {
		pkgType := 99
		var pkgInfo []PkgInfo
		err := MainEng.Table("download_m").Where("pkg_type = ? AND pkg_os = ?", pkgType, downloadReq.TargetOs).
			Cols("pkg_id,pkg_order,pkg_size").
			OrderBy("pkg_id ASC, pkg_order ASC").Find(&pkgInfo)
		CheckErr(err)

		for _, pkg := range pkgInfo {
			pkgList = append(pkgList, model.UpdateRes{
				Size:    pkg.Size,
				URL:     fmt.Sprintf("%s/%s/archives/%d_%d_%d.zip", CdnUrl, downloadReq.TargetOs, pkgType, pkg.Id, pkg.Order),
				Version: config.PackageVersion,
			})
		}
	}

	updateResp := model.UpdateResp{
		ResponseData: pkgList,
		ReleaseInfo:  []any{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(updateResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func DownloadUrl(ctx *gin.Context) {
	// Extract SQL: SELECT CAST(pkg_type AS TEXT) || '_' || CAST(pkg_id AS TEXT) || '_' || CAST(pkg_order AS TEXT) || '.zip' AS zip_name FROM download_m ORDER BY pkg_type ASC,pkg_id ASC, pkg_order ASC;
	// Extract Cmd: cat list.txt | while read line; do; unzip -o $line; done
	downloadReq := model.UrlReq{}
	if err := json.Unmarshal([]byte(ctx.PostForm("request_data")), &downloadReq); err != nil {
		panic(err)
	}
	urlList := []string{}
	for _, v := range downloadReq.PathList {
		urlList = append(urlList, fmt.Sprintf("%s/%s/extracted/%s", CdnUrl, downloadReq.Os, strings.ReplaceAll(v, "\\", "")))
	}
	urlResp := model.UrlResp{
		ResponseData: model.UrlRes{
			UrlList: urlList,
		},
		ReleaseInfo: []any{},
		StatusCode:  200,
	}
	resp, err := json.Marshal(urlResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

func DownloadEvent(ctx *gin.Context) {
	eventResp := model.EventResp{
		ResponseData: []any{},
		ReleaseInfo:  []any{},
		StatusCode:   200,
	}
	resp, err := json.Marshal(eventResp)
	CheckErr(err)

	nonce := ctx.GetInt("nonce")
	nonce++

	ctx.Header("user_id", ctx.GetString("userid"))
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), ctx.GetString("token"), nonce, ctx.GetString("userid"), ctx.GetInt64("req_time")))
	ctx.Header("X-Message-Sign", base64.StdEncoding.EncodeToString(encrypt.RSA_Sign_SHA1(resp, "privatekey.pem")))

	ctx.String(http.StatusOK, string(resp))
}

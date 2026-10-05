package middleware

import (
	"bytes"
	"fmt"
	"honoka-chan/internal/session"
	honokautils "honoka-chan/internal/utils"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Common(ctx *gin.Context) {
	// The iOS client sometimes sends an unquoted boundary containing '/'.
	// Go rejects that Content-Type before reading the otherwise valid body.
	var rawBody []byte
	contentType := ctx.GetHeader("Content-Type")
	_, _, mediaTypeErr := mime.ParseMediaType(contentType)
	if mediaTypeErr != nil && strings.HasPrefix(contentType, "multipart/form-data; boundary=") &&
		ctx.Request.ContentLength > 0 && ctx.Request.ContentLength <= 1<<20 {
		if body, err := io.ReadAll(ctx.Request.Body); err == nil {
			rawBody = body
			ctx.Request.Body = io.NopCloser(bytes.NewReader(body))
		}
	}
	reqData := ""
	form, formErr := ctx.MultipartForm()
	if formErr == nil {
		if v, ok := form.Value["request_data"]; ok && len(v) > 0 {
			reqData = v[0]
		}
	}
	if reqData == "" {
		reqData = ctx.PostForm("request_data")
	}
	if reqData == "" {
		reqData = honokautils.RecoverMultipartRequestData(contentType, rawBody)
	}
	ctx.Set("request_data", reqData)

	uid := ctx.GetHeader("User-ID")

	ss := session.Attach(ctx)
	ctx.Set("session", ss)
	defer ss.FinalizeOrRollback()

	if ctx.IsAborted() {
		return
	}

	authorize := ctx.GetHeader("Authorize")
	params, err := url.ParseQuery(authorize)
	if err != nil {
		ss.AbortWithStatus(http.StatusNotFound, honokautils.NewNotFoundContent(ctx.Request.URL.Path))
		return
	}

	nonce, _ := strconv.Atoi(params.Get("nonce"))
	nonce++
	ctx.Set("nonce", nonce)

	token := params.Get("token")
	ctx.Set("token", token)

	if !honokautils.IsMainLoginEndpoint(ctx.Request.URL.Path) {
		userID, err := strconv.Atoi(uid)
		if err != nil || userID <= 0 {
			ss.AbortWithStatus(http.StatusNotFound, honokautils.NewNotFoundContent(ctx.Request.URL.Path))
			return
		}

		valid, err := ss.IsAuthorizeTokenForUser(token, userID)
		if ss.CheckErr(err) {
			return
		}
		if !valid {
			ss.AbortWithStatus(http.StatusNotFound, honokautils.NewNotFoundContent(ctx.Request.URL.Path))
			return
		}

		ctx.Set("userid", uid)
		ss.LoadUser(uid)
		if ss.Done() {
			return
		}

		if ss.UserPref.ForceRelogin {
			ss.AbortWithStatus(http.StatusNotFound, honokautils.NewNotFoundContent(ctx.Request.URL.Path))
			return
		}
	}

	ctx.Header("user_id", uid)
	ctx.Header("authorize", fmt.Sprintf("consumerKey=lovelive_test&timeStamp=%d&version=1.1&token=%s&nonce=%d&user_id=%s&requestTimeStamp=%d", time.Now().Unix(), token, nonce, uid, time.Now().Unix()))

	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.Header("X-Powered-By", "KLab Native APP Platform")
	ctx.Header("server_version", "20120129")
	ctx.Header("Server-Version", "97.4.6")
	ctx.Header("version_up", "0")
	ctx.Header("status_code", strconv.Itoa(http.StatusOK))

	ctx.Next()
}

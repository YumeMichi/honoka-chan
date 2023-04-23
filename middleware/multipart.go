package middleware

import (
	"io"
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseMultipartForm(ctx *gin.Context) {
	if ctx.Request.Header.Get("OS") != "Android" {
		// I don't know why mime.ParseMediaType() is failed
		// mime.ParseMediaType(ctx.Request.Header.Get("Content-Type"))
		boundary := strings.ReplaceAll(ctx.Request.Header.Get("Content-Type"), "multipart/form-data; boundary=", "")

		var reqData []byte
		mReader := multipart.NewReader(ctx.Request.Body, boundary)
		for {
			part, err := mReader.NextPart()
			if err == io.EOF {
				break
			}
			CheckErr(err)

			data, err := io.ReadAll(part)
			CheckErr(err)

			reqData = data
		}
		ctx.Set("request_data", string(reqData))
	} else {
		ctx.Set("request_data", ctx.PostForm("request_data"))
	}

	ctx.Next()
}

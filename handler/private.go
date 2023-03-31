package handler

import (
	"encoding/base64"
	"fmt"
	"honoka-chan/encrypt"
	"honoka-chan/utils"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/forgoer/openssl"
	"github.com/gin-gonic/gin"
)

var (
	randKey string
)

func ActiveHandler(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, `{ "code": 0, "msg": "ok", "data": { "message": "ok", "result": 0 } }`)
}

func PublicKeyHandler(ctx *gin.Context) {
	publicKey := utils.ReadAllText("publickey.pem")
	publicKey = strings.ReplaceAll(publicKey, "\n", "")
	publicKey = strings.ReplaceAll(publicKey, "-----BEGIN PUBLIC KEY-----", "")
	publicKey = strings.ReplaceAll(publicKey, "-----END PUBLIC KEY-----", "")
	publicKey = strings.ReplaceAll(publicKey, "/", "\\/")
	// fmt.Println(publicKey)
	resp := fmt.Sprintf(`{ "code": 0, "msg": "", "data": { "result": 0, "message": "ok", "key": "%s", "method": "rsa" } }`, publicKey)
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func HandshakeHandler(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		panic(err)
	}
	defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	body64, err := base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		panic(err)
	}
	decryptedBody := encrypt.RSA_Decrypt(body64, "privatekey.pem")
	// fmt.Println(decryptedBody)
	// fmt.Println(string(decryptedBody))

	params, err := url.ParseQuery(string(decryptedBody))
	if err != nil {
		panic(err)
	}
	randKey = params.Get("randkey")
	// fmt.Println(randKey)

	token := `{"token":"YUMEMICHI14104EF796D769AE404C1FC6"}`
	encryptedToken, err := openssl.Des3ECBEncrypt([]byte(token), []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}
	encryptedToken64 := base64.StdEncoding.EncodeToString(encryptedToken)
	// fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedToken64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func InitializeHandler(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	// body64, err := base64.StdEncoding.DecodeString(string(body))
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(body64))

	// decryptedBody, err := openssl.Des3ECBDecrypt(body64, []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(decryptedBody))

	token := `{"message":"Hello, world!"}`
	encryptedToken, err := openssl.Des3ECBEncrypt([]byte(token), []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}
	encryptedToken64 := base64.StdEncoding.EncodeToString(encryptedToken)
	// fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedToken64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func GetCodeHandler(ctx *gin.Context) {
	resp := `{ "code": 0, "msg": "ok", "data": { "codeArray": [ { "btntext": "好的", "code": "-10264022", "msg_from": 2, "text": "", "title": "短信验证码被阻止", "type": 1 }, { "btntext": " ", "code": "-10869623", "msg_from": 2, "text": "", "title": "网络连接失败，无法一键登录", "type": 2 }, { "btntext": " ", "code": "10298300", "msg_from": 2, "text": "", "title": " ", "type": 3 }, { "btntext": "", "code": "10298311", "msg_from": 2, "text": "", "title": "", "type": 3 }, { "btntext": " ", "code": "10298312", "msg_from": 2, "text": "", "title": " ", "type": 3 }, { "btntext": " ", "code": "10298313", "msg_from": 2, "text": "", "title": " ", "type": 1 }, { "btntext": " ", "code": "10298321", "msg_from": 2, "text": "", "title": " ", "type": 3 }, { "btntext": " ", "code": "10298322", "msg_from": 2, "text": "", "title": " ", "type": 3 } ], "codeVersion": "1.0.5" } }`
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func LoginAutoHandler(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	// body64, err := base64.StdEncoding.DecodeString(string(body))
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(body64))

	// decryptedBody, err := openssl.Des3ECBDecrypt(body64, []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(decryptedBody))

	token := `{"message":"Hello, world!"}`
	encryptedToken, err := openssl.Des3ECBEncrypt([]byte(token), []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}
	encryptedToken64 := base64.StdEncoding.EncodeToString(encryptedToken)
	// fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedToken64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func LoginAreaHandler(ctx *gin.Context) {
	userId := ctx.PostForm("userid")
	if userId != "" {
		// fmt.Println(userId)
		resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": { "userid": "%s" } }`, userId)
		ctx.Header("Content-Type", "text/html;charset=utf-8")
		ctx.String(http.StatusOK, resp)
	}
}

func AccountLoginHandler(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	// body64, err := base64.StdEncoding.DecodeString(string(body))
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(body64))

	// decryptedBody, err := openssl.Des3ECBDecrypt(body64, []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(decryptedBody))

	token := `{"message":"Hello, world!"}`
	encryptedToken, err := openssl.Des3ECBEncrypt([]byte(token), []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}
	encryptedToken64 := base64.StdEncoding.EncodeToString(encryptedToken)
	// fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedToken64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func ReportRoleHandler(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	// body64, err := base64.StdEncoding.DecodeString(string(body))
	// if err != nil {
	// 	panic(err)
	// }
	// // fmt.Println(string(body64))

	// decryptedBody, err := openssl.Des3ECBDecrypt(body64, []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(decryptedBody))

	// decrypted, err := url.QueryUnescape(string(decryptedBody))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(decrypted)

	token := `{"message":"Hello, world!"}`
	encryptedToken, err := openssl.Des3ECBEncrypt([]byte(token), []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}
	encryptedToken64 := base64.StdEncoding.EncodeToString(encryptedToken)
	fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedToken64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func GetProductListHandler(ctx *gin.Context) {
	resp := `{ "code": 0, "msg": "ok", "data": { "message": [ ], "result": 0 } }`
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func ReportLog(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, "")
}

func ReportApp(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))
	resp := `{ "code": 0, "msg": "", "data": { "needReport": 0 } }`
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func AgreementHandler(ctx *gin.Context) {
	resp := `{ "return_code": 0, "error_type": 0, "return_message": "", "data": { } }`
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"honoka-chan/database"
	"honoka-chan/encrypt"
	"honoka-chan/tools"
	"honoka-chan/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/forgoer/openssl"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type LoginResp struct {
	Activation           int    `json:"activation"`
	Autokey              string `json:"autokey"`
	CaptchaParams        string `json:"captchaParams"`
	CheckCodeGUID        string `json:"checkCodeGuid"`
	CheckCodeURL         string `json:"checkCodeUrl"`
	HasExtendAccs        int    `json:"hasExtendAccs"`
	HasRealInfo          int    `json:"has_realInfo"`
	ImagecodeType        int    `json:"imagecodeType"`
	IsNewUser            int    `json:"isNewUser"`
	Message              string `json:"message"`
	NextAction           int    `json:"nextAction"`
	PromptMsg            string `json:"prompt_msg"`
	RealInfoNotification string `json:"realInfoNotification"`
	RealInfoForce        int    `json:"realInfo_force"`
	RealInfoForcePay     int    `json:"realInfo_force_pay"`
	RealInfoStatus       int    `json:"realInfo_status"`
	RealInfoStatusPay    int    `json:"realInfo_status_pay"`
	Result               int    `json:"result"`
	SdgHeight            int    `json:"sdg_height"`
	SdgWidth             int    `json:"sdg_width"`
	Ticket               string `json:"ticket"`
	UserAttribute        string `json:"userAttribute"`
	Userid               string `json:"userid"`
}

type LoginAutoResp struct {
	Result  int    `json:"result"`
	Message string `json:"message"`
	Autokey string `json:"autokey"`
	Userid  string `json:"userid"`
	Ticket  string `json:"ticket"`
}

type InitializeResp struct {
	BrandLogo                 string        `json:"brand_logo"`
	BrandName                 string        `json:"brand_name"`
	DaoyuClientid             string        `json:"daoyu_clientid"`
	DaoyuDownloadURL          string        `json:"daoyu_download_url"`
	DeviceFeature             string        `json:"device_feature"`
	DisplayThirdaccout        int           `json:"display_thirdaccout"`
	ForceShowAgreement        int           `json:"force_show_agreement"`
	GreportLogLevel           string        `json:"greport_log_level"`
	GuestEnable               int           `json:"guest_enable"`
	IsMatch                   int           `json:"is_match"`
	LogLevel                  string        `json:"log_level"`
	LoginButton               []string      `json:"login_button"`
	LoginIcon                 []interface{} `json:"login_icon"`
	LoginLimitEnable          int           `json:"login_limit_enable"`
	NeedFloatWindowPermission int           `json:"need_float_window_permission"`
	NewDeviceIDServer         string        `json:"new_device_id_server"`
	QqAppID                   string        `json:"qq_appId"`
	QqKey                     string        `json:"qq_key"`
	ShowGuestConfirm          int           `json:"show_guest_confirm"`
	VoicetipButton            int           `json:"voicetip_button"`
	VoicetipOne               string        `json:"voicetip_one"`
	VoicetipTwo               string        `json:"voicetip_two"`
	WegameAppid               string        `json:"wegame_appid"`
	WegameAppkey              string        `json:"wegame_appkey"`
	WegameClientid            string        `json:"wegame_clientid"`
	WegameCompanyID           string        `json:"wegame_companyId"`
	WegameLoginURL            string        `json:"wegame_loginUrl"`
	WeiboAppKey               string        `json:"weibo_appKey"`
	WeiboRedirectURL          string        `json:"weibo_redirectUrl"`
	WeixinAppID               string        `json:"weixin_appId"`
	WeixinKey                 string        `json:"weixin_key"`
}

func ActiveHandler(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// CheckErr(err)
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

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
	CheckErr(err)
	defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	body64, err := base64.StdEncoding.DecodeString(string(body))
	CheckErr(err)
	decryptedBody := encrypt.RSA_Decrypt(body64, "privatekey.pem")
	// fmt.Println(decryptedBody)
	// fmt.Println(string(decryptedBody))

	params, err := url.ParseQuery(string(decryptedBody))
	CheckErr(err)
	randKey := params.Get("randkey")
	deviceId := params.Get("deviceid")
	// fmt.Println(randKey)
	// fmt.Println(deviceId)

	err = database.LevelDb.Put([]byte(deviceId), []byte(randKey))
	CheckErr(err)

	token := strings.ToUpper(utils.RandomStr(33))
	token = fmt.Sprintf(`{"message":"ok","result":0,"token":"%s"}`, token)
	encryptedToken, err := openssl.Des3ECBEncrypt([]byte(token), []byte(randKey)[0:24], openssl.PKCS7_PADDING)
	CheckErr(err)
	encryptedToken64 := base64.StdEncoding.EncodeToString(encryptedToken)
	// fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedToken64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func InitializeHandler(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// CheckErr(err)
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	// body64, err := base64.StdEncoding.DecodeString(string(body))
	// CheckErr(err)
	// fmt.Println(string(body64))

	deviceId := ctx.Request.Header.Get("X-DEVICEID")
	randKey, err := database.LevelDb.Get([]byte(deviceId))
	CheckErr(err)
	// decryptedBody, err := openssl.Des3ECBDecrypt(body64, randKey[0:24], openssl.PKCS7_PADDING)
	// CheckErr(err)
	// fmt.Println(string(decryptedBody))

	initResp := InitializeResp{
		BrandLogo:                 "http://gskd.sdo.com/ghome/ztc/logo/og/logo_xhdpi.png",
		BrandName:                 "盛趣游戏",
		ForceShowAgreement:        1,
		GreportLogLevel:           "off",
		LogLevel:                  "off",
		LoginButton:               []string{"official"},
		LoginIcon:                 []interface{}{},
		NeedFloatWindowPermission: 1,
		NewDeviceIDServer:         strings.ToUpper(openssl.Md5ToString(deviceId)),
		ShowGuestConfirm:          1,
		VoicetipButton:            1,
	}
	data, err := json.Marshal(initResp)
	CheckErr(err)
	encryptedData, err := openssl.Des3ECBEncrypt([]byte(data), randKey[0:24], openssl.PKCS7_PADDING)
	CheckErr(err)
	encryptedData64 := base64.StdEncoding.EncodeToString(encryptedData)
	// fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedData64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func GetCodeHandler(ctx *gin.Context) {
	resp := `{ "code": 0, "msg": "ok", "data": { "codeArray": [ { "btntext": "好的", "code": "-10264022", "msg_from": 2, "text": "", "title": "短信验证码被阻止", "type": 1 }, { "btntext": " ", "code": "-10869623", "msg_from": 2, "text": "", "title": "网络连接失败，无法一键登录", "type": 2 }, { "btntext": " ", "code": "10298300", "msg_from": 2, "text": "", "title": " ", "type": 3 }, { "btntext": "", "code": "10298311", "msg_from": 2, "text": "", "title": "", "type": 3 }, { "btntext": " ", "code": "10298312", "msg_from": 2, "text": "", "title": " ", "type": 3 }, { "btntext": " ", "code": "10298313", "msg_from": 2, "text": "", "title": " ", "type": 1 }, { "btntext": " ", "code": "10298321", "msg_from": 2, "text": "", "title": " ", "type": 3 }, { "btntext": " ", "code": "10298322", "msg_from": 2, "text": "", "title": " ", "type": 3 } ], "codeVersion": "1.0.5" } }`
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func LoginAutoHandler(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	CheckErr(err)
	defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	body64, err := base64.StdEncoding.DecodeString(string(body))
	CheckErr(err)
	// fmt.Println(string(body64))

	deviceId := ctx.Request.Header.Get("X-DEVICEID")
	randKey, err := database.LevelDb.Get([]byte(deviceId))
	CheckErr(err)

	decryptedBody, err := openssl.Des3ECBDecrypt(body64, randKey[0:24], openssl.PKCS7_PADDING)
	CheckErr(err)
	queryStr := string(decryptedBody)
	// fmt.Println(queryStr)

	params, err := url.ParseQuery(queryStr)
	CheckErr(err)
	autoKey := params.Get("autokey")
	// fmt.Println(autoKey)
	if autoKey == "" {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	stmt, err := UserEng.DB().Prepare("SELECT userid,ticket AS ct FROM users WHERE autokey = ?")
	CheckErr(err)
	defer stmt.Close()

	var userid, ticket string
	err = stmt.QueryRow(autoKey).Scan(&userid, &ticket)
	CheckErr(err)

	var resp string
	if userid != "" {
		autoResp := LoginAutoResp{
			Result:  0,
			Message: "ok",
			Autokey: autoKey,
			Userid:  userid,
			Ticket:  ticket,
		}
		data, err := json.Marshal(autoResp)
		// fmt.Println(string(data))
		CheckErr(err)
		encryptedData, err := openssl.Des3ECBEncrypt(data, randKey[0:24], openssl.PKCS7_PADDING)
		CheckErr(err)
		encryptedData64 := base64.StdEncoding.EncodeToString(encryptedData)
		// fmt.Println(encryptedData64)

		resp = fmt.Sprintf(`{ "code": 0, "msg": "ok", "data": "%s" }`, encryptedData64)
	} else {
		data := `{"message":"账号不存在或者登陆状态已过期！","result":31}`
		encryptedData, err := openssl.Des3ECBEncrypt([]byte(data), randKey[0:24], openssl.PKCS7_PADDING)
		CheckErr(err)
		encryptedData64 := base64.StdEncoding.EncodeToString(encryptedData)
		// fmt.Println(encryptedData64)

		resp = fmt.Sprintf(`{ "code": 31, "msg": "账号不存在或者登陆状态已过期！", "data": "%s" }`, encryptedData64)
	}

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
	body, err := io.ReadAll(ctx.Request.Body)
	CheckErr(err)
	defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	body64, err := base64.StdEncoding.DecodeString(string(body))
	CheckErr(err)
	// fmt.Println(string(body64))

	deviceId := ctx.Request.Header.Get("X-DEVICEID")
	randKey, err := database.LevelDb.Get([]byte(deviceId))
	CheckErr(err)

	decryptedBody, err := openssl.Des3ECBDecrypt(body64, randKey[0:24], openssl.PKCS7_PADDING)
	CheckErr(err)
	queryStr, err := url.QueryUnescape(string(decryptedBody))
	CheckErr(err)

	params, err := url.ParseQuery(queryStr)
	CheckErr(err)

	phone, password := params.Get("phone"), params.Get("password")
	if phone == "" || password == "" {
		ctx.String(http.StatusForbidden, ErrorMsg)
		return
	}

	// sql := `CREATE TABLE "users" (
	// 	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	// 	"phone" TEXT,
	// 	"password" TEXT,
	// 	"autokey" TEXT,
	// 	"ticket" TEXT,
	// 	"userid" INTEGER,
	// 	"last_login_time" INTEGER
	//   );
	//   CREATE TABLE "user_key" (
	// 	"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	// 	"userid" INTEGER,
	// 	"key" INTEGER
	//   );`
	stmt, err := UserEng.DB().Prepare("SELECT password,autokey,ticket,userid FROM users WHERE phone = ?")
	CheckErr(err)
	defer stmt.Close()

	var pass, autokey, ticket, userid string
	rows, err := stmt.Query(phone)
	CheckErr(err)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&pass, &autokey, &ticket, &userid)
		CheckErr(err)
	}

	loginResp := LoginResp{}
	loginCode := 0
	loginMsg := "ok"
	loginTime := time.Now().Unix()
	if pass == "" {
		// 未注册 - 自动注册
		session := UserEng.NewSession()
		defer session.Close()

		if err = session.Begin(); err != nil {
			session.Rollback()
			panic(err)
		}

		pass = openssl.Md5ToString(password)
		autokey = "AUTO" + strings.ToUpper(utils.RandomStr(32))
		userid = strconv.Itoa(int(loginTime))
		ticket = "9999999" + userid + userid
		userStmt, err := session.DB().Prepare("INSERT INTO users(phone,password,autokey,ticket,userid,last_login_time) VALUES (?,?,?,?,?,?)")
		if err != nil {
			session.Rollback()
			panic(err)
		}
		defer userStmt.Close()

		_, err = userStmt.Exec(phone, pass, autokey, ticket, userid, loginTime)
		if err != nil {
			session.Rollback()
			panic(err)
		}
		// id, _ := res.LastInsertId()
		// fmt.Println("LastInsertId:", id)

		keyStmt, err := session.DB().Prepare("INSERT INTO user_key(userid,key) VALUES(?,?)")
		if err != nil {
			session.Rollback()
			panic(err)
		}
		// 方便起见初始化 userid 和 key 一样
		// 注意：user_key 表中的 key 是上文生成的用于登录的 userid，而 userid 则是用于 Authorize Token 生成用的
		_, err = keyStmt.Exec(userid, userid)
		if err != nil {
			session.Rollback()
			panic(err)
		}

		if err = session.Commit(); err != nil {
			session.Rollback()
			panic(err)
		}

		tools.InitUserData(int(loginTime))

		// Login Response
		loginResp.Autokey = autokey
		loginResp.HasRealInfo = 1
		loginResp.Message = "ok"
		loginResp.RealInfoForce = 1
		loginResp.Ticket = ticket
		loginResp.UserAttribute = "0"
		loginResp.Userid = userid
	} else {
		// 已注册
		if openssl.Md5ToString(password) == pass {
			// 密码正确
			// Login Response
			loginResp.Autokey = autokey // 注意：更换设备（deviceId 发生变化）应重新生成 autokey
			loginResp.HasRealInfo = 1
			loginResp.Message = "ok"
			loginResp.RealInfoForce = 1
			loginResp.Ticket = "9999999" + userid + strconv.Itoa(int(loginTime)) // 实际登录用的密码（每次登录都会重新生成新的）
			loginResp.UserAttribute = "0"
			loginResp.Userid = userid // 实际登录用的账号

			// 更新信息
			userStmt, err := UserEng.DB().Prepare("UPDATE users SET autokey=?,ticket=?,last_login_time=? WHERE userid=?")
			CheckErr(err)
			defer userStmt.Close()

			_, err = userStmt.Exec(autokey, ticket, loginTime, userid)
			CheckErr(err)
			// aff, _ := res.RowsAffected()
			// fmt.Println("RowsAffected:", aff)
		} else {
			// 密码错误
			loginCode = 31
			loginMsg = "账号不存在或者密码有误！"
		}
	}

	data, err := json.Marshal(loginResp)
	CheckErr(err)
	// fmt.Println(string(data))
	encryptedData, err := openssl.Des3ECBEncrypt(data, randKey[0:24], openssl.PKCS7_PADDING)
	CheckErr(err)
	encryptedData64 := base64.StdEncoding.EncodeToString(encryptedData)
	// fmt.Println(encryptedToken64)

	resp := fmt.Sprintf(`{ "code": %d, "msg": %s, "data": "%s" }`, loginCode, loginMsg, encryptedData64)

	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, resp)
}

func ReportRoleHandler(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// CheckErr(err)
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))

	// body64, err := base64.StdEncoding.DecodeString(string(body))
	// CheckErr(err)
	// fmt.Println(string(body64))

	deviceId := ctx.Request.Header.Get("X-DEVICEID")
	randKey, err := database.LevelDb.Get([]byte(deviceId))
	CheckErr(err)

	// decryptedBody, err := openssl.Des3ECBDecrypt(body64, randKey[0:24], openssl.PKCS7_PADDING)
	// CheckErr(err)
	// fmt.Println(string(decryptedBody))

	// decrypted, err := url.QueryUnescape(string(decryptedBody))
	// CheckErr(err)
	// fmt.Println(decrypted)

	// Unable to decrypt server data
	token := `{"message":"Hello, world!"}`
	encryptedToken, err := openssl.Des3ECBEncrypt([]byte(token), randKey[0:24], openssl.PKCS7_PADDING)
	CheckErr(err)
	encryptedToken64 := base64.StdEncoding.EncodeToString(encryptedToken)
	// fmt.Println(encryptedToken64)

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
	// CheckErr(err)
	// defer ctx.Request.Body.Close()
	// fmt.Println(string(body))
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	ctx.String(http.StatusOK, "")
}

func ReportApp(ctx *gin.Context) {
	// body, err := io.ReadAll(ctx.Request.Body)
	// CheckErr(err)
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

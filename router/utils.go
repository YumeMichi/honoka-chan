package router

import (
	"fmt"
	"honoka-chan/config"
	"honoka-chan/encrypt"
	"time"
)

func SignResp(ep, body, key string) (resp string) {
	signBody := fmt.Sprintf("%d,\"%s\",0,%s", time.Now().UnixMilli(), config.MasterVersion, body)
	sign := encrypt.HMAC_SHA1_Encrypt([]byte(ep+" "+signBody), []byte(key))
	// fmt.Println(sign)

	resp = fmt.Sprintf("[%s,\"%s\"]", signBody, sign)
	return
}

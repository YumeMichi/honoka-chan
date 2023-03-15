package encrypt

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
)

func HMAC_SHA1_Encrypt(plainText []byte, key []byte) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(plainText)
	cipherText := fmt.Sprintf("%x", mac.Sum(nil))
	return cipherText
}

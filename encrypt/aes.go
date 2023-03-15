package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

var (
	iv = []byte("12345678abcdefgh")
)

func Padding(plainText []byte, blockSize int) []byte {
	n := blockSize - len(plainText)%blockSize
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

func UnPadding(cipherText []byte) []byte {
	end := cipherText[len(cipherText)-1]
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}

func AES_CBC_Encrypt(plainText []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	plainText = Padding(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText
}

func AES_CBC_Decrypt(cipherText []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	plainText := make([]byte, len(cipherText))
	blockMode := cipher.NewCBCDecrypter(block, iv)
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = UnPadding(plainText)
	return plainText
}

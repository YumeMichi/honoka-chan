package encrypt

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"runtime"
)

func RSA_Gen(bits int) {
	//get current path
	_, currentpath, _, _ := runtime.Caller(0)
	currentpath = filepath.Dir(currentpath)

	//----------------------------------------------private key

	// GenerateKey generates an RSA keypair of the given bit size using the
	// random source random (for example, crypto/rand.Reader).
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	//serialize privatekey to ASN.1 der by x509.MarshalPKCS8PrivateKey
	x509privatekey, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		panic(err)
	}

	//encode x509 to pem and save to file
	//1. create privatefile
	privatekeyfile, err := os.Create(currentpath + "/privatekey.pem")
	if err != nil {
		panic(err)
	}
	defer privatekeyfile.Close()
	//2. new a pem block struct object
	privatekeyblock := pem.Block{
		Type:    "PRIVATE KEY",
		Headers: nil,
		Bytes:   x509privatekey,
	}
	//3. save to file
	pem.Encode(privatekeyfile, &privatekeyblock)

	//----------------------------------------------public key

	//get public key
	publickey := privateKey.PublicKey
	//serialize publickey to ASN.1 der by x509.MarshalPKCS8PublicKey
	x509publickey, _ := x509.MarshalPKIXPublicKey(&publickey)

	//encode x509 to pem and save to file
	//1. create publickeyfile
	publickeyfile, err := os.Create(currentpath + "/publickey.pem")
	if err != nil {
		panic(err)
	}
	defer publickeyfile.Close()

	//2. new a pem block struct object
	publickeyblock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   x509publickey,
	}

	//3. save to file
	pem.Encode(publickeyfile, &publickeyblock)
}
func RSA_Encrypt(plainText []byte, publickeypath string) []byte {
	//open publickeyfile
	publickeyfile, err := os.Open(publickeypath)
	if err != nil {
		panic(err)
	}
	defer publickeyfile.Close()

	//get publickeyfile info
	publickeyfileInfo, _ := publickeyfile.Stat()

	//read publickeyfile content
	//1. make size
	buf := make([]byte, publickeyfileInfo.Size())
	//2. read file to buf
	publickeyfile.Read(buf)
	//3. decode pem
	publickeyDecodeBlock, _ := pem.Decode(buf)
	//4. x509 decode
	publicKeyInterface, err := x509.ParsePKIXPublicKey(publickeyDecodeBlock.Bytes)
	if err != nil {
		panic(err)
	}
	//assert
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//encrypt plainText
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		panic(err)
	}

	return cipherText
}

func RSA_Decrypt(cipherText []byte, privatekeypath string) []byte {
	//open privatekeyfile
	privatekeyfile, err := os.Open(privatekeypath)
	if err != nil {
		panic(err)
	}
	defer privatekeyfile.Close()
	//get privatekeyfile content
	privatekeyinfo, _ := privatekeyfile.Stat()
	buf := make([]byte, privatekeyinfo.Size())
	privatekeyfile.Read(buf)
	//pem decode
	privatekeyblock, _ := pem.Decode(buf)
	//X509 decode
	parseKey, err := x509.ParsePKCS8PrivateKey(privatekeyblock.Bytes)
	if err != nil {
		panic(err)
	}
	privateKey := parseKey.(*rsa.PrivateKey)
	//decrypt the cipher
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		panic(err)
	}

	return plainText
}

func RSA_Sign_SHA1(cipherText []byte, privatekeypath string) []byte {
	//open privatekeyfile
	privatekeyfile, err := os.Open(privatekeypath)
	if err != nil {
		panic(err)
	}
	defer privatekeyfile.Close()
	//get privatekeyfile content
	privatekeyinfo, _ := privatekeyfile.Stat()
	buf := make([]byte, privatekeyinfo.Size())
	privatekeyfile.Read(buf)
	//pem decode
	privatekeyblock, _ := pem.Decode(buf)
	//X509 decode
	parseKey, err := x509.ParsePKCS8PrivateKey(privatekeyblock.Bytes)
	if err != nil {
		panic(err)
	}
	privateKey := parseKey.(*rsa.PrivateKey)

	msgHash := sha1.New()
	_, err = msgHash.Write(cipherText)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA1, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	return signature
}

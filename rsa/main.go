package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func mmain() {
	//rsa 密钥文件产生
	GenRsaKey(1024)
}

//RSA公钥私钥产生
func GenRsaKey(bits int) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func FileLoad(filepath string) []byte {

	privatefile, err := os.Open(filepath)

	defer privatefile.Close()

	if err != nil {

		return nil

	}
	privateKey := make([]byte, 2048)

	num, err := privatefile.Read(privateKey)

	return privateKey[:num]
}

//rsa加密解密 签名验签
var publickey = FileLoad("public.pem")
var privatekey = FileLoad("private.pem")

func RSAEncrypt(orgidata []byte) ([]byte, error) {
	block, _ := pem.Decode(publickey)
	if block == nil {
		return nil, errors.New("public key is bad")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, orgidata) //加密
}

func RSADecrypt(cipertext []byte) ([]byte, error) {
	block, _ := pem.Decode(privatekey)
	if block == nil {
		return nil, errors.New("public key is bad")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, cipertext)
}

func main() {
	var data []byte
	var err error
	// nw/aWrN8bNf4gWgr6OziAetD2QeB8yrrJc/No5MRblUts7LjAlHR2KqmVXx+OnqBj+733jC+ISC4RCdl0JM8EK+/Z5CtC+0qT7jOiUN3qJpDB4KbeC36MSIUPMeYnqwSTyFCjtba2LCg2ZbB9f/Xler7EfoIWSfH4BgvEtNjPPk=
	data, err = RSAEncrypt([]byte("dfkjaeuijaf"))
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println("加密：", base64.StdEncoding.EncodeToString(data))
	bt, err := base64.StdEncoding.DecodeString(base64.StdEncoding.EncodeToString(data))
	if err != nil {
		fmt.Println(err.Error())
	}
	origData, err := RSADecrypt(bt) //解密
	if err != nil {
		fmt.Println("错误", err)
	}
	fmt.Println("解密:", string(origData)) //
	pk := FileLoad("myprivatekey.pem")
	fmt.Println(string(pk))
}

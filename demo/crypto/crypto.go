package crypto

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

// CryptoMain 获取唯一相应的密码值
func CryptoMain() {
	dk, _ := scrypt.Key([]byte("456789"), []byte("shentongtong"), 32768, 8, 1, 32)
	println(dk)
}

// Encrypt base64 加密
func Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

// Decode base64 解密
func Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

// Base64Main Base64加解密
func Base64Main() {

	// encode
	hello := "你好世界，Hello,world!"
	enbyte := Encode([]byte(hello))
	fmt.Println("Encoded:", string(enbyte))

	// decode
	decode, err := Decode([]byte(enbyte))
	if err != nil {
		fmt.Println("Decode error:", err.Error())
	}

	fmt.Println(hello)

	fmt.Println(string(decode))

	if hello != string(decode) {
		fmt.Println("encode not equal decode!")
	}

}

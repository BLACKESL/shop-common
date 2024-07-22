package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempstr := h.Sum(nil)
	return hex.EncodeToString(tempstr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakePassword(password, salt string) string {
	return Md5Encode(password + salt)
}

// 解密
func ValidPassword(plainpwd, salt, password string) bool {
	md := Md5Encode(plainpwd + salt)
	fmt.Println(md + "				" + password)
	return md == password

}

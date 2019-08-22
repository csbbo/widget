package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

//先sha256加密再截取前中后三段进行md5加密
func Encrypt(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	sha256_str := hex.EncodeToString(h.Sum(nil)) //sha256后得到64位的字符串

	length := len(sha256_str)/8
	slice_str := sha256_str[:length]+sha256_str[4*length:5*length]+sha256_str[7*length:]

	m := md5.New()
	m.Write([]byte(slice_str))
	return hex.EncodeToString(m.Sum(nil)) //md5后得到32位的字符串
}

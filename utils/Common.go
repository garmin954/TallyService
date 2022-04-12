package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
)

// 随机字符串
func RandomStr(n int) string {
	var l = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	length := len(l)
	for i := range b {
		b[i] = l[rand.Intn(length)]
	}
	return string(b)
}

// 返回一个32位md5加密后的字符串
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func JsonToMap(str string) map[string]interface{} {

	var tempMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &tempMap)

	if err != nil {
		panic(err)
	}

	return tempMap
}

func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

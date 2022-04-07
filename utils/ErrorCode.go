package utils

var ERROR_TOKEN_EXIST int = 4100
var AUTHORIZATION_HAS_EXPIRED = 4104

func GetErrMsg(code int) string {
	errorMsgs := map[int]string{
		ERROR_TOKEN_EXIST:         "token不存在",
		AUTHORIZATION_HAS_EXPIRED: "token授权已过期,请重新登录",
	}

	msg := errorMsgs[code]
	return msg
}

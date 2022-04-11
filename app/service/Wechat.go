package service

import (
	"encoding/json"
	"fmt"

	"github.com/idoubi/goz"
)

type Wechat struct{}

type JsCodeSession struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// 获取xopenid
func (wechat *Wechat) GetCode(code string) (rep *JsCodeSession, err error) {
	cli := goz.NewClient()

	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		"wx8ee1d2607ce3dd67",
		"5c7b99bbae37070f81c2ac2b1f1dac22",
		code,
	)

	resp, err := cli.Get(url)
	if err != nil {
		return nil, err
	}
	body, _ := resp.GetBody()
	json.Unmarshal(body, &rep)

	return rep, nil
}

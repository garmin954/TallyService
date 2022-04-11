package service

import (
	"fmt"
	"io/ioutil"

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

	resp, err := cli.Get(url, goz.Options{JSON: struct {
		Key1 string   `json:"key1"`
		Key2 []string `json:"key2"`
		Key3 int      `json:"key3"`
	}{"value1", []string{"value21", "value22"}, 333}})
	if err != nil {
		return nil, err
	}

	body, err := resp.GetParsedBody()
	r, _ := ioutil.ReadAll(resp.GetRequest().Body)
	fmt.Println(r)
	fmt.Println(body)

	if err != nil {
		return nil, err
	}

	// contents := utils.JsonToMap(body.GetContents())
	// rep.Openid contents
	return rep, nil
}

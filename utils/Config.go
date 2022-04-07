package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Database DatabaseConfig `json:"database"`
	Sms      SmsConfig      `json:"sms"`
	Jwt      JwtConfig      `json:"jwt"`
}

type SmsConfig struct {
	SignName        string `json:"sign_name"`
	TemplateCode    string `json:"template_code"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

type JwtConfig struct {
	Key    string `json:"key"`
	Expire int64  `json:"expire"`
}

var Configs *Config = nil

func init() {

	_, err := ParesConfig("./config/app.json")
	fmt.Println("-----------smg---------------------", Configs.Jwt)
	if err != nil {
		panic(err.Error())
	}
}

func GetConfig() *Config {
	return Configs
}

func ParesConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoer := json.NewDecoder(reader)

	if err = decoer.Decode(&Configs); err != nil {
		return nil, err
	}
	return Configs, nil
}

package utils

import (
	"bufio"
	"encoding/json"
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

var _cfg *Config = nil

func GetConfig() *Config {
	return _cfg
}

func ParesConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoer := json.NewDecoder(reader)

	if err = decoer.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}

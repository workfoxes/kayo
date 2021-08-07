package config

import (
	"encoding/json"
	"github.com/workfoxes/kayo/pkg/utils"
	"io/ioutil"
)

var C *Config

type Config struct {
	TraderKey          string `json:"TraderKey"`
	Env                string `json:"ENV"`
	GoogleClientId     string `json:"GoogleClientId"`
	GoogleClientSecret string `json:"GoogleClientSecret"`
	HOST               string `json:"Host"`
	RedisHost          string `json:"RedisHost"`
	RedisPassword      string `json:"RedisPassword"`
}

func init() {
	C = GetConfig()
}

var FileName = "config.json"

func GetConfig() *Config {
	env := utils.GetEnv("ENV", "Dev")
	var config Config
	data, _ := ioutil.ReadFile(FileName)
	_ = json.Unmarshal(data, &config)
	config.Env = env
	return &config
}

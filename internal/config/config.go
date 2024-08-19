package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type configuration struct {
	PORT           string
	BOT_TOKEN      string
	SLACK_TOKEN    string
	Linear_api_key string
	DATABASE_URL   string
}

func LoadConfig() (*configuration, error) {
	var config *configuration

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading env \n err:%v", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshal env \n err : %v", err)
	}

	return config, nil

}

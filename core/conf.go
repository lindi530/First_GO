package core

import (
	"GO1/config"
	"GO1/global"
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", viper.ConfigFileUsed())
		return
	}

	cfg := &config.Config{}
	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Println("Error unmarshalling config file:", viper.ConfigFileUsed())
		return
	}
	global.Config = cfg
}

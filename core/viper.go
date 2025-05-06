package core

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("settings")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file:", viper.ConfigFileUsed())
		return
	}
}

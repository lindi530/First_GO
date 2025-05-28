package conf

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
	//allSettings := viper.AllSettings()
	//fmt.Printf("All config: %+v\n", allSettings)

	cfg := &config.Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Println("Error unmarshalling config file:", viper.ConfigFileUsed())
		return
	}
	//fmt.Printf("Snowflake Config: %+v\n", cfg.Snowflake)

	global.Config = cfg
}

package config

import (
	"GO1/config/settings"
	"GO1/config/upload"
)

type Config struct {
	Mysql       Mysql             `mapstructure:"mysql"`
	Logger      Logger            `mapstructure:"logger"`
	System      System            `mapstructure:"system"`
	Snowflake   Snowflake         `mapstructure:"snowflake"`
	RedisClient RedisClient       `mapstructure:"redis"`
	JWT         JWT               `mapstructure:"jwt"`
	Upload      upload.Upload     `mapstructure:"upload"`
	Settings    settings.Settings `mapstructure:"settings"`
}

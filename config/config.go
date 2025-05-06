package config

type Config struct {
	Mysql  Mysql  `mapstructure:"mysql"`
	Logger Logger `mapstructure:"logger"`
	System System `mapstructure:"system"`
}

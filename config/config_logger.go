package config

type Logger struct {
	FilenameHttp  string `mapstructure:"fileNameHttp"`
	FilenameOther string `mapstructure:"fileNameOther"`
	AppName       string `mapstructure:"appName"`
	MaxSize       int    `mapstructure:"maxSize"`
	MaxBackups    int    `mapstructure:"maxBackups"`
	MaxAge        int    `mapstructure:"maxAge"`
	Compress      bool   `mapstructure:"compress"`
}

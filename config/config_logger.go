package config

type Logger struct {
	Filename      string `mapstructure:"fileName"`
	FilenameError string `mapstructure:"fileNameError"`
	AppName       string `mapstructure:"appName"`
	MaxSize       int    `mapstructure:"maxSize"`
	MaxBackups    int    `mapstructure:"maxBackups"`
	MaxAge        int    `mapstructure:"maxAge"`
	Compress      bool   `mapstructure:"compress"`
}

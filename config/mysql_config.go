package config

import (
	"os"
	"strconv"
)

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	DB       string `mapstructure:"db"`
	User     string `mapstructure:"user"`
	Config   string `mapstructure:"config"` // 高级配置，例如charset
	Password string `mapstructure:"password"`
	LogLevel string `mapstructure:"log_level"`
}

func (m *Mysql) DSN() string {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?charset=utf8mb4&parseTime=True&loc=Local"
	}
	return dsn
}

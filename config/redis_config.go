package config

type RedisClient struct {
	Address     string `mapstructure:"address"`
	KeyPrefix   string `mapstructure:"key_prefix"`
	MaxIdle     int    `mapstructure:"max_idle"`
	IdleTimeout int    `mapstructure:"idle_timeout"`
}

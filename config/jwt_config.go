package config

type JWT struct {
	AccessDuration  int `mapstructure:"access_duration"`
	RefreshDuration int `mapstructure:"refresh_duration"`
}

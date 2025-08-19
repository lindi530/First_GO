package config

import "fmt"

type System struct {
	Host   string `mapstructure:"host"`
	Scheme string `mapstructure:"scheme"`
	Port   int    `mapstructure:"port"`

	Env string `mapstructure:"env"`
}

func (s *System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

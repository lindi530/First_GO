package config

import "GO1/config/run_code"

type RunCodePath struct {
	Local   run_code.Local   `mapstructure:"local"`
	Service run_code.Service `mapstructure:"service"`
}

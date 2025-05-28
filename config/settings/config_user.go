package settings

type User struct {
	Avatar string `mapstructure:"avatar"`
	Quote  string `mapstructure:"quote"`
	Gender string `mapstructure:"gender"`
}

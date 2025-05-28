package upload

type Image struct {
	Path string `mapstructure:"path"`
	Size int64  `mapstructure:"size"`
}

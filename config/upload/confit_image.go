package upload

type Image struct {
	Path string `yaml: "path"`
	Size int64  `yaml: "size"`
}

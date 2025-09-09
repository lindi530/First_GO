package problem_service

import "GO1/global"

func getRunPath(dir, pattern *string) {
	var baseDir, prefix string
	switch env := global.Config.System.Env; env {
	case "debug":
		baseDir = global.Config.RunCode.Service.BaseDir
		prefix = global.Config.RunCode.Service.BaseDir
	case "release":
		baseDir = global.Config.RunCode.Local.BaseDir
		prefix = global.Config.RunCode.Local.Prefix
	}
	dir = &baseDir
	pattern = &prefix
}

package problem_api

import "strings"

// 代码安全检查
func isSafeCode(code, language string) bool {
	dangerMap := map[string][]string{
		"c":      {"system", "exec", "fork", "popen"},
		"cpp":    {"system", "exec", "fork", "popen"},
		"python": {"os.system", "subprocess", "eval", "exec", "shutil"},
		"java":   {"Runtime.getRuntime", "ProcessBuilder", "FilePermission"},
		"go":     {"os/exec", "syscall", "ioutil.WriteFile"},
	}

	if banned, ok := dangerMap[language]; ok {
		for _, keyword := range banned {
			if strings.Contains(code, keyword) {
				return false
			}
		}
	}
	return true
}

package problem_api

import (
	"GO1/middlewares/response"
	"GO1/models/problem_model"
	"GO1/pkg/jwt"
	"GO1/service/problem_service"
	"github.com/gin-gonic/gin"
	"strings"
)

func (ProblemAPI) SubmitCode(c *gin.Context) {
	var codeSubmission problem_model.CodeSubmission
	if err := c.ShouldBindJSON(&codeSubmission); err != nil {
		response.FailWithMessage("解析信息失败", c)
		return
	}
	if ok := isCodeSafe(codeSubmission.Code, codeSubmission.Language); !ok {
		response.FailWithMessage("不安全代码", c)
		return
	}
	userid := jwt.GetUserIdFromToken(c.GetHeader("Authorization"))

	resp := problem_service.SubmitCode(userid, codeSubmission)

	if resp.Code == 1 {
		response.FailWithMessage(resp.Message, c)
		return
	}

	response.OkWithData(resp.Data, c)
}

// 代码安全检查
func isCodeSafe(code, language string) bool {
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

package problem_service

import (
	"GO1/models/problem_model"
	"GO1/models/ws_model"
	"GO1/service/ws_service"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RunCode(userid int64, code, language string, testcases *[]problem_model.Example, message *ws_model.EditStatus) []problem_model.RunResult {
	dir := ""
	if tempDirEnv := os.Getenv("TEMP_DIR"); tempDirEnv != "" {
		dir = tempDirEnv
	}
	tempDir, err := os.MkdirTemp(dir, "ojcode-*")
	if err != nil {
		return []problem_model.RunResult{{Passed: false, Error: "Cannot create temp directory"}}
	}
	defer os.RemoveAll(tempDir)

	var codeFileName, compileCmd, runCmd, image string

	switch language {
	case "cpp":
		codeFileName = "code.cpp"
		compileCmd = "g++ /app/code.cpp -o /app/a.out"
		runCmd = "/app/a.out"
		image = "gcc:15"
	case "python":
		codeFileName = "code.py"
		runCmd = "python3 /app/code.py"
		image = "python:3.12"
	case "java":
		codeFileName = "Main.java"
		compileCmd = "javac /app/Main.java"
		runCmd = "java -cp /app Main"
		image = "openjdk:latest"
	case "go":
		codeFileName = "code.go"
		compileCmd = "go build -o /app/a.out /app/code.go"
		runCmd = "/app/a.out"
		image = "golang:latest"
	default:
		return []problem_model.RunResult{{Passed: false, Error: "Unsupported language"}}
	}

	// 写入代码文件
	codePath := filepath.Join(tempDir, codeFileName)
	err = os.WriteFile(codePath, []byte(code), 0644)
	if err != nil {
		return []problem_model.RunResult{{Passed: false, Error: "Cannot write code file"}}
	}

	// 写入所有样例的输入文件
	for i, tc := range *testcases {
		inputPath := filepath.Join(tempDir, fmt.Sprintf("input%d.txt", i))
		err := os.WriteFile(inputPath, []byte(tc.Input), 0644)
		if err != nil {
			return []problem_model.RunResult{{Passed: false, Error: fmt.Sprintf("Cannot write input %d", i)}}
		}
	}

	// 写入运行脚本 run.sh
	runScript := "#!/bin/sh\nset -e\n"
	for i := range *testcases {
		runScript += fmt.Sprintf("%s < /app/input%d.txt > /app/output%d.txt\n", runCmd, i, i)
	}
	runScriptPath := filepath.Join(tempDir, "run.sh")
	err = os.WriteFile(runScriptPath, []byte(runScript), 0755)
	if err != nil {
		return []problem_model.RunResult{{Passed: false, Error: "Cannot write run.sh"}}
	}

	// 构建 Docker 命令
	var dockerCmd []string
	if compileCmd != "" {
		dockerCmd = []string{
			"run", "--rm",
			"-v", fmt.Sprintf("%s:/app", tempDir),
			image,
			"sh", "-c",
			fmt.Sprintf("%s && sh /app/run.sh", compileCmd),
		}
	} else {
		dockerCmd = []string{
			"run", "--rm",
			"-v", fmt.Sprintf("%s:/app", tempDir),
			image,
			"sh", "-c",
			"sh /app/run.sh",
		}
	}
	ws_service.WsHub.SendEditData(message, "Running")
	// 执行 Docker 命令
	cmd := exec.Command("docker", dockerCmd...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		return []problem_model.RunResult{{Passed: false, Error: stderr.String()}}
	}

	// 读取每个样例的输出并比较
	results := make([]problem_model.RunResult, len(*testcases))
	for i, tc := range *testcases {
		outPath := filepath.Join(tempDir, fmt.Sprintf("output%d.txt", i))
		outBytes, err := os.ReadFile(outPath)
		if err != nil {
			results[i] = problem_model.RunResult{
				Passed: false,
				Error:  fmt.Sprintf("Cannot read output%d.txt: %s", i, err.Error()),
			}
			continue
		}
		output := strings.TrimSpace(string(outBytes))
		expected := strings.TrimSpace(tc.Output)
		results[i] = problem_model.RunResult{
			Passed:   output == expected,
			Output:   output,
			Expected: expected,
			Error:    "Wrong Answer", // stderr 已清空
		}
	}

	return results
}

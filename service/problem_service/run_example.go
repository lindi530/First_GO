package problem_service

import (
	"GO1/models/problem_model"
	"GO1/models/ws_model"
	"GO1/service/ws_service"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func RunExample(code, language, input string, memoryLimit, timeLimit int, message *ws_model.MessageWs) problem_model.RunResult {
	ws_service.WsHub.CodeStateWs(message, "Pending")
	tempDir, err := os.MkdirTemp("", "ojcode-*")
	if err != nil {
		return problem_model.RunResult{Passed: false, Error: "Cannot create temp directory"}
	}
	defer os.RemoveAll(tempDir)

	var codeFileName, compileCmd, runCmd, image string

	switch language {
	case "cpp":
		codeFileName = "code.cpp"
		compileCmd = "g++ /app/code.cpp -o /app/a.out"
		runCmd = "/app/a.out < /app/input.txt > /app/output.txt"
		image = "gcc:15"
	case "python":
		codeFileName = "code.py"
		compileCmd = "" // 无需编译
		runCmd = "python3 /app/code.py < /app/input.txt > /app/output.txt"
		image = "python:3.12"
	case "java":
		codeFileName = "Main.java"
		compileCmd = "javac /app/Main.java"
		runCmd = "java -cp /app Main < /app/input.txt > /app/output.txt"
		image = "openjdk:latest"
	case "go":
		codeFileName = "code.go"
		compileCmd = "go build -o /app/a.out /app/code.go"
		runCmd = "/app/a.out < /app/input.txt > /app/output.txt"
		image = "golang:latest"
	default:
		return problem_model.RunResult{Passed: false, Error: "Unsupported language"}
	}

	codePath := filepath.Join(tempDir, codeFileName)
	inputPath := filepath.Join(tempDir, "input.txt")
	outputPath := filepath.Join(tempDir, "output.txt")

	// 写入代码和输入
	err = os.WriteFile(codePath, []byte(code), 0644)
	if err != nil {
		return problem_model.RunResult{Passed: false, Error: "Cannot write code file"}
	}

	err = os.WriteFile(inputPath, []byte(input), 0644)
	if err != nil {
		return problem_model.RunResult{Passed: false, Error: "Cannot write input file"}
	}

	// 构建 docker 命令

	dockerCmd := buildDockerCmd(image, tempDir, compileCmd, runCmd, memoryLimit, timeLimit)
	ws_service.WsHub.CodeStateWs(message, "Running")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeLimit)*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "docker", dockerCmd...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err = cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return problem_model.RunResult{Passed: false, Error: "Time Limit Exceeded"}
	}

	if err != nil {
		errStr := stderr.String()
		// 你也可以在这里加上额外的错误码判断，比如 exit status 137（内存爆）
		if strings.Contains(errStr, "exit status 137") {
			return problem_model.RunResult{Passed: false, Error: "Memory Limit Exceeded"}
		}
		return problem_model.RunResult{Passed: false, Error: errStr}
	}

	outBytes, err := os.ReadFile(outputPath)
	if err != nil {
		return problem_model.RunResult{Passed: false, Error: "Cannot read output"}
	}
	ws_service.WsHub.CodeStateWs(message, "Finished")
	outStr := strings.TrimSpace(string(outBytes))
	return problem_model.RunResult{
		Output: outStr,
	}
}

func buildDockerCmd(image, tempDir, compileCmd, runCmd string, memoryLimit, timeLimit int) []string {
	memLimit := fmt.Sprintf("--memory=%dm", memoryLimit)
	timeoutCmd := fmt.Sprintf("timeout %ds %s", timeLimit, runCmd)

	var fullCmd string
	if compileCmd != "" {
		fullCmd = fmt.Sprintf("%s && %s", compileCmd, timeoutCmd)
	} else {
		fullCmd = timeoutCmd
	}

	return []string{
		"run", "--rm",
		memLimit,
		"-v", fmt.Sprintf("%s:/app", tempDir),
		image,
		"sh", "-c",
		fullCmd,
	}
}

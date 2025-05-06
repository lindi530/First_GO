package middlewares

import (
	"GO1/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logFile := &lumberjack.Logger{
		Filename:   global.Config.Logger.FilenameHttp, // 路径建议 logs 目录存在
		MaxSize:    global.Config.Logger.MaxSize,      // 文件最大大小
		MaxBackups: global.Config.Logger.MaxBackups,   // 最多保留文件个数
		MaxAge:     global.Config.Logger.MaxAge,       // 保留时间天
		Compress:   global.Config.Logger.Compress,     // 是否压缩旧文件
	}
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		appName := global.Config.Logger.AppName
		method := c.Request.Method
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		path := c.Request.URL.Path
		timestamp := time.Now().Format(time.DateTime)

		methodColor := SetColor(status)
		statusColor := SetColor(status)
		resetColor := SetColor(0)

		colorLog := fmt.Sprintf("[%s] | [%s] | %s%-6s%s | %s%d%s | %13v | %15s | %s\n",
			appName,
			timestamp,
			methodColor, method, resetColor,
			statusColor, status, resetColor,
			latency,
			clientIP,
			path,
		)

		// 无色输出到文件
		plainLog := fmt.Sprintf("[%s] | [%s] | %-6s | %3d | %13v | %15s | %s\n",
			appName,
			timestamp,
			method,
			status,
			latency,
			clientIP,
			path,
		)
		// 写入
		fmt.Fprint(os.Stdout, colorLog)
		fmt.Fprint(logFile, plainLog)
	}
}

func SetColor(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[32m" // 绿色    2xx（成功）	         请求成功
	case code >= 300 && code < 400:
		return "\033[36m" // 青色    3xx（重定向）		需要进一步操作
	case code >= 400 && code < 500:
		return "\033[33m" // 黄色    4xx（客户端错误）     客户端出现问题
	case code >= 500:
		return "\033[31m" // 红色    5xx（服务器错误）  	服务器内部错误
	default:
		return "\033[0m" //  默认    1xx（信息性响应）	  	指令性通知
	}
}

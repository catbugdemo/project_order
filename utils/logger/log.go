package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	Log = logrus.New()
)

func init() {
	InitLogger()
}

func InitLogger() {
	// 日志格式化为JSON而不是默认的ASCII
	Log.SetFormatter(&logrus.JSONFormatter{})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	Log.SetOutput(os.Stdout)

	// 只记录严重或以上警告
	Log.SetLevel(logrus.WarnLevel)
}

// New 用来创建新的 log 日志
func New() *logrus.Logger {
	return Log
}

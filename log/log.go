package log

import (
	"bytes"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

func init() {
	InitLogger()
}

func InitLogger() {
	// 日志格式化为JSON而不是默认的ASCII
	log.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})

	// 输出stdout而不是默认的stderr，也可以是一个文件
	log.SetOutput(os.Stdout)

	/*	// 只记录严重或以上警告
		log.SetLevel(logrus.WarnLevel)*/
}

// 自定义 输出格式
type MyFormatter struct {
	logrus.Logger
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func Info(args ...interface{}) {
	log.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Error(args ...interface{}) {
	log.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

package stdlog

import (
	"log"
	"os"
)

var defaultLogger *Logger

func init() {
	defaultLogger = &Logger{
		logger: log.New(os.Stderr, defaultPrefix, log.LstdFlags),
		level:  LevelInfo,
		color:  true,
	}
}

func Debug(message string) {
	defaultLogger.Debug(message)
}
func Info(message string) {
	defaultLogger.Info(message)
}

func Warn(message string) {
	defaultLogger.Warn(message)
}

// 无论是什么日志级别都要输出
func Error(message string) {
	defaultLogger.Error(message)
}

// 无论是什么日志级别都要输出
func Panic(message string) {
	defaultLogger.Panic(message)
}

// 无论是什么日志级别都要输出
func Fatal(message string) {
	defaultLogger.Fatal(message)
}

func DebugF(template string, args ...interface{}) {
	defaultLogger.DebugF(template, args...)
}
func InfoF(template string, args ...interface{}) {
	defaultLogger.InfoF(template, args...)
}

func WarnF(template string, args ...interface{}) {
	defaultLogger.WarnF(template, args...)
}

// 无论是什么日志级别都要输出
func ErrorF(template string, args ...interface{}) {
	defaultLogger.ErrorF(template, args...)
}

// 无论是什么日志级别都要输出
func PanicF(template string, args ...interface{}) {
	defaultLogger.PanicF(template, args...)
}

// 无论是什么日志级别都要输出
func FatalF(template string, args ...interface{}) {
	defaultLogger.FatalF(template, args...)
}

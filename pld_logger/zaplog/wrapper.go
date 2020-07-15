package zaplog

import (
	"fmt"
	"go.uber.org/zap"
)

func Copy() {

}
func Debug(message string, fields ...zap.Field) {
	zapLogger.Debug(message, fields...)
}

func DebugF(template string, args ...interface{}) {
	zapLogger.Debug(fmt.Sprintf(template, args...))
}

func Info(message string, fields ...zap.Field) {
	zapLogger.Info(message, fields...)
}
func InfoF(template string, args ...interface{}) {
	zapLogger.Info(fmt.Sprintf(template, args...))
}

func Warn(message string, fields ...zap.Field) {
	zapLogger.Warn(message, fields...)
}

func WarnF(template string, args ...interface{}) {
	zapLogger.Warn(fmt.Sprintf(template, args...))
}

func Error(message string, fields ...zap.Field) {
	zapLogger.Error(message, fields...)
}

func ErrorF(template string, args ...interface{}) {
	zapLogger.Error(fmt.Sprintf(template, args...))
}

func DPanic(message string, fields ...zap.Field) {
	zapLogger.DPanic(message, fields...)
}

func DPanicF(template string, args ...interface{}) {
	zapLogger.DPanic(fmt.Sprintf(template, args...))
}
func Panic(message string, fields ...zap.Field) {
	zapLogger.Panic(message, fields...)
}

func PanicF(template string, args ...interface{}) {
	zapLogger.Panic(fmt.Sprintf(template, args...))
}
func Fatal(message string, fields ...zap.Field) {
	zapLogger.Fatal(message, fields...)
}

func FatalF(template string, args ...interface{}) {
	zapLogger.Fatal(fmt.Sprintf(template, args...))
}

package zaplog

import (
	"fmt"
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger(dev bool, logPrefix string) *Logger {
	zapCfg := getZapConfig(dev, logPrefix)
	// 因为我们做了一层包装，所以需要跳过一层caller
	// 否则，日志的caller位置，始终显示的是当前logger包中的位置
	callerOption := zap.AddCallerSkip(2)
	var err error
	zapLogger, err = zapCfg.Build(callerOption)
	if err != nil {
		panic(err)
	}
	defer func() {
		zapLogger.Sync()
	}()
	return &Logger{logger: zapLogger}
}
func (l *Logger) getZapFields(values ...interface{}) (fields []zap.Field) {
	for idx, value := range values {
		if val, ok := value.(zap.Field); ok {
			fields = append(fields, val)
		} else {
			fields = append(fields, zap.Any(fmt.Sprintf("v_%d", idx), value))
		}
	}
	return
}
func (l *Logger) Debug(message string, values ...interface{}) {
	l.logger.Debug(message, l.getZapFields(values...)...)
}

func (l *Logger) Info(message string, values ...interface{}) {
	l.logger.Info(message, l.getZapFields(values...)...)
}

func (l *Logger) Warn(message string, values ...interface{}) {
	l.logger.Warn(message, l.getZapFields(values...)...)
}

func (l *Logger) Error(message string, values ...interface{}) {
	l.logger.Error(message, l.getZapFields(values...)...)
}

func (l *Logger) Panic(message string, values ...interface{}) {
	l.logger.Panic(message, l.getZapFields(values...)...)
}

func (l *Logger) Fatal(message string, values ...interface{}) {
	l.logger.Fatal(message, l.getZapFields(values...)...)
}

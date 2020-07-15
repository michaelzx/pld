package pld_logger

import (
	"github.com/michaelzx/pld/pld_logger/stdlog"
	"github.com/michaelzx/pld/pld_logger/zaplog"
	"io"
	"os"
)

type Logger interface {
	Debug(message string, values ...interface{})
	Info(message string, values ...interface{})
	Warn(message string, values ...interface{})
	Error(message string, values ...interface{})
	Panic(message string, values ...interface{})
	Fatal(message string, values ...interface{})
}

var (
	defaultLogger Logger
)

func init() {
	defaultLogger = stdlog.NewLogger(os.Stderr, stdlog.LevelDebug, true)
}

func UseStdLog(out io.Writer, level int, color bool) {
	defaultLogger = stdlog.NewLogger(out, level, color)
}
func UseZapLog(dev bool, logPrefix string) {
	defaultLogger = zaplog.NewLogger(dev, logPrefix)
}

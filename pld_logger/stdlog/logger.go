package stdlog

import (
	"github.com/fatih/color"
	"io"
	"log"
)

type Logger struct {
	logger *log.Logger
	level  int
	color  bool
}

func NewLogger(out io.Writer, level int, color bool) *Logger {
	return &Logger{logger: log.New(out, defaultPrefix, log.LstdFlags), level: level, color: color}
}

func (l *Logger) setPrefix(basicPrefix string) {
	if l.color {
		switch basicPrefix {
		case debugPrefix:
			l.logger.SetPrefix(color.GreenString(basicPrefix))
		case infoPrefix:
			l.logger.SetPrefix(color.BlueString(basicPrefix))
		case warnPrefix:
			l.logger.SetPrefix(color.YellowString(basicPrefix))
		case errorPrefix:
			l.logger.SetPrefix(color.RedString(basicPrefix))
		case panicPrefix:
			l.logger.SetPrefix(color.RedString(basicPrefix))
		case fatalPrefix:
			l.logger.SetPrefix(color.RedString(basicPrefix))
		}
	} else {
		l.logger.SetPrefix(basicPrefix)
	}
}
func (l *Logger) println(message string, values ...interface{}) {
	if len(values) == 0 {
		l.logger.Println(message)
	} else {
		l.logger.Printf("%s %+v", message, values)
	}
}
func (l *Logger) Debug(message string, values ...interface{}) {
	if l.level >= LevelDebug {
		l.setPrefix(debugPrefix)
		l.println(message, values...)
	}
}
func (l *Logger) Info(message string, values ...interface{}) {
	if l.level >= LevelInfo {
		l.setPrefix(infoPrefix)
		l.println(message, values...)
	}
}

func (l *Logger) Warn(message string, values ...interface{}) {
	if l.level >= LevelWarn {
		l.setPrefix(warnPrefix)
		l.println(message, values...)
	}
}

// 无论是什么日志级别都要输出
func (l *Logger) Error(message string, values ...interface{}) {
	if l.level >= LevelError {
		l.setPrefix(errorPrefix)
		l.println(message, values...)
	}
}

// 无论是什么日志级别都要输出
func (l *Logger) Panic(message string, values ...interface{}) {
	l.setPrefix(panicPrefix)
	if len(values) == 0 {
		l.logger.Panicln(message)
	} else {
		l.logger.Panicf("%s %v", message, values)
	}
}

// 无论是什么日志级别都要输出
func (l *Logger) Fatal(message string, values ...interface{}) {
	l.setPrefix(fatalPrefix)
	if len(values) == 0 {
		l.logger.Fatalln(message)
	} else {
		l.logger.Fatalf("%s %v", message, values)
	}
}

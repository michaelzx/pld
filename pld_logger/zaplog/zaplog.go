package zaplog

import (
	"github.com/michaelzx/pld/pld_fs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"path/filepath"
)

var (
	zapLogger *zap.Logger
	logDir    = "./app.log"
)

func GetZap() *zap.Logger {
	return zapLogger
}

func getZapConfig(devMode bool, logPrefix string) zap.Config {
	var loggingLevel zapcore.Level
	var OutputPaths []string
	var ErrorOutputPaths []string
	var levelEncoder zapcore.LevelEncoder
	if devMode {
		// 开发模式
		loggingLevel = zap.DebugLevel
		OutputPaths = []string{"stdout"}
		ErrorOutputPaths = []string{"stderr"}
		levelEncoder = zapcore.CapitalColorLevelEncoder
	} else {
		// 生产模式
		// TODO 对日志进行分割
		loggingLevel = zap.InfoLevel
		if logPrefix != "" {
			logDirPath := filepath.Join(pld_fs.PathAppRunning(), "log")
			log.Println(logDirPath)
			if !pld_fs.Exists(logDirPath) {
				if err := os.Mkdir(logDirPath, os.ModePerm); err != nil {
					log.Fatal("无法创建log目录", err.Error())
				}
			}
			OutputPaths = []string{"stdout", filepath.Join(logDirPath, logPrefix+".log")}
			ErrorOutputPaths = []string{"stderr", filepath.Join(logDirPath, logPrefix+".err.log")}
		} else {
			OutputPaths = []string{"stdout"}
			ErrorOutputPaths = []string{"stderr"}
		}
		levelEncoder = zapcore.CapitalLevelEncoder
		levelEncoder = zapcore.CapitalLevelEncoder
	}
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(loggingLevel),
		Development: devMode,
		Encoding:    "console",
		// 如果需要忽略输出，则赋值：zapcore.OmitKey
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "callerKey",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    levelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.NanosDurationEncoder,
			EncodeCaller:   customCallerEncoder,
		},
		OutputPaths:      OutputPaths,
		ErrorOutputPaths: ErrorOutputPaths,
	}
}
func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// TODO 考虑优化caller显示方式
	enc.AppendString(caller.TrimmedPath())
}

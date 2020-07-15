package pld_logger

func Debug(message string, values ...interface{}) {
	defaultLogger.Debug(message, values...)
}

func Info(message string, values ...interface{}) {
	defaultLogger.Info(message, values...)
}
func Warn(message string, values ...interface{}) {
	defaultLogger.Warn(message, values...)
}
func Error(message string, values ...interface{}) {
	defaultLogger.Error(message, values...)
}
func Panic(message string, values ...interface{}) {
	defaultLogger.Panic(message, values...)
}
func Fatal(message string, values ...interface{}) {
	defaultLogger.Fatal(message, values...)
}

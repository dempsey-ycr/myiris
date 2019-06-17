package logger

func Info(args ...interface{}) {
	_log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	_log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	_log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	_log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	_log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	_log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	_log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	_log.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	_log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	_log.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	_log.Debug(args...)
}
func Debugf(format string, args ...interface{}) {
	_log.Debugf(format, args...)
}

package log

import (
	"github.com/thep0y/go-logger/basic"
	"github.com/thep0y/go-logger/log"
)

var (
	logger 		*log.Logger
)

func InitLog(logLevel basic.LogLevel) {
	logger = log.NewLogger()
	logger.SetLogLevel(logLevel)
}

// Trace default Trace method
func Trace(v ...interface{}) {
	logger.Trace(v...)
}

// Tracef default Tracef method
func Tracef(format string, v ...interface{}) {
	logger.Tracef(format, v...)
}

// Info default Info method
func Info(v ...interface{}) {
	logger.Info(v...)
}

// Infof default Infof method
func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Debug default Debug method
func Debug(v ...interface{}) {
	logger.Debug(v...)
}

// Debugf default Debugf method
func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

// Warn default Warn method
func Warn(v ...interface{}) {
	logger.Warn(v...)
}

// Warnf default Warnf method
func Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Error default Error method
func Error(v ...interface{}) {
	logger.Error(v...)
}

// Errorf default Errorf method
func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

// Fatal default Fatal method
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Fatalf default Fatalf method
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

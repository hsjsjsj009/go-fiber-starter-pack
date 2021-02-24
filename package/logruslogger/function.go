package logruslogger

import log "github.com/sirupsen/logrus"

func Debug(message string, context string, scope string, corr ...interface{}) {
	Log(log.DebugLevel,message,context,scope,corr...)
}

func Info(message string, context string, scope string, corr ...interface{}) {
	Log(log.InfoLevel,message,context,scope,corr...)
}

func Warn(message string, context string, scope string, corr ...interface{}) {
	Log(log.WarnLevel,message,context,scope,corr...)
}

func Error(message string, context string, scope string, corr ...interface{}) {
	Log(log.ErrorLevel,message,context,scope,corr...)
}

func Fatal(message string, context string, scope string, corr ...interface{}) {
	Log(log.FatalLevel,message,context,scope,corr...)
}

func Panic(message string, context string, scope string, corr ...interface{}) {
	Log(log.PanicLevel,message,context,scope,corr...)
}

package log

import (
	"github.com/gookit/color"
	"spm/core/util"
	"strings"
)

const (
	LogDebug int = iota + 1
	LogInfo
	LogWarn
	LogError
)

const (
	LogDebugString = "debug"
	LogInfoString = "info"
	LogWarnString = "warning"
	LogErrorString = "error"
)

//日志级别
var logLevel = LogWarn

func SetLogLevel(lv int){
	logLevel = lv
}

func SetLogLevelString(lv string) {
	lv = strings.ToLower(strings.TrimSpace(lv))
	switch lv {
	case LogDebugString:
		logLevel = LogDebug
	case LogInfoString:
		logLevel = LogInfo
	case LogWarnString:
		logLevel = LogWarn
	case LogErrorString:
		logLevel = LogError
	}
}

func Debug(msg ...interface{}){
	if logLevel <= LogDebug {
		color.Gray.Println(util.SlicePrepend(msg, "DEBUG | "))
	}
}

func Info(msg ...interface{}){
	if logLevel <= LogInfo {
		color.Info.Println(msg...)
	}
}

func Warning(msg ...interface{}){
	if logLevel <= LogWarn {
		color.Warn.Println(msg...)
	}
}

func Error(msg ...interface{}){
	if logLevel <= LogError {
		color.Danger.Println(msg...)
	}
}

func IsDebug() bool{
	return logLevel == LogDebug
}

func IsInfo() bool{
	return logLevel == LogInfo
}

func IsWarning() bool{
	return logLevel == LogWarn
}

func IsError() bool{
	return logLevel == LogError
}

package log

import (
	"fmt"
	"log"
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

func Debug(msg ...string){
	if logLevel <= LogDebug {
		log.SetPrefix("DEBUG | ")
		log.Println(msg)
	}
}

func Info(msg ...string){
	if logLevel <= LogInfo {
		log.SetPrefix("INFO  | ")
		log.Println(msg)
	}
}

func Warning(msg ...string){
	if logLevel <= LogWarn {
		//log.SetPrefix("WARN  | ")
		//log.Println(msg)
		fmt.Println(msg)
	}
}

func Error(msg ...string){
	if logLevel <= LogError {
		//log.SetPrefix("ERROR | ")
		//log.Println(msg)
		fmt.Println(msg)
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

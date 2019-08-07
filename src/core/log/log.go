package log

import "log"

const (
	LOG_DEBUG int = iota + 1
	LOG_INFO
	LOG_WARN
	LOG_ERROR
)

var logLevel int

func SetLogLevel(lv int){
	logLevel = lv
}


func Debug(msg ...string){
	if logLevel <= LOG_DEBUG {
		log.Println("DEBUG: ", msg)
	}
}

func Info(msg ...string){
	if logLevel <= LOG_INFO {
		log.Println("INFO: ", msg)
	}
}

func Warning(msg ...string){
	if logLevel <= LOG_WARN {
		log.Println("WARN: ", msg)
	}
}

func Error(msg ...string){
	if logLevel <= LOG_ERROR {
		log.Println("ERROR: ", msg)
	}
}

package commands

import "log"

type Command struct {
	Log *log.Logger
}

func (c Command) Info(msg ...string){
	c.Log.Println("INFO: ", msg)
}

func (c Command) Warning(msg ...string){
	c.Log.Println("WARN: ", msg)
}

func (c Command) Error(msg ...string){
	c.Log.Println("ERROR: ", msg)
}

type Commander interface {
	Description() string
	Run() error
}







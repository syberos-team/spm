package commands

import (
	"core/util"
	"flag"
	"log"
)

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
	RegisterFlags(flags *flag.FlagSet)
}

//pwd 当前所在目录
var pwd string

func init(){
	path, err := util.Pwd()
	if err!=nil {
		panic(err)
	}
	pwd = path
}







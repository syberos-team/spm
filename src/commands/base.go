package commands

import (
	"core/util"
	"flag"
)

const SpmJsonFilename = "spm.json"

//pwd 当前所在目录
var pwd string

func init(){
	path, err := util.Pwd()
	if err!=nil {
		panic(err)
	}
	pwd = path
}

type Command struct {
}


type Commander interface {
	Description() string
	Run() error
	RegisterFlags(flags *flag.FlagSet)
}








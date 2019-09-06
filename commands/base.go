package commands

import (
	"spm/core/util"
)

const SpmJsonFilename = "spm.json"

//pwd 当前所在目录
var pwd string

func InitPwd(){
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
	RegisterArgs(args ...string)
	ArgsDescription() []ArgsDescription
}

type ArgsDescription struct {
	Name string
	Description string
	Required bool
	IsArray bool
}








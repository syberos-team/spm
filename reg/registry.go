package reg

import (
	"flag"
	"os"
	"spm/commands"
	"spm/core/log"
	"spm/core/util"
)

type Registry struct {
	//存储指令的实现，key指令，value为实现
	commandImpl map[string]commands.Commander
	//存储指令是否执行，key指令，value为是否执行
	commandRun map[string]*bool
}


func (r *Registry) RegistryCommand(name string, commander commands.Commander){
	var value = false
	flag.BoolVar(&value, name, false, commander.Description())

	r.commandImpl[name] = commander
	r.commandRun[name] = &value
}

func (r *Registry) RunCommand(){
	if r.runCopy() {
		return
	}
	for name, run := range r.commandRun {
		if *run {
			cmd := r.commandImpl[name]
			cmd.RegisterArgs(os.Args[2:]...)
			err := cmd.Run()
			if err!=nil {
				log.Error(err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}
	}
	commands.Usage()
}

func (r *Registry) runCopy() bool{
	if len(os.Args) == 4 && os.Args[1]=="copy" {
		src := os.Args[2]
		dst := os.Args[3]
		_, err := util.CopyFile(src, dst)
		if err!=nil {
			log.Error(err.Error())
		}
		return true
	}
	return false
}

func NewRegistry() *Registry{
	return &Registry{
		commandImpl: map[string]commands.Commander{},
		commandRun: map[string]*bool{},
	}
}
package reg

import (
	"commands"
	"core/log"
	"flag"
	"os"
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
	for name, run := range r.commandRun {
		if *run {
			cmd := r.commandImpl[name]
			cmd.RegisterFlags(flag.NewFlagSet(os.Args[0], flag.ExitOnError))
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

func NewRegistry() *Registry{
	return &Registry{
		commandImpl: map[string]commands.Commander{},
		commandRun: map[string]*bool{},
	}
}
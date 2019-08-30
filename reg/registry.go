package reg

import (
	"github.com/gookit/gcli/v2"
	"spm/commands"
	"spm/core"
	"spm/core/conf"
	"spm/core/log"
	"spm/core/util"
)

type Registry struct {
	app *gcli.App
}


func (r *Registry) RegistryCommand(name string, commander commands.Commander){
	cmd := gcli.NewCommand(name, commander.Description(), nil)
	argsDescription := commander.ArgsDescription()
	if argsDescription!=nil {
		for _, desc := range argsDescription {
			cmd.AddArg(desc.Name, desc.Description, desc.Required, desc.IsArray)
		}
	}
	cmd.SetFunc(func(c *gcli.Command, args []string) error{
		commander.RegisterArgs(args...)
		return commander.Run()
	})
	r.app.Add(cmd)
}

func (r *Registry) RunCommand(){
	if r.runCopy() {
		return
	}
	r.app.Run()
}

func (r *Registry) runCopy() bool{
	args := r.app.OsArgs()
	if len(args) == 4 && args[1]=="copy" {
		src := args[2]
		dst := args[3]
		log.Debug("copy", src, dst)
		_, err := util.CopyFile(src, dst)
		if err!=nil {
			log.Error(err.Error())
		}
		return true
	}
	return false
}

func NewRegistry() *Registry{
	app := gcli.NewApp()
	app.Name = conf.FILENAME
	app.Version = core.VERSION
	app.Description = "spm is a tool for managing syberos app dependencies"
	return &Registry{
		app: app,
	}
}
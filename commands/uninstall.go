package commands

import (
	"spm/core/log"
	"strings"
)

type UninstallCommand struct {
	Command
	packageName string		//包名，通过参数获得
}

func (u *UninstallCommand) Description() string {
	return "Uninstalls a package"
}

func (u *UninstallCommand) Run() error {
	log.Error("uninstall command unrealized")
	return nil
}

func (u *UninstallCommand) RegisterArgs(args ...string) {
	u.packageName = strings.TrimSpace(args[0])
}

func (u *UninstallCommand) ArgsDescription() []ArgsDescription{
	return []ArgsDescription{
		{
			Name:        "package",
			Description: "Package name to be uninstalled",
			Required:    true,
			IsArray:     false,
		},
	}
}


func NewUninstallCommand() *UninstallCommand{
	return &UninstallCommand{}
}




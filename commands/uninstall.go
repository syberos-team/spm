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
	if args==nil || len(args)==0 {
		return
	}
	u.packageName = strings.TrimSpace(args[0])
}

func NewUninstallCommand() *UninstallCommand{
	return &UninstallCommand{}
}




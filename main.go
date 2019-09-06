package main

import (
	"spm/commands"
	"spm/core/conf"
	_ "spm/core/conf"
	"spm/reg"
)

var registry *reg.Registry

func init(){
	conf.InitConfig()
	commands.InitPwd()

	registry = reg.NewRegistry()

	registry.RegistryCommand("init", commands.NewInitCommand())
	registry.RegistryCommand("search", commands.NewSearchCommand())
	registry.RegistryCommand("info", commands.NewInfoCommand())
	registry.RegistryCommand("publish", commands.NewPublishCommand())
	registry.RegistryCommand("install", commands.NewInstallCommand())
	registry.RegistryCommand("uninstall", commands.NewUninstallCommand())
	registry.RegistryCommand("upgrade", commands.NewUpgradeCommand())
}


func main(){
	registry.RunCommand()
}

package main

import (
	"flag"
	"spm/commands"
	_ "spm/core/conf"
	"spm/reg"
)

var registry *reg.Registry

func init(){
	registry = reg.NewRegistry()

	registry.RegistryCommand("help", commands.NewHelpCommand())
	registry.RegistryCommand("init", commands.NewInitCommand())
	registry.RegistryCommand("search", commands.NewSearchCommand())
	registry.RegistryCommand("info", commands.NewInfoCommand())
	registry.RegistryCommand("publish", commands.NewPublishCommand())
	registry.RegistryCommand("install", commands.NewInstallCommand())
	registry.RegistryCommand("uninstall", commands.NewUninstallCommand())
	registry.RegistryCommand("version", commands.NewVersionCommand())
	registry.RegistryCommand("upgrade", commands.NewUpgradeCommand())
}


func main(){
	flag.Parse()
	registry.RunCommand()
}

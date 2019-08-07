package main

import (
	"commands"
	"flag"
	"reg"
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
}


func main(){
	flag.Parse()
	registry.RunCommand()
}

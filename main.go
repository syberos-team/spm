package main

import (
	"spm/commands"
	"spm/core/conf"
	_ "spm/core/conf"
	"spm/reg"
)

// spm 版本号，spm版本升级时在此处修改版本号
const VERSION = "1.0.0"

var registry *reg.Registry

func init(){
	conf.InitConfig(VERSION)
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

package commands

import (
	"core"
	"fmt"
)

type UpgradeCommand struct {
	Command
}

func (u *UpgradeCommand) Description() string {
	return "Upgrade spm"
}

func (u *UpgradeCommand) Run() error {
	versionManage := core.NewVersionManage()
	needUpgrade, err := versionManage.CheckVersion()
	if err!=nil {
		return err
	}
	if needUpgrade {
		return versionManage.Upgrade()
	}
	fmt.Println("version:", core.VERSION)
	fmt.Println("The current version is up to date")
	return nil
}

func (u *UpgradeCommand) RegisterArgs(args ...string) {
}

func NewUpgradeCommand() *UpgradeCommand {
	return &UpgradeCommand{}
}




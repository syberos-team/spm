package commands

import "core/log"

type UpgradeCommand struct {
	Command
}

func (u *UpgradeCommand) Description() string {
	return "Upgrade spm"
}

func (u *UpgradeCommand) Run() error {
	log.Error("upgrade command unrealized")
	return nil
}

func (u *UpgradeCommand) RegisterArgs(args ...string) {
}

func NewUpgradeCommand() *UpgradeCommand {
	return &UpgradeCommand{}
}




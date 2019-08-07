package commands

import (
	"core"
	"fmt"
)

type VersionCommand struct {
	Command
}

func (v VersionCommand) Description() string {
	return "Output product version"
}

func (v VersionCommand) Run() error {
	fmt.Println(core.VERSION)
	return nil
}

func (v VersionCommand) RegisterArgs(args ...string) {
}

func NewVersionCommand() *VersionCommand {
	return &VersionCommand{}
}



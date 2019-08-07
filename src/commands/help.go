package commands

import (
	"flag"
	"fmt"
	"os"
)

type HelpCommand struct {
	Command
}

func (h HelpCommand) Description() string {
	return "Shows the help text for a command"
}

func (h HelpCommand) Run() error {
	Usage()
	return nil
}

func (h HelpCommand) RegisterArgs(args ...string) {
}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

func Usage(){
	_,_ = fmt.Fprintf(os.Stdout, `spm is a tool for managing syberos app dependencies

Usage: spm [command] <...args>

Options:
`)
	flag.PrintDefaults()
}




package main

import (
	"commands"
)

func registry(){

}

func main(){
	cmd := &commands.InitCommand{}
	if err := cmd.Run(); err!=nil {
		cmd.Error(err.Error())
	}
}

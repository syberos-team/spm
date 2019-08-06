package main

import (
	"commands"
	"core/util"
)

func registry(){

}

func main(){
	//cmd := &commands.InitCommand{}
	//if err := cmd.Run(); err!=nil {
	//	cmd.Error(err.Error())
	//}

	//_ = util.GitClone("https://github.com/abeir/7788.git", "/home/abeir/workspace/go/abc")
	//_ = util.RemoveDotGit("/home/abeir/workspace/go/abc")

	spmJsonPath := "/home/abeir/doc/test/qpm.json"

	spmJsonConent := &commands.SpmJsonContent{}
	var content interface{} = spmJsonConent

	_ = util.LoadJsonFile(spmJsonPath, &content)
	spmJsonConent.Dependencies = append(spmJsonConent.Dependencies, "com.abeir@124")
	_ = util.WriteStruct(spmJsonPath, spmJsonConent)
}

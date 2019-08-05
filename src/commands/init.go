package commands

import (
	"core/conf"
	"core/util"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//pwd 当前所在目录
var pwd string

func init(){
	path, err := util.Pwd()
	if err!=nil {
		panic(err)
	}
	pwd = path
}


type InitCommand struct {
	Command
	spmJson *conf.SpmJson
}

func (i InitCommand) Description() string {
	return "Initializes a new module in the current directory"
}

func (i InitCommand) Run() error {
	//if err := i.check(); err!=nil {
	//	return err
	//}
	cwd := filepath.Base(pwd)

	i.spmJson = conf.NewSpmJson()

	i.spmJson.Author.Name, _ = <-util.Prompt("Your name:", "")
	i.spmJson.Author.Email, _ = <-util.Prompt("Your email:", "")

	suggestedName := i.extractReverseDomain(i.spmJson.Author.Email) + "." + cwd
	i.spmJson.Name, _ = <-util.Prompt("Unique package name:", suggestedName)
	if i.spmJson.Name == "" {
		return errors.New("must be filled in with unique package name")
	}
	i.spmJson.Description, _ = <-util.Prompt("Briefly describe the project:", "")

	i.spmJson.Repository.Url, _ = <-util.Prompt("Git repository url:", "")
	priFilename, _ := <-util.Prompt("Package .pri file:", i.recommendPriFilename(i.spmJson.Name))
	if priFilename == "" {
		return errors.New("must be filled in with package .pri file")
	}
	if !strings.HasSuffix(priFilename, ".pri"){
		priFilename = priFilename + ".pri"
	}
	i.spmJson.PriFilename = priFilename

	if err := i.generateSpmJson(); err!=nil{
		return err
	}
	if err := i.generateBoilerplate(); err!=nil {
		return err
	}
	return nil
}



func (i InitCommand) extractReverseDomain(email string) string {
	emailParts := strings.Split(email, "@")
	if len(emailParts) != 2 {
		return ""
	}
	domainParts := strings.Split(emailParts[1], ".")
	for i, j := 0, len(domainParts)-1; i < j; i, j = i+1, j-1 {
		domainParts[i], domainParts[j] = domainParts[j], domainParts[i]
	}
	return strings.Join(domainParts, ".")
}

func (i InitCommand) recommendPriFilename(name string) string{
	return strings.ReplaceAll(name, ".", "_") + ".pri"
}

//check 检查当前所在目录必须是空目录
func (i InitCommand) check() error{
	f, err := os.Open(pwd)
	if err!=nil {
		return err
	}
	defer util.CloseQuietly(f)
	dirname, _ := f.Readdirnames(1)
	if len(dirname) > 0 {
		return errors.New("current directory is not empty")
	}
	return nil
}

func (i InitCommand) generateBoilerplate() error{
	subDir := strings.ReplaceAll(i.spmJson.Name, ".", "/")
	if err := os.MkdirAll(subDir, os.ModePerm); err!=nil {
		return err
	}
	prefixName := filepath.Base(i.spmJson.PriFilename)
	model := &util.TemplateModel{
		QrcFile:   prefixName + ".qrc",
		QrcPrefix:  subDir,
		Name:      i.spmJson.Name,
	}
	priPath := filepath.Join(pwd, i.spmJson.PriFilename)
	if err := util.WriteTemplate(priPath, util.PriTemplate, *model); err!=nil {
		return err
	}
	qrcPath := filepath.Join(pwd, model.QrcFile)
	if err := util.WriteTemplate(qrcPath, util.QrcTemplate, *model); err!=nil {
		return err
	}
	qmldirPath := filepath.Join(pwd, "qmldir")
	if err := util.WriteTemplate(qmldirPath, util.QmldirTemplate, *model); err!=nil {
		return err
	}
	return nil
}

func (i InitCommand) generateSpmJson() error{
	data, err := json.MarshalIndent(i.spmJson, "", "\t")
	if err!=nil {
		return err
	}
	return ioutil.WriteFile(filepath.Join(pwd, "spm.json"), data, os.FileMode(0666))
}









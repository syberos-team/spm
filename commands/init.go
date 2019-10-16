package commands

import (
	"errors"
	"fmt"
	"path/filepath"
	"spm/core/conf"
	"spm/core/log"
	"spm/core/util"
	"strings"
)

//InitCommand 创建项目
type InitCommand struct {
	Command
	spmJson *conf.SpmJson
	//是否生成样板
	boilerplate bool
}

func (i *InitCommand) RegisterArgs(args ...string) {
}

func (i *InitCommand) Description() string {
	return "Initializes a new module in the current directory"
}

func (i *InitCommand) ArgsDescription() []ArgsDescription{
	return nil
}

func (i *InitCommand) Run() error {
	if err := i.beforeCheck(); err!=nil {
		return err
	}
	var author string
	var email string
	var repoUrl string

	//查询git仓库信息
	if util.IsGitRepository(pwd) {
		author, email, repoUrl = i.gitInfo()
	}

	i.spmJson = conf.NewSpmJson()
	i.spmJson.Author.Name, _ = <-util.Prompt("Your name:", author)
	i.spmJson.Author.Email, _ = <-util.Prompt("Your email:", email)

	cwd := filepath.Base(pwd)
	suggestedName := i.extractReverseDomain(i.spmJson.Author.Email) + "." + cwd
	if strings.HasPrefix(suggestedName, ".") {
		suggestedName = suggestedName[1:]
	}
	i.spmJson.Name, _ = <-util.Prompt("Unique package name:", suggestedName)
	if i.spmJson.Name == "" {
		return errors.New("must be filled in with unique package name")
	}
	i.spmJson.Version, _ = <-util.Prompt("package version:", i.spmJson.Version)
	if i.spmJson.Version == "" {
		return errors.New("must be filled in with version")
	}
	i.spmJson.Description, _ = <-util.Prompt("Briefly describe the project:", "")

	i.spmJson.Repository.Url, _ = <-util.Prompt("Git repository url:", repoUrl)
	priFilename, _ := <-util.Prompt("Package .pri file:", i.recommendPriFilename(i.spmJson.Name))
	if priFilename == "" {
		return errors.New("must be filled in with package .pri file")
	}
	if !strings.HasSuffix(priFilename, ".pri"){
		priFilename = priFilename + ".pri"
	}
	i.spmJson.PriFilename = priFilename

	bootstrap := <-util.Prompt("Generate boilerplate:", "Y/n")
	if len(bootstrap) == 0 || strings.ToLower(string(bootstrap[0])) == "y" {
		i.boilerplate = true
	}

	//生成文件前再次检查
	if err := i.afterCheck(); err!=nil {
		return err
	}

	if err := i.generateSpmJson(); err!=nil{
		return err
	}

	if i.boilerplate {
		if err := i.generateBoilerplate(); err!=nil {
			return err
		}
	}
	log.Info("Initialized module ", i.spmJson.Name)
	return nil
}



func (i *InitCommand) extractReverseDomain(email string) string {
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

func (i *InitCommand) recommendPriFilename(name string) string{
	return strings.ReplaceAll(name, ".", "_") + ".pri"
}

//beforeCheck 检查当前所在目录下没有spm.json
func (i *InitCommand) beforeCheck() error{
	spmJsonPath := filepath.Join(pwd, SpmJsonFilename)
	if util.IsExists(spmJsonPath) {
		return fmt.Errorf("%s is exists", SpmJsonFilename)
	}
	return nil
}

func (i *InitCommand) afterCheck() error {
	priPath := filepath.Join(pwd, i.spmJson.PriFilename)
	if util.IsExists(priPath) {
		return fmt.Errorf("%s is exists", i.spmJson.PriFilename)
	}
	return nil
}

func (i *InitCommand) generateBoilerplate() error{
	subDir := strings.ReplaceAll(i.spmJson.Name, ".", "/")
	prefixName := strings.TrimSuffix(i.spmJson.PriFilename, ".pri")
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

func (i *InitCommand) generateSpmJson() error{
	return util.WriteStruct(filepath.Join(pwd, SpmJsonFilename), i.spmJson)
}

func (i *InitCommand) gitInfo() (author, email, repoUrl string){
	git := util.NewGit()
	author, _ = git.LastCommitAuthorName()
	email, _ = git.LastCommitEmail()
	repoUrl, _ = git.RepositoryURL()
	return
}

func NewInitCommand() *InitCommand{
	return &InitCommand{}
}









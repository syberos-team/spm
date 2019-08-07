package commands

import (
	"core"
	"core/conf"
	"core/log"
	"core/util"
	"errors"
	"flag"
	"path"
	"strings"
)

//PublishCommand 将包信息上传至服务器
type PublishCommand struct {
	Command
}

func (p *PublishCommand) RegisterFlags(flags *flag.FlagSet) {
}

func (p *PublishCommand) Description() string {
	return "Publishes a new module"
}

func (p *PublishCommand) Run() error {
	spmClient := core.NewSpmClient()
	req := &core.PublishRequest{
		Package:      core.Package{
			Name: "",
			Description: "",
		},
		Author:       core.Author{
			Name: "",
			Email: "",
			Description: "",
		},
		Repository:   core.Repository{
			Url: "",
		},
		Version:      "",
		Dependencies: nil,
		PriFilename:  "",
		Force:        "",
	}
	resp, err := spmClient.Publish(req)
	if err!=nil {
		return err
	}
	if core.CODE_ERROR == resp.Code {
		return errors.New(resp.Msg)
	}
	log.Info("Published module ", req.Package.Name)
	return nil
}

func (p *PublishCommand) loadSpmJson() (*conf.SpmJson, error){
	spmJson := conf.NewSpmJson()
	err := spmJson.Load(path.Join(pwd, SpmJsonFilename))
	if err!=nil {
		return nil, err
	}
	return spmJson, nil
}

func (p *PublishCommand) checkSpmJson(spmJson *conf.SpmJson) (isChecked bool, fail string){
	isChecked = false
	if strings.TrimSpace(spmJson.Name)=="" {
		fail = "the package name must be configured"
		return
	}
	if strings.TrimSpace(spmJson.Version)=="" {
		fail = "the version must be configured"
		return
	}
	if strings.TrimSpace(spmJson.Repository.Url)=="" {
		fail = "the repository url must be configured"
		return
	}
	if strings.TrimSpace(spmJson.PriFilename)=="" {
		fail = "the priFilename must be configured"
		return
	}
	if !util.IsExists(path.Join(pwd, spmJson.PriFilename)) {
		fail = "priFilename configuration must be consistent with file name"
		return
	}
	return true, ""
}

func NewPublishCommand() *PublishCommand {
	return &PublishCommand{}
}




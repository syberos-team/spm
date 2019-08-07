package commands

import (
	"core"
	"core/conf"
	"core/log"
	"core/util"
	"errors"
	"path"
	"strings"
)

//PublishCommand 将包信息上传至服务器
type PublishCommand struct {
	Command
	force bool		//是否强制推送
}

func (p *PublishCommand) RegisterArgs(args ...string) {
	if args==nil || len(args)==0 {
		return
	}
	if strings.TrimSpace(args[0])=="force" {
		p.force = true
	}
}

func (p *PublishCommand) Description() string {
	return "Publishes a new module"
}

func (p *PublishCommand) Run() error {
	log.Debug("run publish command...")
	spmJson, err := p.loadSpmJson()
	if err!=nil {
		return err
	}
	err = p.checkSpmJson(spmJson)
	if err!=nil {
		return err
	}

	req := p.spmJsonToRequest(spmJson)
	spmClient := core.NewSpmClient()
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

func (p *PublishCommand) checkSpmJson(spmJson *conf.SpmJson) error{
	if strings.TrimSpace(spmJson.Name)=="" {
		return errors.New("the package name must be configured")
	}
	if strings.TrimSpace(spmJson.Version)=="" {
		return errors.New("the version must be configured")
	}
	if strings.TrimSpace(spmJson.Repository.Url)=="" {
		return errors.New("the repository url must be configured")
	}
	if strings.TrimSpace(spmJson.PriFilename)=="" {
		return errors.New("the priFilename must be configured")
	}
	if !util.IsExists(path.Join(pwd, spmJson.PriFilename)) {
		return errors.New("priFilename configuration must be consistent with file name")
	}
	//不支持添加依赖
	if spmJson.Dependencies!=nil && len(spmJson.Dependencies)>0 {
		return errors.New("adding dependencies is not supported for the time being")
	}
	return nil
}

func (p *PublishCommand) spmJsonToRequest(spmJson *conf.SpmJson) *core.PublishRequest{
	req := &core.PublishRequest{
		Package:      core.Package{
			Name: spmJson.Name,
			Description: spmJson.Description,
		},
		Author:       core.Author{
			Name: spmJson.Author.Name,
			Email: spmJson.Author.Email,
			Description: spmJson.Author.Description,
		},
		Repository:   core.Repository{
			Url: spmJson.Repository.Url,
		},
		Version:      spmJson.Version,
		Dependencies: spmJson.Dependencies,
		PriFilename:  spmJson.PriFilename,
	}
	if p.force {
		req.Force = "1"
	}
	return req
}

func NewPublishCommand() *PublishCommand {
	return &PublishCommand{}
}




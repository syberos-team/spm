package commands

import (
	"core"
	"errors"
	"flag"
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
	p.Info("Published module ", req.Package.Name)
	return nil
}




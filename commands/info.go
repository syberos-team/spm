package commands

import (
	"errors"
	"html/template"
	"os"
	"spm/core"
	"spm/core/util"
)

//InfoCommand 查询包详情
type InfoCommand struct {
	Command
	packageName string
	version string
}

func (i *InfoCommand) Description() string {
	return "Displays information about the specified package"
}

var infoTemplate = template.Must(template.New("infoTemplate").Parse(`
Name: {{.Package.Name}}
Author: {{.Author.Name}} ({{.Author.Email}})
Repository: {{.Repository.Url}}
Description: {{.Package.Description}}
Dependencies:
{{- with .Dependencies}}
	{{range $index, $dependency := . }}
	{{$dependency}}
	{{end}}
{{else}}
	None.
{{end}}
`))
func (i *InfoCommand) Run() error {
	if i.packageName=="" {
		return errors.New("you must enter the package name")
	}
	data, err := PackageInfo(i.packageName, i.version)
	if err!=nil {
		return err
	}
	if data==nil {
		return errors.New("no package found")
	}
	return i.printData(*data)
}

func (i *InfoCommand) RegisterArgs(args ...string) {
	if args==nil || len(args)==0 {
		return
	}
	i.packageName, i.version = util.ParsePackageInfo(args[0])
}

func (i *InfoCommand) printData(data core.InfoResponseData) error{
	return infoTemplate.Execute(os.Stdout, data)
}

//PackageInfo 访问服务端获取包信息
func PackageInfo(packageName, version string) (*core.InfoResponseData, error){
	client := core.NewSpmClient()
	req := &core.InfoRequest{
		PackageName: packageName,
		Version:     version,
	}
	resp, err := client.Info(req)
	if err!=nil {
		return nil, err
	}
	if core.CODE_ERROR==resp.Code {
		return nil, errors.New(resp.Msg)
	}
	return resp.Data, nil
}

func NewInfoCommand() *InfoCommand{
	return &InfoCommand{}
}


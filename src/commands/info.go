package commands

import (
	"core"
	"errors"
	"flag"
	"html/template"
	"os"
	"strings"
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
	return i.printData(*data)
}

func (i *InfoCommand) RegisterFlags(flags *flag.FlagSet) {
	arg := flags.Arg(0)
	if arg=="" {
		return
	}
	packageInfo := strings.Split(arg, "@")
	if packageInfo==nil || len(packageInfo)==0{
		return
	}
	packageInfoLen := len(packageInfo)
	if packageInfoLen==1{
		i.packageName = packageInfo[0]
	}else if packageInfoLen==2 {
		i.packageName = packageInfo[0]
		i.version = packageInfo[1]
	}
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
	return &resp.Data, nil
}

func NewInfoCommand() *InfoCommand{
	return &InfoCommand{}
}


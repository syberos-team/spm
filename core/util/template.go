package util

import (
	"os"
	"strings"
	"text/template"
)

var (
	PriTemplate = template.Must(template.New("modulePri").Parse(`
RESOURCES += \
    $$PWD/{{.QrcFile}}
`))
	QrcTemplate = template.Must(template.New("moduleQrc").Parse(`
<RCC>
    <qresource prefix="/{{.QrcPrefix}}">
        <file>qmldir</file>
    </qresource>
</RCC>
`))
	QmldirTemplate = template.Must(template.New("qmldir").Parse(`
module {{.Name}}
`))
	VendorTemplate = template.Must(template.New("vendorPri").Parse(`
INCLUDEPATH += $$PWD
QML_IMPORT_PATH += $$PWD

{{- with .IncludePris}}
{{range $index, $pri := . }}{{$pri}}
{{end}}{{end}}`))
)

type TemplateModel struct {
	QrcFile string
	QrcPrefix string
	Name string
	IncludePris []string
}

//WriteTemplate 使用模板生成文件
func WriteTemplate(filePath string, tpl *template.Template, data TemplateModel) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer CloseQuietly(file)

	err = tpl.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}

//TemplateToString 使用模板生成字符串
func TemplateToString(tpl *template.Template, data TemplateModel) (string, error){
	builder := &strings.Builder{}
	err := tpl.Execute(builder, data)
	if err!=nil {
		return "", err
	}
	return builder.String(), nil
}
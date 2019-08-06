package util

import (
	"os"
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

{{- with .Dependencies}}
	{{range $index, $dependency := . }}
include($$PWD/{{$dependency}}.pri)
	{{end}}
{{end}}
`))
)

type TemplateModel struct {
	QrcFile string
	QrcPrefix string
	Name string
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
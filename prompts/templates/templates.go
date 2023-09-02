package templates

import (
	"fmt"
	"text/template"

	"github.com/manifoldco/promptui"
)

const (
	active   = `▸ {{ .Name | cyan | bold }}{{ if .Arn }} ({{ .Arn }}){{end}}`
	inactive = `  {{ .Name | cyan }}{{ if .Arn }} ({{ .Arn }}){{end}}`
	Selected = `{{ "✔" | green }} %s: {{ .Name | cyan }}{{ if .Arn }} ({{ .Arn }}){{end}}`
	success  = `{{ "✔" | green }} {{ .Label }}: {{ .Value }}`
	failure  = `{{ "✗" | red }} {{ .Label }}: {{ .Value }}`
)

type input struct {
	Label string
	Value string
}

var (
	tplSuccess *template.Template
	tplFailure *template.Template
)

func init() {
	tplSuccess = template.Must(template.New("").Funcs(promptui.FuncMap).Parse(success))
	tplFailure = template.Must(template.New("").Funcs(promptui.FuncMap).Parse(failure))
}

func TplAwsResource(resourceType string) *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Active:   active,
		Inactive: inactive,
		Selected: fmt.Sprintf(Selected, resourceType),
	}
}

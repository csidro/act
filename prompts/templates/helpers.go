package templates

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/manifoldco/promptui"
)

func render(tpl *template.Template, data interface{}) string {
	var buf bytes.Buffer
	err := tpl.Execute(&buf, data)
	if err != nil {
		return fmt.Sprintf("%+v", data)
	}
	return buf.String()
}

func PromptSuccess(label, value string) string {
	data := input{Label: label, Value: value}
	return render(tplSuccess, data)
}

func PromptFailure(label, value string) string {
	data := input{Label: label, Value: value}
	return render(tplFailure, data)
}

func SelectSuccess(tpls *promptui.SelectTemplates, data interface{}) string {
	tpl, err := template.New("").Funcs(promptui.FuncMap).Parse(tpls.Selected)
	if err != nil {
		return fmt.Sprintf("%+v", data)
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		return fmt.Sprintf("%+v", data)
	}
	return buf.String()
}

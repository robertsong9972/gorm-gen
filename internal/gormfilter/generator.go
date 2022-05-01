package gormfilter

import (
	"os"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/robertsong9972/gorm-gen/util"
)

var funcMap = template.FuncMap{
	"camel":        util.GoCamelCase,
	"snake":        strcase.ToSnake,
	"goType":       goType,
	"GoType":       GoType,
	"isString":     isString,
	"isNumber":     isNumber,
	"notLastIndex": notLastIndex,
}

type GenOptions struct {
	FilePath string
	GoMod    string
}

func Generate(m Model, options GenOptions) string {
	t, err := template.New("model").Funcs(funcMap).Parse(FilterTemplate)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(options.FilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, m)
	if err != nil {
		panic(err)
	}
	return ""
}

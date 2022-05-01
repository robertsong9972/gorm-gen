package gormfilter

import (
	"fmt"
	"strings"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
)

type ModelField struct {
	Name string
	Type descriptor.FieldDescriptorProto_Type
}

type Model struct {
	Package      string       // e.g. gen
	DbPackage    string       // e.g. git.garena.com/xust/goorm-dot
	DbModuleName string       // e.g. test
	Name         string       // e.g. XTab
	Fields       []ModelField // e.g. field_a, field_b
}

type ParseOptions struct {
	FilePath string
}

func Parse(options ParseOptions) []Model {
	p := protoparse.Parser{}
	desc, err := p.ParseFiles(options.FilePath)
	if err != nil {
		panic(fmt.Sprintf("failed to parse file|filepath=%s|err=%v", options.FilePath, err))
	}

	models := make([]Model, 0)
	for _, d := range desc {
		dbPackage, dbModuleName := getGoPackage(d)
		for _, m := range d.GetMessageTypes() {
			model := Model{
				Package:      "gen",
				DbPackage:    dbPackage,
				DbModuleName: dbModuleName,
				Name:         m.GetName(),
			}
			for _, f := range m.GetFields() {
				model.Fields = append(model.Fields, ModelField{
					Name: f.GetName(),
					Type: f.GetType(),
				})
			}

			models = append(models, model)
		}
	}

	return models
}

// there are 2 possible formats in the go_package option
// 1. internal/query/db_proto;item_attribute"
// 2. internal/query/db_proto"
func getGoPackage(fd *desc.FileDescriptor) (dbPackage string, dbModuleName string) {
	goPackage := *fd.GetFileOptions().GoPackage
	pkg := strings.Split(*fd.GetFileOptions().GoPackage, ";")

	if len(pkg) == 1 {
		lastSlash := strings.LastIndex(goPackage, "/")
		return goPackage, goPackage[lastSlash+1:]
	} else {
		return pkg[0], pkg[1]
	}
}

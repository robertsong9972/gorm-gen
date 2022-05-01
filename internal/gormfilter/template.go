package gormfilter

const FilterTemplate string = `package {{ .Package }}
{{$condName := printf "%sCond" $.Name}}
import (
	"strings"

	gormgen "github.com/robertsong9972/gorm-gen"
)

type {{ $.Name }}Filter struct {
{{ range .Fields }}{{ if isNumber .Type }}
	{{ camel .Name }}Eq    *{{ goType .Type }}
	{{ camel .Name }}Gte   *{{ goType .Type }}
	{{ camel .Name }}Gt    *{{ goType .Type }}
	{{ camel .Name }}Lte   *{{ goType .Type }}
	{{ camel .Name }}Lt    *{{ goType .Type }}
	{{ camel .Name }}Neq   *{{ goType .Type }}
	{{ camel .Name }}In    []{{ goType .Type }}
	{{ camel .Name }}NotIn []{{ goType .Type }}
{{ else if isString .Type}}
	{{ camel .Name }}Eq    *{{ goType .Type }}
	{{ camel .Name }}Neq   *{{ goType .Type }}
	{{ camel .Name }}In    []{{ goType .Type }}
	{{ camel .Name }}NotIn []{{ goType .Type }}
	{{ camel .Name }}Like  *{{ goType .Type }}
{{end}}{{end}}
	Limit      *int
	Offset     *int
	OrderBy    *string
	Select     []string

    condType   string
    builder    strings.Builder
    wheres     string
    args       []interface{}
}

func (f *{{ $.Name }}Filter) GetLimit() int {
	if f != nil && f.Limit != nil {
		return *f.Limit
	}
	return 0
}

func (f *{{ $.Name }}Filter) GetOffset() int {
	if f != nil && f.Offset != nil {
		return *f.Offset
	}
	return 0
}

func (f *{{ $.Name }}Filter) GetOrderBy() string {
	if f != nil && f.OrderBy != nil {
		return *f.OrderBy
	}
	return ""
}

func (f *{{ $.Name }}Filter) ToAndCond() (string, []interface{}) {
	return f.buildCond(gormgen.AND)
}

func (f *{{ $.Name }}Filter) ToOrCond() (string, []interface{}) {
	return f.buildCond(gormgen.OR)
}


func (f *{{ $.Name }}Filter) buildCond(condType string) (string, []interface{}) {
    if f == nil {
		return gormgen.NOTFOUNDCOND, nil
	}
	f.condType = condType
{{ range .Fields }}{{ if isNumber .Type }}
    if f.{{ camel .Name }}Eq != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Eq)
    }
    if f.{{ camel .Name }}Neq != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Neq)
    }
    if f.{{ camel .Name }}Gte != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Gte)
    }
   	if f.{{ camel .Name }}Gt != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Gt)
    }
   	if f.{{ camel .Name }}Lte != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Lte)
    }
   	if f.{{ camel .Name }}Lt != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Lt)
    }
   	if f.{{ camel .Name }}In != nil{
        f.writeCond(" {{snake .Name}}=? ", f.{{ camel .Name }}In)
    }
	if f.{{ camel .Name }}NotIn != nil{
        f.writeCond(" {{snake .Name}}=? ", f.{{ camel .Name }}NotIn)
    }
{{ else if isString .Type}}
    if f.{{ camel .Name }}Eq != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Eq)
    }
    if f.{{ camel .Name }}Neq != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Neq)
    }
	if f.{{ camel .Name }}Like != nil{
        f.writeCond(" {{snake .Name}}=? ", *f.{{ camel .Name }}Like)
    }
   	if f.{{ camel .Name }}In != nil{
        f.writeCond(" {{snake .Name}}=? ", f.{{ camel .Name }}In)
    }
	if f.{{ camel .Name }}NotIn != nil{
        f.writeCond(" {{snake .Name}}=? ", f.{{ camel .Name }}NotIn)
    }
{{end}}{{end}}
	f.wheres = f.builder.String()
	if f.wheres == "" {
		return gormgen.NOTFOUNDCOND, nil
	}
	return f.wheres, f.args
}

func (f *{{ $.Name }}Filter) writeCond(cond string, arg interface{}) {
	if len(f.args) != 0 {
		f.builder.WriteString(f.condType)
	}
	f.builder.WriteString(cond)
	f.args = append(f.args, arg)
}

func (f *{{ $.Name }}Filter) test() bool {
{{ range .Fields }}{{ if or (isNumber .Type) (isString .Type) }}
	if f.{{ camel .Name }}In != nil && len(f.{{ camel .Name }}In) == 0{
		return false
	}
{{end}}{{end}}
	return true
}
`

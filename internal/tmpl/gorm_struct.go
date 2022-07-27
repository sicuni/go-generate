package tmpl

const StructTemplate = NotEditMark + `
package {{.PackageName}}

import "time"

const TableName{{.StructName}} = "{{.TableName}}"

// {{.StructName}} mapped from table <{{.TableName}}>
type {{.StructName}} struct {
    {{range .Fields}}
    {{.Name}} {{.Type}} ` + "`gorm:\"{{.GORMTag}}\" json:\"{{.JSONTag}}\"` " +
	"{{if .ColumnComment}}// {{.ColumnComment}}{{end}}" +
	`{{end}}
}

// TableName {{.StructName}}'s table name
func ({{.S}} *{{.StructName}}) TableName() string {
    return TableName{{.StructName}}
}
`

package tmpl

const NotEditMark = `
// Code generated by github.com/sicuni/go-generate.

`

const HeaderTmpl = NotEditMark + `
package {{.}}

import(
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"
)
`

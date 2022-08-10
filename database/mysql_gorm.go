package database

import (
	"github.com/sicuni/go-generate/internal/model"
	"github.com/sicuni/go-generate/internal/model/column"
	"gorm.io/gorm"
)

type MysqlGorm struct {
	db 		*gorm.DB
	dbConnect string
}

func (g MysqlGorm) GetStructFields(tableName string) (result []*model.Field, err error) {
	columns := make([]column.MysqlColumn, 0)
	schemaName := GetMysqlSchemaName(g.dbConnect)
	if err = g.db.Debug().Raw(MYSQLQUERY, schemaName, tableName).Scan(&columns).Error; err != nil {
		return nil, err
	}

	for _, mysqlColumn := range columns {
		field := mysqlColumn.ToField()
		field.Name = snakeToHump(field.ColumnName)
		result = append(result, field)
	}
	return result, nil
}

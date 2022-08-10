package database

import (
	"fmt"
	"github.com/sicuni/go-generate/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	GetStructFields(tableName string) ([]*model.Field, error)
}

// OpenGorm 连接数据库
func OpenGorm(driverName, dbConnect string) (Database, error) {
	switch driverName {
	case MYSQL:
		db, err := gorm.Open(mysql.Open(dbConnect))
		return &MysqlGorm{db: db, dbConnect: dbConnect}, err
	default:
		return nil, fmt.Errorf("driver name error: driver name: %v", driverName)
	}
}

// snakeToHump 根据数据表列名生成ColumnName
func snakeToHump(str string) string {
	toStr := make([]byte, 0, len(str))
	i := 0
	for i < len(str) {
		if (str[i] < 97 || str[i] > 122) && str[i] != '_' {
			toStr = append(toStr, str[i])
			i++
			continue
		}
		if i == 0 && str[i] >= 97 && str[i] <= 122 {
			toStr = append(toStr, str[i]-32)
			i++
			continue
		}
		if i == len(str)-1 && str[i] == '_' {
			i++
			continue
		}
		if str[i] == '_' {
			if str[i+1] >= 97 && str[i+1] <= 122 {
				toStr = append(toStr, str[i+1]-32)
				i += 2
			}
			continue
		}
		toStr = append(toStr, str[i])
		i++
	}
	return string(toStr)
}

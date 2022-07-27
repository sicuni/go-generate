package test

import (
	"fmt"
	go_generate "go-generate"
	"go-generate/config"
	"go-generate/database"
	"testing"
)

const MysqlConnectstring = "root:root@(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func TestGenBaseStructs(t *testing.T) {
	gen := go_generate.New(config.New(
		config.WithModelPath("./"), //model 代码输出路径
		config.WithDaoPath("./"),   //dao 代码输出路径
	))
	gen.UseDB(database.OpenGorm(database.MYSQL, MysqlConnectstring)) //使用gorm mysql
	gen.BindModel(gen.GenModelAs("user", "User"))             //绑定模型
	if err := gen.Execute(); err != nil {
		fmt.Println(err)
	}
}
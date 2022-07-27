## go-generate
**安装：** 
```azure
go get github.com/sicuni/generate
```
自动生成model、dao代码

**使用说明：**
```azure
        gen := go_generate.New(config.New(
		config.WithModelPath("./"), //model 代码输出路径
		config.WithDaoPath("./"),   //dao 代码输出路径
	))
	gen.UseDB(database.OpenGorm(database.MYSQL, MysqlConnectstring)) //使用gorm mysql
	gen.BindModel(gen.GenModelAs("user", "User"))             //绑定模型
	if err := gen.Execute(); err != nil {
		fmt.Println(err)
	}
```

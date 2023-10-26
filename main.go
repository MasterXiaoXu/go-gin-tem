
package main

import (
	"bv/util"
	"bv/routers"
	"bv/models"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

func main() {
	//连接数据库
	err := util.InitMySql()
	if err !=nil{
		panic(err)
	}
	//程序退出关闭数据库连接
	defer util.Close()
	//绑定模型
	util.MysqlDb.AutoMigrate(&models.User{})
	//注册路由
	r :=routes.SetRouter()

	r.Run(":8080")

	
}
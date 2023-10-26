package util

import (
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/go-ini/ini"
	"log"
	"fmt"
 )


var MysqlDb *gorm.DB

func InitMySql()(err error)  {

	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	
	sec, err := cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		sec.Key("USER").String(),
		sec.Key("PASSWORD").String(),
		sec.Key("HOST").String(),
		sec.Key("PORT").String(),
		sec.Key("DB_NAME").String(),
	)
	//连接数据库
	MysqlDb,err =gorm.Open("mysql",dsn)
	if err !=nil{
		log.Fatalf("Fail to connect mysql: %v", err)
	}
	//验证数据库连接是否成功，若成功，则无异常
	return MysqlDb.DB().Ping()
}

//关闭数据库连接
func Close()  {
	MysqlDb.Close()
}


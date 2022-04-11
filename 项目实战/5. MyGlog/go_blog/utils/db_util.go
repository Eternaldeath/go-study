package utils

import (
	"github.com/astaxie/beego"
)

/*
	初始化时，连接数据库
*/
func init() {
	// 读取配置文件参数
	// 这里可以改进，单独用一个函数进行读取配置
	user := beego.AppConfig.String("mysql_user")
	password := beego.AppConfig.String("mysql_password")
	host := beego.AppConfig.String("mysql_host")
	dbname := beego.AppConfig.String("mysql_dbname")

}

package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqldb := beego.AppConfig.String("mysqldb")
	url := mysqluser+":"+mysqlpass+"@/"+mysqldb+"?charset=utf8"
	orm.RegisterDataBase("default", "mysql", url)
	//orm.RegisterModel(new(User))
}
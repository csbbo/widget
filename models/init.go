package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqldb := beego.AppConfig.String("mysqldb")
	url := mysqluser+":"+mysqlpass+"@/"+mysqldb+"?charset=utf8"
	orm.RegisterDataBase("default", "mysql", url)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

type User struct {
	Username string `orm:"size(20)"`
	Password string `orm:"size(64)"`
}
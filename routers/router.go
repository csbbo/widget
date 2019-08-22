package routers

import (
	"zero/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.UserController{},"post:Login")
	beego.Router("/regist", &controllers.UserController{},"post:Regist")
}

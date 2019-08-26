package routers

import (
	"widget/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.UserController{},"post:Login")
	beego.Router("/regist", &controllers.UserController{},"post:Regist")

	//请求错误
	beego.Router("/401", &controllers.ErrorController{},"*:Unauthorized")

	//样式主题
	beego.Router("/addtheme", &controllers.ThemeController{},"post:AddTheme")
	beego.Router("/deletetheme", &controllers.ThemeController{},"post:DeleteTheme")
	beego.Router("/updatetheme", &controllers.ThemeController{},"post:UpdateTheme")
	beego.Router("/readtheme", &controllers.ThemeController{},"post:ReadTheme")
	beego.Router("/readalltheme", &controllers.ThemeController{},"post:ReadAllTheme")
}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"widget/utils"
)

func init() {
	var FilterUser = func(c *context.Context) {
		token := c.Request.Header["Token"]
		if len(token) == 0 {
			c.Redirect(302, "/401")
		} else if CheckToken(token[0]){
			c.Redirect(302, "/401")
		}
	}
	beego.InsertFilter("/addtheme",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/deletetheme",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/updatetheme",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/readtheme",beego.BeforeRouter,FilterUser)
	beego.InsertFilter("/readalltheme",beego.BeforeRouter,FilterUser)
}

func CheckToken(token string) bool{
	_, err := utils.ParseJWT(token)
	if err != nil{
		return false
	}
	return true
}
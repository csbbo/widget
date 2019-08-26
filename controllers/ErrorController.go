package controllers

import "github.com/astaxie/beego"

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Unauthorized(){
	c.Data["json"] = map[string]string{"message":"please log in"}
	c.ServeJSON()
}
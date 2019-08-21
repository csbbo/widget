package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	beego.Controller
}

type LoginData struct {
	Username string `valid:"Required"`
	Password string `valid:"Required"`
}
func (c *UserController) Login() {
	var loginData LoginData
	valid := validation.Validation{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &loginData)
	v, _ := valid.Valid(&loginData)
	if !v {
		c.Data["json"] = valid.Errors
	}else {
		c.Data["json"] = map[string]interface{}{"status": "succes"}
	}
	c.ServeJSON()
}
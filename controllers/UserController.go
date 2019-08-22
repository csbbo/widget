package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"zero/models"
	"zero/utils"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if len(models.Valid(&user)) != 0{
		c.Data["json"] = models.Valid(&user)[0]
	}else {
		ob := models.ReadUserByName(user)
		if ob.Username == ""{
			c.Data["json"] = map[string]interface{}{"message":"user not exist"}
		}else if ob.Password != utils.Encrypt(user.Password){
			c.Data["json"] = map[string]interface{}{"message":"password error"}
		}else {
			c.Data["json"] = map[string]interface{}{"message":"login success"}
		}
	}
	c.ServeJSON()
}

func (c *UserController) Regist() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if len(models.Valid(&user)) != 0{
		c.Data["json"] = models.Valid(&user)[0]
	}else {
		err := models.AddUser(user)
		if err != nil{
			c.Data["json"] = err.Error()
		}else {
			c.Data["json"] = map[string]interface{}{"message":"add user success"}
		}
	}
	c.ServeJSON()
}
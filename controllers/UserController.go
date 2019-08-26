package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
	"widget/models"
	"widget/utils"
)

type UserController struct {
	beego.Controller
}

//登录
func (c *UserController) Login() {
	var user models.User
	if json.Unmarshal(c.Ctx.Input.RequestBody, &user) != nil {
		fmt.Println("JSON解析出错")
	}
	if len(models.Valid(&user)) != 0 {
		c.Data["json"] = models.Valid(&user)[0]
	} else {
		ob := models.ReadUserByName(user)
		if ob.Username == "" {
			c.Data["json"] = map[string]interface{}{"message":"user not exist"}
		}else if ob.Password != utils.Encrypt(user.Password){
			c.Data["json"] = map[string]interface{}{"message":"password error"}
		}else {
			//get jwt
			jwtToken,_ := utils.CreateJWT(ob.Id,ob.Username)
			//redis
			rs, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				fmt.Println("Connect to redis error", err)
				return
			}
			defer rs.Close()
			_,erro := rs.Do("SET",jwtToken, ob.Id,"EX",4*60*60)
			if erro != nil{
				fmt.Println("redis set failed:", erro)
			}

			c.Data["json"] = map[string]interface{}{"message":"login success","token":jwtToken}
		}
	}
	c.ServeJSON()
}

//注册
func (c *UserController) Regist() {
	var user models.User
	if json.Unmarshal(c.Ctx.Input.RequestBody, &user) != nil {
		fmt.Println("JSON解析出错")
	}
	if len(models.Valid(&user)) != 0{
		c.Data["json"] = models.Valid(&user)[0]
	}else {
		err := models.AddUser(user)
		if err != nil {
			c.Data["json"] = err.Error()
		}else {
			c.Data["json"] = map[string]interface{}{"message":"add user success"}
		}
	}
	c.ServeJSON()
}
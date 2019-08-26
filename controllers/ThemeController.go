package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"widget/models"
)

type ThemeController struct {
	beego.Controller
}
//添加主题
func (c *ThemeController) AddTheme(){
	var theme models.Theme
	if json.Unmarshal(c.Ctx.Input.RequestBody, &theme) != nil {
		fmt.Println("JSON解析出错")
	}
	if len(models.Valid(&theme)) != 0{
		c.Data["json"] = models.Valid(&theme)[0]
	} else {
		err := models.AddTheme(theme)
		if err != nil{
			c.Data["json"] = err.Error()
		}else {
			c.Data["json"] = map[string]interface{}{"message":"add theme success"}
		}
	}
	c.ServeJSON()
}
//删除主题
func (c *ThemeController) DeleteTheme(){
	data := map[string]int{}
	if json.Unmarshal(c.Ctx.Input.RequestBody, &data) != nil {
		fmt.Println("JSON解析出错")
	}
	err := models.DeleteTheme(data["id"])
	if err != nil{
		c.Data["json"] = err.Error()
	}else {
		c.Data["json"] = map[string]interface{}{"message":"delete theme success"}
	}
	c.ServeJSON()
}
//修改主题
func (c *ThemeController) UpdateTheme(){
	var theme models.Theme
	if json.Unmarshal(c.Ctx.Input.RequestBody, &theme) != nil {
		fmt.Println("JSON解析出错")
	}
	if len(models.Valid(&theme)) != 0{
		c.Data["json"] = models.Valid(&theme)[0]
	} else {
		err := models.UpdateTheme(theme)
		if err != nil{
			c.Data["json"] = err.Error()
		}else {
			c.Data["json"] = map[string]interface{}{"message":"update theme success"}
		}
	}
	c.ServeJSON()
}
//查询主题
func (c *ThemeController) ReadTheme(){
	data := map[string]int{}
	if json.Unmarshal(c.Ctx.Input.RequestBody, &data) != nil {
		fmt.Println("JSON解析出错")
	}
	err, theme := models.ReadTheme(data["id"])
	if err != nil {
		c.Data["json"] = err
	}else {
		c.Data["json"] = theme[0]
	}
	c.ServeJSON()
}
//查询所有主题
func (c *ThemeController) ReadAllTheme(){
	err, theme := models.ReadAllTheme()
	if err != nil {
		c.Data["json"] = err
	}else {
		c.Data["json"] = theme
	}
	c.ServeJSON()
}

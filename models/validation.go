package models

import (
	"github.com/astaxie/beego/validation"
)

type User struct {
	Id int
	Username string `valid:"Required"`
	Password string `valid:"Required"`
}
type Theme struct {
	Id int
	Name string `valid:"Required"`
	Image_url string `valid:"Required;MaxSize(200)"`
	Zip_url string `valid:"Required;MaxSize(200)"`
}

//表单验证
func Valid(data interface{})([]*validation.Error){
	valid := validation.Validation{}
	v, _ := valid.Valid(data)
	if !v {
		return valid.Errors
	}
	return valid.Errors
}

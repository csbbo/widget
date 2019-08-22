package models

import "github.com/astaxie/beego/validation"

type User struct {
	Username string `valid:"Required"`
	Password string `valid:"Required"`
}

func Valid(data interface{})([]*validation.Error){
	valid := validation.Validation{}
	v, _ := valid.Valid(data)
	if !v {
		return valid.Errors
	}
	return valid.Errors
}

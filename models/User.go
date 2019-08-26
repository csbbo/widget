package models

import (
	"github.com/astaxie/beego/orm"
	"widget/utils"
)

func AddUser(user User)(error){
	o := orm.NewOrm()
	_, err := o.Raw("INSERT INTO user(username,password) VALUES(?,?)",user.Username,utils.Encrypt(user.Password)).Exec()
	return err
}
func ReadUserByName(user User)(User){
	o := orm.NewOrm()
	o.Raw("SELECT * FROM user WHERE username=?",user.Username).QueryRow(&user)
	return user
}
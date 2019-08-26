package models

import (
	"github.com/astaxie/beego/orm"
)

func AddTheme(theme Theme)(error){
	o := orm.NewOrm()
	values := []string{theme.Name,theme.Image_url,theme.Zip_url}
	_, err := o.Raw("INSERT INTO theme(name,image_url,zip_url) VALUES(?,?,?)",values).Exec()
	return err
}
func DeleteTheme(id int)(error){
	o := orm.NewOrm()
	_, err := o.Raw("DELETE FROM theme WHERE id=?",id).Exec()
	return err
}
func UpdateTheme(theme Theme)(error){
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE theme SET name=?,image_url=?,zip_url=? WHERE id=?",theme.Name,theme.Image_url,theme.Zip_url,theme.Id).Exec()
	return err
}
func ReadTheme(id int)(error,[]orm.Params){
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw("SELECT *FROM theme WHERE id=?",id).Values(&maps)
	return err,maps
}
func ReadAllTheme()(error,[]orm.Params){
	o := orm.NewOrm()
	var maps []orm.Params
	_, err := o.Raw("SELECT *FROM theme").Values(&maps)
	return err,maps
}
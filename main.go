package main

import (
	"github.com/astaxie/beego"
	_ "widget/models"
	_ "widget/routers"
)

func main() {
	beego.Run()
}


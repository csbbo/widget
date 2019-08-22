package main

import (
	"github.com/astaxie/beego"
	_ "zero/models"
	_ "zero/routers"
)

func main() {
	beego.Run()
}


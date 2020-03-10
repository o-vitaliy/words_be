package main

import (
	"github.com/astaxie/beego"
	"words/controllers"
)

func main() {
	beego.Router("/", &controllers.WordsController{}, "get:Get;post:Post")
	beego.Run()
}

package main

import (
	"github.com/astaxie/beego"
)

type MainController struct{
	beego.Controller
}

func (self *MainController) Get(){
	self.Ctx.WriteString("hello,world..")
}

func main()  {
	beego.Router("/",&MainController{})
	beego.Run()
}


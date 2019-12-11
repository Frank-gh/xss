package main

import (
	"github.com/astaxie/beego"
	_ "github.com/Frank-gh/xss/models"
	_ "github.com/Frank-gh/xss/routers"
)

func main() {

	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.BConfig.Listen.HTTPAddr = "0.0.0.0"
	// beego.BConfig.Listen.HTTPPort = 9090

	//beego.InsertFilter("/api/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8360

}

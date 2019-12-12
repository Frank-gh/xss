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

	//beego.InsertFilter("/api/*", beego.BeforeExec, services.FilterFunc, true, true)

	beego.Run() // listen and serve on 0.0.0.0:8360

}

package routers

import (
	"github.com/astaxie/beego"
	"github.com/Frank-gh/xss/controllers"
)
func init() {
	beego.Router("api/index/index", &controllers.IndexController{}, "get:Index_Index")
}
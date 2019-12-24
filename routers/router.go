package routers

import (
	"github.com/astaxie/beego"
	"github.com/Frank-gh/xss/controllers"
)
func init() {
	beego.Router("api/index/index", &controllers.IndexController{}, "get:Index_Index")
	beego.Router("api/goods/category", &controllers.GoodsController{}, "get:Goods_Category")
	beego.Router("api/goods/list", &controllers.GoodsController{}, "get:Goods_List")
}
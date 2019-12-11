package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/Frank-gh/xss/models"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index_Index() {
	o := orm.NewOrm()
	var banners []models.XssAd
	ad := new(models.XssAd)
	o.QueryTable(ad).Filter("ad_position_id", 1).All(&banners)
}
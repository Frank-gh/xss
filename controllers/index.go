package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/Frank-gh/xss/models"
	"github.com/Frank-gh/xss/utils"
)

type IndexController struct {
	beego.Controller
}

type IndexRtnJson struct {
	Banners      []models.XssAd				`json:"banner"`
	Channels     []models.XssChannel		`json:"channel"`
}

func (this *IndexController) Index_Index() {
	o := orm.NewOrm()

	var banners []models.XssAd
	ad := new(models.XssAd)
	o.QueryTable(ad).Filter("ad_position_id", 1).All(&banners)


	var channels []models.XssChannel
	channel := new(models.XssChannel)
	o.QueryTable(channel).OrderBy("sort_order").All(&channels)

	utils.ReturnHTTPSuccess(&this.Controller, IndexRtnJson{banners,channels})

	this.ServeJSON()
}
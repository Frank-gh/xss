package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/Frank-gh/xss/models"
	"github.com/Frank-gh/xss/utils"
)

type CategoryRtnJson struct {
	CurCategory     models.XssChannel   `json:"currentCategory"`
	BrotherCategory []models.XssChannel `json:"brotherCategory"`
}

type GoodsController struct {
	beego.Controller
}

func (this *GoodsController) Goods_Index() {
}

func (this *GoodsController) Goods_Category() {
	channelId := this.GetString("id")
	intchannelId := utils.String2Int(channelId)

	o := orm.NewOrm()

	var curcategory models.XssChannel
	var brothercategory []models.XssChannel

	channel := new(models.XssChannel)

	o.QueryTable(channel).Filter("id", intchannelId).One(&curcategory)
	o.QueryTable(channel).All(&brothercategory)

	utils.ReturnHTTPSuccess(&this.Controller, CategoryRtnJson{CurCategory: curcategory, BrotherCategory: brothercategory})
	this.ServeJSON()
}
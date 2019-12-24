package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/Frank-gh/xss/models"
	"github.com/Frank-gh/xss/utils"
	"net/url"
)

type CategoryRtnJson struct {
	CurCategory     models.XssChannel   `json:"currentCategory"`
	BrotherCategory []models.XssChannel `json:"brotherCategory"`
}

type GoodsController struct {
	beego.Controller
}

type FilterCategory struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Checked bool   `json:"checked"`
}

type ListRtnJson struct {
	utils.PageData
	GoodsList        []orm.Params     `json:"goodsList"`
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

func (this *GoodsController) Goods_List() {
	categoryId := this.GetString("categoryId")
	brandId := this.GetString("brandId")
	keyword := this.GetString("keyword")
	isNew := this.GetString("isNew")
	isHot := this.GetString("isHot")
	page := this.GetString("page")
	size := this.GetString("size")
	sort := this.GetString("sort")
	order := this.GetString("order")

	var intsize int = 10
	if size != "" {
		intsize = utils.String2Int(size)
	}

	var intpage int = 1
	if page != "" {
		intpage = utils.String2Int(page)
	}

	o := orm.NewOrm()
	goodstable := new(models.XssGoods)
	rs := o.QueryTable(goodstable)

	if isNew != "" {
		rs = rs.Filter("is_new", isNew)
	}
	if isHot != "" {
		rs = rs.Filter("is_hot", isHot)
	}

	keyword, _ = url.QueryUnescape(keyword)
	if keyword != "" {

		rs = rs.Filter("name__icontains", keyword)
		searchhistory := models.XssSearchHistory{Keyword: keyword, UserId: utils.Int2String(getLoginUserId()),
			AddTime: utils.GetTimestamp()}
		o.Insert(&searchhistory)
	}
	if brandId != "" {
		rs = rs.Filter("brand_id", brandId)
	}

	if categoryId != "" {
		intcategoryId := utils.String2Int(categoryId)
		if intcategoryId > 0 {
			rs = rs.Filter("category_id", intcategoryId)
		}
	}

	if sort == "price" {
		orderstr := "retail_price"
		if order == "desc" {
			orderstr = "-" + orderstr
		}
		rs = rs.OrderBy(orderstr)
	} else {
		rs = rs.OrderBy("-id")
	}

	var rawData []orm.Params
	rs.Values(&rawData, "id", "name", "list_pic_url", "retail_price")
	updateJsonKeysGoods(rawData)

	pageData := utils.GetPageData(rawData, intpage, intsize)

	utils.ReturnHTTPSuccess(&this.Controller, ListRtnJson{PageData: pageData, GoodsList: pageData.Data.([]orm.Params)})
	this.ServeJSON()
}

func updateJsonKeysGoods(vals []orm.Params) {

	for _, val := range vals {
		for k, v := range val {
			switch k {
			case "Id":
				delete(val, k)
				val["id"] = v
			case "Name":
				delete(val, k)
				val["name"] = v
			case "ListPicUrl":
				delete(val, k)
				val["list_pic_url"] = v
			case "RetailPrice":
				delete(val, k)
				val["retail_price"] = v
			}
		}
	}
}

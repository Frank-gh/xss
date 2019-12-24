package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type XssAd struct {
	AdPositionId int    `json:"ad_position_id"`
	Content      string `json:"content"`
	Enabled      int    `json:"enabled"`
	EndTime      int    `json:"end_time"`
	Id           int    `json:"id"`
	ImageUrl     string `json:"image_url"`
	Link         string `json:"link"`
	MediaType    int    `json:"media_type"`
	Name         string `json:"name"`
}

type XssChannel struct {
	IconUrl   string `json:"icon_url"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	SortOrder int    `json:"sort_order"`
	Url       string `json:"url"`
}
type XssGoods struct {
	Id                int    `json:"id"`
	CategoryId        int    `json:"category_id"`
	AddTime           int    `json:"add_time"`
	AppExclusivePrice string `json:"app_exclusive_price"`
	AttributeCategory int    `json:"attribute_category"`
	BrandId           int    `json:"brand_id"`
	CounterPrice      string `json:"counter_price"`
	ExtraPrice        string `json:"extra_price"`
	GoodsBrief        string `json:"goods_brief"`
	GoodsDesc         string `json:"goods_desc"`
	GoodsNumber       int    `json:"goods_number"`
	GoodsSn           string `json:"goods_sn"`
	GoodsUnit         string `json:"goods_unit"`
	IsAppExclusive    int    `json:"is_app_exclusive"`
	IsDelete          bool   `json:"is_delete"`
	IsHot             int    `json:"is_hot"`
	IsLimited         int    `json:"is_limited"`
	IsNew             int    `json:"is_new"`
	IsOnSale          int    `json:"is_on_sale"`
	Keywords          string `json:"keywords"`
	ListPicUrl        string `json:"list_pic_url"`
	Name              string `json:"name"`
	PrimaryPicUrl     string `json:"primary_pic_url"`
	PrimaryProductId  int    `json:"primary_product_id"`
	PromotionDesc     string `json:"promotion_desc"`
	PromotionTag      string `json:"promotion_tag"`
	RetailPrice       string `json:"retail_price"`
	SellVolume        int    `json:"sell_volume"`
	SortOrder         int    `json:"sort_order"`
	UnitPrice         string `json:"unit_price"`
}
type XssSearchHistory struct {
	AddTime int64  `json:"add_time"`
	From    string `json:"from"`
	Id      int    `json:"id"`
	Keyword string `json:"keyword"`
	UserId  string `json:"user_id"`
}
func init() {
	// 开启debug日志
	orm.Debug = true
	// 1. 注册数据驱动, mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/xss?charset=utf8mb4", 30)

	// register model
	orm.RegisterModel(new(XssAd))
	orm.RegisterModel(new(XssChannel))
	orm.RegisterModel(new(XssGoods))
	
	// 自动创建表 参数二为是否开启创建表   参数三是否更新表
	// orm.RunSyncdb("default", true, true)
}
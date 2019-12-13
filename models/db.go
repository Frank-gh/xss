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
	// 自动创建表 参数二为是否开启创建表   参数三是否更新表
	// orm.RunSyncdb("default", true, true)
}
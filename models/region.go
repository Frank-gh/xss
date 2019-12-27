package models

import (
	"github.com/astaxie/beego/orm"
)

func GetRegionName(regionid int) string {

	o := orm.NewOrm()
	regiontable := new(XssRegion)
	var region XssRegion
	o.QueryTable(regiontable).Filter("id", regionid).One(&region)

	return region.Name

}

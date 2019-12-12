package utils

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type HTTPData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func ReturnHTTPSuccess(this *beego.Controller, val interface{}) {

	rtndata := HTTPData{
		ErrNo:  0,
		ErrMsg: "",
		Data:   val,
	}

	data, err := json.Marshal(rtndata)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = json.RawMessage(string(data))
	}
}

func GetHTTPRtnJsonData(errno int, errmsg string) interface{} {

	rtndata := HTTPData{
		ErrNo:  errno,
		ErrMsg: errmsg,
		Data:   nil,
	}
	data, _ := json.Marshal(rtndata)

	return json.RawMessage(string(data))

}

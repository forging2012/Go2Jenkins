package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about index
type IndexController struct {
	beego.Controller
}

// @Title GetAll
// @Description hellp api
// @Success 200 hello api
// @router / [get]
func (i *IndexController) Get() {
	i.Data["json"] = "hello api"
	i.ServeJSON()
}

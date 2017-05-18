package controllers

import (
	"devcloud/models"

	"github.com/astaxie/beego"
)

// Operations about devclouds
type DevCloudController struct {
	beego.Controller
}

// @Title DevCloud Create Project
// @Description Create project
// @Param       project_name             query    string  true            "project name"
// @Param       svn_url             query    string  true            "project svn"
// @Success 200 {"result":"Create sucess |Create has error"}
// @Failure 50X app has error
// @router /create [get]
func (d *DevCloudController) Create() {
        ProjectName := d.GetString("project_name")
        SvnUrl := d.GetString("svn_url")
        ret := models.GetCreateResult(ProjectName,SvnUrl)
        d.Data["json"] = map[string]string{"result": ret}
        d.ServeJSON()
}

// @Title DevCloud checkout code
// @Description checkout code
// @Param       project_name             query    string  true            "project name"
// @Success 200 {"result":"Checkout sucess |Checkout has error"}
// @Failure 50X app has error
// @router /checkout [get]
func (d *DevCloudController) CheckOut() {
        ProjectName := d.GetString("project_name")
        ret := models.GetCheckOutResult(ProjectName)
        d.Data["json"] = map[string]string{"result": ret}
        d.ServeJSON()
}

// @Title DevCloud compile code
// @Description compile code
// @Param       project_name             query    string  true            "project name"
// @Success 200 {"result":"Compile sucess |Compile has error"}
// @Failure 50X app has error
// @router /compile [get]
func (d *DevCloudController) Compile() {
        ProjectName := d.GetString("project_name")
        ret := models.GetCompileResult(ProjectName)
        d.Data["json"] = map[string]string{"result": ret}
        d.ServeJSON()
}

// @Title DevCloud Pack Project
// @Description Pack project
// @Param       project_name             query    string  true            "project_name"
// @Param       version             query    string  true            "version"
// @Success 200 {"result":"Pack sucess http://****|Pack has error"}
// @Failure 50X app has error
// @router /pack [get]
func (d *DevCloudController) Pack() {
	ProjectName := d.GetString("project_name")
	Version := d.GetString("version")
	ret := models.GetPackResult(ProjectName,Version)
	d.Data["json"] = map[string]string{"result": ret}
	d.ServeJSON()
}

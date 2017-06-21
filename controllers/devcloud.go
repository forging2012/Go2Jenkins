package controllers

import (
	//"strings"
	"devcloud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about devclouds
type DevCloudController struct {
	beego.Controller
}

// @Title DevCloud Create Project
// @Description Create project
// @Param       project_name             query    string  true            "project name"
// @Param       svn_url             query    string  true            "project svn"
// @Success 200 {object} models.Result
// @Success 403 forbidden
// @Failure 50X app has error
// @router /create [get]
func (d *DevCloudController) Create() {
	var ret map[string]string
	clientip := d.Ctx.Input.IP()
	ProjectName := d.GetString("project_name")
	SvnUrl := d.GetString("svn_url")
	ret = models.GetCreateResult(ProjectName, SvnUrl)
	logs.Info(clientip + " create " + ProjectName + " " + SvnUrl)
	d.Data["json"] = ret
	d.ServeJSON()
}

// @Title DevCloud checkout code
// @Description checkout code
// @Param       project_name             query    string  true            "project name"
// @Success 200 {object} models.Result
// @Success 403 forbidden
// @Failure 50X app has error
// @router /checkout [get]
func (d *DevCloudController) CheckOut() {
	var ret map[string]string
	clientip := d.Ctx.Input.IP()
	ProjectName := d.GetString("project_name")
	ret = models.GetCheckOutResult(ProjectName)
	logs.Info(clientip + " checkout " + ProjectName)
	d.Data["json"] = ret
	d.ServeJSON()
}

// @Title DevCloud code check
// @Description code check
// @Param       project_name             query    string  true            "project name"
// @Success 200 {object} models.Result
// @Success 403 forbidden
// @Failure 50X app has error
// @router /codecheck [get]
func (d *DevCloudController) CodeCheck() {
	var ret map[string]string
	clientip := d.Ctx.Input.IP()
	ProjectName := d.GetString("project_name")
	ret = models.GetCodeCheckResult(ProjectName)
	logs.Info(clientip + " codecheck " + ProjectName)
	d.Data["json"] = ret
	d.ServeJSON()
}

// @Title DevCloud compile code
// @Description compile code
// @Param       project_name             query    string  true            "project name"
// @Param       jdk_version             query    string  true            "jdk version {1.5 1.6 1.7 1.8}"
// @Success 200 {object} models.Result
// @Success 403 forbidden
// @Failure 50X app has error
// @router /compile [get]
func (d *DevCloudController) Compile() {
	var ret map[string]string
	clientip := d.Ctx.Input.IP()
	ProjectName := d.GetString("project_name")
	JdkVersion := d.GetString("jdk_version")
	ret = models.GetCompileResult(ProjectName, JdkVersion)
	logs.Info(clientip + " compile " + ProjectName)
	d.Data["json"] = ret
	d.ServeJSON()
}

// @Title DevCloud Pack Project
// @Description Pack project
// @Param       project_name             query    string  true            "project name"
// @Param       version             query    string  true            "project version"
// @Param       isE             query    string  true            "是否紧急发布Y/N"
// @Success 200 {object} models.Result
// @Success 403 forbidden
// @Failure 50X app has error
// @router /pack [get]
func (d *DevCloudController) Pack() {
	var ret map[string]string
	clientip := d.Ctx.Input.IP()
	ProjectName := d.GetString("project_name")
	Version := d.GetString("version")
	IsE := d.GetString("isE")
	ret = models.GetPackResult(ProjectName, Version, IsE)
	logs.Info(clientip + " pack " + ProjectName)
	d.Data["json"] = ret
	d.ServeJSON()
}

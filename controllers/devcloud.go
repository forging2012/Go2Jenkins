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
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /create [get]
func (d *DevCloudController) Create() {
	clientip := d.Ctx.Input.IP()
	ProjectName := d.GetString("project_name")
	SvnUrl := d.GetString("svn_url")
	resp := make(map[string]interface{})
	if ProjectName != "" && SvnUrl != "" {
		logs.Info(clientip + " create " + ProjectName + " " + SvnUrl)
		resp = models.GetCreateResult(ProjectName, SvnUrl)
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	d.Data["json"] = resp
	d.ServeJSON()
}

// @Title DevCloud checkout code
// @Description checkout code
// @Param       project_name             query    string  true            "project name"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /checkout [get]
func (d *DevCloudController) CheckOut() {
	if ProjectName := d.GetString("project_name"); ProjectName != "" {
		resp := make(map[string]string)
		logs.Info(d.Ctx.Input.IP() + " checkout " + ProjectName)
		resp = models.GetCheckOutResult(ProjectName)
		d.Data["json"] = resp
		d.ServeJSON()
	} else {
		resp := make(map[string]interface{})
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
		d.Data["json"] = resp
		d.ServeJSON()
	}
}

// @Title DevCloud code check
// @Description code check
// @Param       project_name             query    string  true            "project name"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /codecheck [get]
func (d *DevCloudController) CodeCheck() {
	if ProjectName := d.GetString("project_name"); ProjectName != "" {
		resp := make(map[string]string)
		logs.Info(d.Ctx.Input.IP() + " codecheck " + ProjectName)
		resp = models.GetCodeCheckResult(ProjectName)
		d.Data["json"] = resp
		d.ServeJSON()
	} else {
		resp := make(map[string]interface{})
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
		d.Data["json"] = resp
		d.ServeJSON()
	}
}

// @Title DevCloud compile code
// @Description compile code
// @Param       project_name             query    string  true            "project name"
// @Param       jdk_version             query    string  true            "jdk version {1.5 1.6 1.7 1.8}"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /compile [get]
func (d *DevCloudController) Compile() {
	ProjectName := d.GetString("project_name")
	JdkVersion := d.GetString("jdk_version")
	if ProjectName != "" && JdkVersion != "" {
		logs.Info(d.Ctx.Input.IP() + " compile " + ProjectName)
		resp := make(map[string]string)
		resp = models.GetCompileResult(ProjectName, JdkVersion)
		d.Data["json"] = resp
		d.ServeJSON()
	} else {
		resp := make(map[string]interface{})
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
		d.Data["json"] = resp
		d.ServeJSON()
	}

}

// @Title DevCloud Pack Project
// @Description Pack project
// @Param       project_name             query    string  true            "project name"
// @Param       version             query    string  true            "project version"
// @Param       isE             query    string  true            "是否紧急发布Y/N"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /pack [get]
func (d *DevCloudController) Pack() {
	ProjectName := d.GetString("project_name")
	Version := d.GetString("version")
	IsE := d.GetString("isE")
	if ProjectName != "" && Version != "" && IsE != "" {
		resp := make(map[string]string)
		logs.Info(d.Ctx.Input.IP() + " pack " + ProjectName)
		resp = models.GetPackResult(ProjectName, Version, IsE)
		d.Data["json"] = resp
		d.ServeJSON()
	} else {
		resp := make(map[string]interface{})
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
		d.Data["json"] = resp
		d.ServeJSON()
	}
}

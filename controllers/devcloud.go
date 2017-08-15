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
		if ret, err := models.GetJsonFromMsi(resp); err != nil {
			logs.Info(err.Error())
		} else {

			logs.Info(ret)
		}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	d.Data["json"] = resp
	d.ServeJSON()
}

// @Title DevCloud checkout code
// @Description checkout code
// @Param       project_name             query    string  true            "project name"
// @Param       flow_id             query    string  true            "每个持续集成流程id"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /checkout [get]
func (d *DevCloudController) CheckOut() {
	ProjectName := d.GetString("project_name")
	Flowid := d.GetString("flow_id")
	if ProjectName != "" && Flowid != "" {
		resp := make(map[string]string)
		logs.Info(d.Ctx.Input.IP() + " checkout " + ProjectName + " " + Flowid)
		resp = models.GetCheckOutResult(ProjectName, Flowid)
		if ret, err := models.GetJsonFromMss(resp); err != nil {
			logs.Info(err.Error())
		} else {

			logs.Info(ret)
		}
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
// @Param       flow_id             query    string  true            "每个持续集成流程id"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /codecheck [get]
func (d *DevCloudController) CodeCheck() {
	ProjectName := d.GetString("project_name")
	Flowid := d.GetString("flow_id")
	if ProjectName != "" && Flowid != "" {
		resp := make(map[string]string)
		logs.Info(d.Ctx.Input.IP() + " codecheck " + ProjectName + " " + Flowid)
		resp = models.GetCodeCheckResult(ProjectName, Flowid)
		if ret, err := models.GetJsonFromMss(resp); err != nil {
			logs.Info(err.Error())
		} else {

			logs.Info(ret)
		}
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
// @Param       flow_id             query    string  true            "每个持续集成流程id"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /compile [get]
func (d *DevCloudController) Compile() {
	ProjectName := d.GetString("project_name")
	JdkVersion := d.GetString("jdk_version")
	Flowid := d.GetString("flow_id")
	if ProjectName != "" && JdkVersion != "" && Flowid != "" {
		logs.Info(d.Ctx.Input.IP() + " compile " + ProjectName + " " + Flowid)
		resp := make(map[string]string)
		resp = models.GetCompileResult(ProjectName, JdkVersion, Flowid)
		if ret, err := models.GetJsonFromMss(resp); err != nil {
			logs.Info(err.Error())
		} else {

			logs.Info(ret)
		}
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
// @Param       jdk_version             query    string  true            "jdk version {1.5 1.6 1.7 1.8}"
// @Param       flow_id             query    string  true            "每个持续集成流程id"
// @Success 200 {object} models.Result
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /pack [get]
func (d *DevCloudController) Pack() {
	ProjectName := d.GetString("project_name")
	Version := d.GetString("version")
	IsE := d.GetString("isE")
	JdkVersion := d.GetString("jdk_version")
	Flowid := d.GetString("flow_id")
	if ProjectName != "" && Version != "" && IsE != "" && JdkVersion != "" && Flowid != "" {
		resp := make(map[string]string)
		logs.Info(d.Ctx.Input.IP() + " pack " + ProjectName + " " + Flowid)
		resp = models.GetPackResult(ProjectName, Version, IsE, JdkVersion, Flowid)
		if ret, err := models.GetJsonFromMss(resp); err != nil {
			logs.Info(err.Error())
		} else {

			logs.Info(ret)
		}
		d.Data["json"] = resp
		d.ServeJSON()
	} else {
		resp := make(map[string]interface{})
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
		d.Data["json"] = resp
		d.ServeJSON()
	}
}

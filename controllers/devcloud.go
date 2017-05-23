package controllers

import (
	"strings"
	"devcloud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about devclouds
type DevCloudController struct {
	beego.Controller
}

func init() {
	logs.SetLogger(logs.AdapterFile,`{"filename":"test.log","level":7}`)
}

func access(d *DevCloudController) (string,bool){ 
	clientip := d.Ctx.Input.IP()
	allow_ip := beego.AppConfig.String("allow_ip")
	s_allow_ip := strings.Split(allow_ip,";")
	for _,ip := range s_allow_ip {
		if (clientip == ip){ 
			return clientip,true
		}
	}
	return clientip,false
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
	clientip,isaccess := access(d)
	var ret map[string]string
	if isaccess {
        	ProjectName := d.GetString("project_name")
        	SvnUrl := d.GetString("svn_url")
        	ret = models.GetCreateResult(ProjectName,SvnUrl)
		logs.Info(clientip+" create "+ProjectName+" "+SvnUrl)
	} else {
		ret = map[string]string{"status": "403"}
		logs.Warn(clientip+" is not allow to access")
	}
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
	clientip,isaccess := access(d)
	var ret map[string]string
	if isaccess {
        	ProjectName := d.GetString("project_name")
        	ret = models.GetCheckOutResult(ProjectName)
		logs.Info(clientip+" checkout "+ProjectName)
	} else {
		ret = map[string]string{"status": "403"}
		logs.Warn(clientip+" is not allow to access")
	}
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
	clientip,isaccess := access(d)
        var ret map[string]string
        if isaccess {
        	ProjectName := d.GetString("project_name")
        	ret = models.GetCodeCheckResult(ProjectName)
		logs.Info(clientip+" codecheck "+ProjectName)
	} else {
		ret = map[string]string{"status": "403"}
		logs.Warn(clientip+" is not allow to access")
	}
        d.Data["json"] = ret
        d.ServeJSON()
}

// @Title DevCloud compile code
// @Description compile code
// @Param       project_name             query    string  true            "project name"
// @Success 200 {object} models.Result
// @Success 403 forbidden
// @Failure 50X app has error
// @router /compile [get]
func (d *DevCloudController) Compile() {
	clientip,isaccess := access(d)
        var ret map[string]string
        if isaccess {
        	ProjectName := d.GetString("project_name")
        	ret = models.GetCompileResult(ProjectName)
		logs.Info(clientip+" compile "+ProjectName)
	} else {
                ret = map[string]string{"status": "403"}
		logs.Warn(clientip+" is not allow to access")
        }
        d.Data["json"] = ret
        d.ServeJSON()
}

// @Title DevCloud Pack Project
// @Description Pack project
// @Param       project_name             query    string  true            "project_name"
// @Param       version             query    string  true            "version"
// @Success 200 {object} models.Result
// @Success 403 forbidden
// @Failure 50X app has error
// @router /pack [get]
func (d *DevCloudController) Pack() {
	clientip,isaccess := access(d)
	var ret map[string]string
	if isaccess {
		ProjectName := d.GetString("project_name")
		Version := d.GetString("version")
		ret = models.GetPackResult(ProjectName,Version)
		logs.Info(clientip+" pack "+ProjectName)
	} else {
                ret = map[string]string{"status": "403"}
		logs.Warn(clientip+" is not allow to access")
        }
	d.Data["json"] = ret
	d.ServeJSON()
}

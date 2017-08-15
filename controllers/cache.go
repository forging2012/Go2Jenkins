package controllers

import (
	"devcloud/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about memory cache
type CacheController struct {
	beego.Controller
}

// @Title DelMemCache
// @Description delete mem cache
// @Param	project_name	query 	string	true		"project name"
// @Success 200 {"status": 200}
// @Failure 10016 Miss required parameter
// @Failure 10018 is not in memory cache
// @Failure 403 Forbidden
// @router /del [get]
func (c *CacheController) Del() {
	resp := make(map[string]interface{})
	if projectname := c.GetString("project_name"); projectname != "" {
		if isexit := models.IsExistInMem(projectname); isexit {
			logs.Info(c.Ctx.Input.IP() + " Del Memory cache " + projectname)
			models.DelCacheFromMem(projectname)
			resp = map[string]interface{}{"status": 200}
		} else {
			resp = map[string]interface{}{"status": 10018, "error": "cache is not exist"}
		}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title Get Memory cache
// @Description Get Memory cache
// @Param       project_name    query   string  true            "project name"
// @Success 200 {"status": 200}
// @Failure 403 Forbidden
// @router /get [get]
func (c *CacheController) Get() {
	resp := make(map[string]interface{})
	if projectname := c.GetString("project_name"); projectname != "" {
		if isexit := models.IsExistInMem(projectname); isexit {
			logs.Info(c.Ctx.Input.IP() + " Get Memory cache " + projectname)
			ret := models.GetCacheFromMem(projectname)
			resp = map[string]interface{}{"status": 200, projectname: ret}
		} else {
			resp = map[string]interface{}{"status": 10018, "error": "cache is not exist"}
		}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	c.Data["json"] = resp
	c.ServeJSON()
}

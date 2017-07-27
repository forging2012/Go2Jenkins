package controllers

import (
	"devcloud/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about run cmd
type SshController struct {
	beego.Controller
}

// @Title runcmd
// @Description run remote cmd
// @Param       body            body    models.SshServer   true            "The server ssh info"
// @Success 200 {object} models.CmdResult
// @Failure 403 body is empty
// @router / [post]
func (s *SshController) Post() {
	var servers models.SshServers
	json.Unmarshal(s.Ctx.Input.RequestBody, &servers)
	if len(servers) > 0 {
		ret := models.Run(servers)
		s.Data["json"] = ret
	} else {
		s.Data["json"] = map[string]string{"error": "args is null"}
	}
	s.ServeJSON()
}

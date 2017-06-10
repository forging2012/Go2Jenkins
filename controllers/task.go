package controllers

import (
	"test/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type AddTaskController struct {
	beego.Controller
}

// @Title AddTask
// @Description add task
// @Param	project_name		query 	string	true		"project_name"
// @Param	spec		query 	string	true		"task time"
// @Success 200 ***** ***** 
// @Failure 403 body is empty
// @router / [get]
func (at *AddTaskController) AddTk() {
	project_name := at.GetString("project_name")
	spec := at.GetString("spec")
	models.AddT(project_name,spec)
	o.Data["json"] = map[string]string{"status": 200}
	o.ServeJSON()
}

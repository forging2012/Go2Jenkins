package controllers

import (
	"devcloud/models"
	//"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type TaskController struct {
	beego.Controller
}

// @Title AddTask
// @Description add task
// @Param	project_name		query 	string	true		"project_name"
// @Param	spec		query 	string	true		"task time"
// @Param	tasklist		query 	string	true		"task list"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /add [get]
func (t *TaskController) Add() {
	project_name := t.GetString("project_name")
	spec := t.GetString("spec")
	tasklist := t.GetString("tasklist")
	models.AddTask(project_name, spec, tasklist)
	t.Data["json"] = map[string]int{"status": 200}
	t.ServeJSON()
}

// @Title DelTask
// @Description delete task
// @Param	project_name		query 	string	true		"project_name"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /del [get]
func (t *TaskController) Del() {
	project_name := t.GetString("project_name")
	models.DelTask(project_name)
	t.Data["json"] = map[string]int{"status": 200}
	t.ServeJSON()
}

// @Title DelTask
// @Description delete task
// @Param       project_name            query   string  true            "project_name"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /delay [get]
func (t *TaskController) Delay() {
	project_name := t.GetString("project_name")
	models.DelayTask(project_name)
	t.Data["json"] = map[string]int{"status": 200}
	t.ServeJSON()
}

package controllers

import (
	"devcloud/models"
	"time"
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
// @Param	tasklist		query 	string	true		"task list:checkout;codecheck;compile;pack"
// @Success 200 {"status": 200,"task_in_estype":"crontask", "task_in_esid":id}
// @Failure 403 body is empty
// @router /add [get]
func (t *TaskController) Add() {
	project_name := t.GetString("project_name")
	//task_name := t.GetString("task_name")
	task_name := models.MD5(time.Now().Format("2006-01-02 15:04:05"))
	spec := t.GetString("spec")
	tasklist := t.GetString("tasklist")
	models.AddTask(project_name, task_name, spec, tasklist)
	t.Data["json"] = map[string]interface{}{"status": 200, "task_in_estype": "crontask", "task_in_esid": project_name + "-" + task_name}
	t.ServeJSON()
}

// @Title DelTask
// @Description delete task
// @Param	taskid		query 	string	true		"task id"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /del [get]
func (t *TaskController) Del() {
	taskid := t.GetString("taskid")
	models.DelTask(taskid)
	t.Data["json"] = map[string]int{"status": 200}
	t.ServeJSON()
}

// @Title Get all Task
// @Description Get all task
// @Success 200 {object} models.CronAllInfo
// @Failure 403 body is empty
// @router /getall [get]
func (t *TaskController) GetALL() {
	ret := models.GetAllTask()
	t.Data["json"] = ret
	t.ServeJSON()
}

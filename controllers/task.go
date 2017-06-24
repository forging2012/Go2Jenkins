package controllers

import (
	"devcloud/models"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about task
type TaskController struct {
	beego.Controller
}

// @Title AddTask
// @Description add task
// @Param	project_name		query 	string	true		"project_name"
// @Param	spec		query 	string	true		"秒 分钟 小时 天 月 星期"
// @Param	tasklist		query 	string	true		"从checkout;codecheck;compile|jdk_version;pack|project_version中组合,多个以分号隔开"
// @Success 200 {"status": 200,"task_in_estype":"crontask", "task_in_esid":id}
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /add [get]
func (t *TaskController) Add() {
	project_name := t.GetString("project_name")
	task_name := models.MD5(time.Now().Format("2006-01-02 15:04:05"))
	spec := t.GetString("spec")
	tasklist := t.GetString("tasklist")
	resp := make(map[string]interface{})
	isok := true
	if project_name != "" && spec != "" && tasklist != "" {
		tks := strings.Split(tasklist, ";")
		for _, tk := range tks {
			tk_name := strings.Split(tk, "|")[0]
			if tk_name == "compile" {
				num := len(strings.Split(tk, "|"))
				if num == 1 {
					resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
					isok = false
				}
			}
			if tk_name == "pack" {
				num := len(strings.Split(tk, "|"))
				if num == 1 {
					resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
					isok = false
				}
			}
		}
		if isok {
			logs.Info(t.Ctx.Input.IP() + " Add Task " + project_name + " " + task_name + " " + spec + " " + tasklist)
			models.AddTask(project_name, task_name, spec, tasklist, time.Now())
			resp = map[string]interface{}{"status": 200, "task_in_estype": "crontask", "task_in_esid": project_name + "-" + task_name}
		} else {
			resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
		}

	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	t.Data["json"] = resp
	t.ServeJSON()
}

// @Title DelTask
// @Description delete task
// @Param	taskid		query 	string	true		"task id"
// @Success 200 {"status": 200}
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /del [get]
func (t *TaskController) Del() {
	resp := make(map[string]interface{})
	if taskid := t.GetString("taskid"); taskid != "" {
		logs.Info(t.Ctx.Input.IP() + " Del Task " + taskid)
		models.DelTask(taskid)
		resp = map[string]interface{}{"status": 200}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	t.Data["json"] = resp
	t.ServeJSON()
}

// @Title Get all Task
// @Description Get all task
// @Success 200 {object} models.CronInfo
// @Failure 403 Forbidden
// @router /getall [get]
func (t *TaskController) GetALL() {
	ret := models.GetAllTask()
	t.Data["json"] = ret
	t.ServeJSON()
}

// @Title StopTask
// @Description stop task
// @Param	taskid		query 	string	true		"task id"
// @Success 200 {"status": 200}
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /stop [get]
func (t *TaskController) Stop() {
	resp := make(map[string]interface{})
	if taskid := t.GetString("taskid"); taskid != "" {
		logs.Info(t.Ctx.Input.IP() + " Stop Task " + taskid)
		models.StopTask(taskid)
		resp = map[string]interface{}{"status": 200}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	t.Data["json"] = resp
	t.ServeJSON()
}

// @Title StartTask
// @Description start task
// @Param	taskid		query 	string	true		"task id"
// @Success 200 {"status": 200}
// @Failure 10016 Miss required parameter
// @Failure 403 Forbidden
// @router /start [get]
func (t *TaskController) Start() {
	resp := make(map[string]interface{})
	if taskid := t.GetString("taskid"); taskid != "" {
		logs.Info(t.Ctx.Input.IP() + " Start Task " + taskid)
		models.StartTask(taskid)
		resp = map[string]interface{}{"status": 200}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	t.Data["json"] = resp
	t.ServeJSON()
}
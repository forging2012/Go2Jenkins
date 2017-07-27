package controllers

import (
	"devcloud/models"
	"strconv"
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
// @Failure 10017 Project already set task
// @Failure 10021 spec can not set all *
// @Failure 10022 0 <= m <= 59
// @Failure 10023 0 <= h < 24
// @Failure 403 Forbidden
// @router /add [get]
func (t *TaskController) Add() {
	project_name := t.GetString("project_name")
	task_name := models.MD5(time.Now().String())
	spec := t.GetString("spec")
	tasklist := t.GetString("tasklist")
	resp := make(map[string]interface{})
	if project_name != "" && spec != "" && tasklist != "" {
		if strings.Count(spec, "*") == 6 {
			resp = map[string]interface{}{"status": 10021, "error": "spec can not set all *"}
			goto finsh
		}
		specs := strings.Split(spec, " ")
		for n, s := range specs {
			if n == 1 {
				m, err := strconv.Atoi(s)
				if err != nil {
					resp = map[string]interface{}{"status": 10022, "error": "0 <= m <= 59"}
					goto finsh
				} else {
					if m < 0 || m > 59 {
						resp = map[string]interface{}{"status": 10022, "error": "0 <= m <= 59"}
						goto finsh
					}
				}
			}
			if n == 2 {
				h, err := strconv.Atoi(s)
				if err != nil {
					resp = map[string]interface{}{"status": 10023, "error": "0 <= h < 24"}
					goto finsh
				} else {
					if h < 0 || h >= 24 {
						resp = map[string]interface{}{"status": 10023, "error": "0 <= h < 24"}
						goto finsh
					}
				}
			}
		}
		if hasproject, _ := models.IsInTaskList(project_name); hasproject {
			resp = map[string]interface{}{"status": 10017, "error": project_name + " already set task"}
			goto finsh
		}
		tks := strings.Split(tasklist, ";")
		for _, tk := range tks {
			tk_name := strings.Split(tk, "|")[0]
			if tk_name == "compile" {
				num := len(strings.Split(tk, "|"))
				if num == 1 {
					resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
					goto finsh
				}
			}
			if tk_name == "pack" {
				num := len(strings.Split(tk, "|"))
				if num == 1 {
					resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
					goto finsh
				}
			}
		}
		logs.Info(t.Ctx.Input.IP() + " Add Task " + project_name + " " + task_name + " " + spec + " " + tasklist)
		models.AddTask(project_name, task_name, spec, tasklist, time.Now())
		resp = map[string]interface{}{"status": 200, "task_in_estype": "crontask", "task_in_esid": task_name}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
finsh:
	t.Data["json"] = resp
	t.ServeJSON()
}

// @Title DelTask
// @Description delete task
// @Param	taskid	query 	string	true		"task name"
// @Success 200 {"status": 200}
// @Failure 10016 Miss required parameter
// @Failure 10018 task in not exist
// @Failure 403 Forbidden
// @router /del [get]
func (t *TaskController) Del() {
	resp := make(map[string]interface{})
	if taskname := t.GetString("taskid"); taskname != "" {
		if isexit, _ := models.IsInTaskList(taskname); isexit {
			logs.Info(t.Ctx.Input.IP() + " Del Task " + taskname)
			models.DelTask(taskname)
			resp = map[string]interface{}{"status": 200}
		} else {
			resp = map[string]interface{}{"status": 10018, "error": "task in not exist"}
		}
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
// @Param	taskid		query 	string	true		"task name"
// @Success 200 {"status": 200}
// @Failure 10016 Miss required parameter
// @Failure 10018 task is not exist
// @Failure 10019 task is already stop
// @Failure 403 Forbidden
// @router /stop [get]
func (t *TaskController) Stop() {
	resp := make(map[string]interface{})
	if taskname := t.GetString("taskid"); taskname != "" {
		if isexist, status := models.IsInTaskList(taskname); isexist && status == "Y" {
			logs.Info(t.Ctx.Input.IP() + " Stop Task " + taskname)
			models.StopTask(taskname)
			resp = map[string]interface{}{"status": 200}
		} else if !isexist {
			resp = map[string]interface{}{"status": 10018, "error": "task in not exist"}
		} else if status == "N" {
			resp = map[string]interface{}{"status": 10019, "error": "task is already stop"}
		}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	t.Data["json"] = resp
	t.ServeJSON()
}

// @Title StartTask
// @Description start task
// @Param	taskid	query 	string	true		"task name"
// @Success 200 {"status": 200}
// @Failure 10016 Miss required parameter
// @Failure 10018 task is not exist
// @Failure 10020 task is already start
// @Failure 403 Forbidden
// @router /start [get]
func (t *TaskController) Start() {
	resp := make(map[string]interface{})
	if taskname := t.GetString("taskid"); taskname != "" {
		if isexist, status := models.IsInTaskList(taskname); isexist && status == "N" {
			logs.Info(t.Ctx.Input.IP() + " Start Task " + taskname)
			models.StartTask(taskname)
			resp = map[string]interface{}{"status": 200}
		} else if !isexist {
			resp = map[string]interface{}{"status": 10018, "error": "task is not exist"}
		} else if status == "Y" {
			resp = map[string]interface{}{"status": 10020, "error": "task is already start"}
		}
	} else {
		resp = map[string]interface{}{"status": 10016, "error": "Miss required parameter"}
	}
	t.Data["json"] = resp
	t.ServeJSON()
}

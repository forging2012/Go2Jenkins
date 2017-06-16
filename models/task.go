package models

import (
	"encoding/json"
	"fmt"
	//"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

const (
	EsCronConfig = "croninfo"
)

var BeegoTaskNameToTaskLists []map[string]string

//save [project+"|"+taskname]tasklist
var BeegoTaskNameToTaskList map[string]string

type CronInfo struct {
	Project  string `json:projectname`
	TaskName string `json:"taskname"`
	Spec     string `json:"spec"`
	TaskList string `json:"tasklist"`
}

func loadCronFromConfig() {
	var croninfos []CronInfo

	configfile_cron := beego.AppConfig.String("cron")
	if configfile_cron != "" {
		//json反序列化可以存储struct的切片中
		if err := json.Unmarshal([]byte(configfile_cron), &croninfos); err != nil {
			logs.Error("load cron from configfile has err" + err.Error())
		}
		for _, croninfo := range croninfos {
			addTask(croninfo)
			logs.Info("load cron from configfile")
			logs.Info(croninfo)
		}
	}
}

func loadCronFromEs() {
	var croninfos []CronInfo

	// es /devclouds_logs/croninfo/croninfo
	es_ret, err := search("EsCronConfig", "EsCronConfig")
	if err != nil {
		logs.Error("search task from es has err " + err.Error())
	} else {
		es_cron := es_ret.Source["msg"]
		if es_cron != nil {
			//json反序列化可以存储struct的切片中
			if err := json.Unmarshal([]byte(es_cron.(string)), &croninfos); err != nil {
				logs.Error("load cron from es has err" + err.Error())
			}
			for _, croninfo := range croninfos {
				addTask(croninfo)
				logs.Info("load cron from es")
				logs.Info(croninfo)
			}
		}
	}
}

//加载配置文件task
func addTask(ci CronInfo) {
	f := func() error { doFunc(ci.Project, ci.TaskName, ci.TaskList); return nil }
	beego_taskname := ci.Project + "-" + ci.TaskName
	tk := toolbox.NewTask(beego_taskname, ci.Spec, f)
	toolbox.AddTask(beego_taskname, tk)
	//每个任务的执行task列表存在TaskList
	BeegoTaskNameToTaskList = make(map[string]string)
	BeegoTaskNameToTaskList[beego_taskname] = ci.TaskList
	BeegoTaskNameToTaskLists = append(BeegoTaskNameToTaskLists, BeegoTaskNameToTaskList)
}

type CiResult struct {
	CHECKOUT  string
	CODECHECK string
	COMPILE   string
	PACK      string
}

func doFunc(project, taskname, tasklist string) {
	var ret CiResult
	tks := strings.Split(tasklist, ";")
	isexit := false
	for _, tk := range tks {
		switch {
		case tk == "checkout":
			result := GetCheckOutResult(project)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.CHECKOUT = result
		case tk == "codecheck":
			result := GetCodeCheckResult(project)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.CODECHECK = result
		case tk == "compile":
			result := GetCompileResult(project)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.COMPILE = result
		case tk == "pack":
			result := GetPackResult(project, "1.0", "N")["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.PACK = result
		}
		if isexit {
			break
			//fmt.Println(isexit)
		}
	}
	if rets, err := json.Marshal(&ret); err != nil {
		panic(err)
	} else {
		logs.Info(string(rets))
		// /devclouds_logs/crontask/$project-$taskname
		beego_taskname := project + "-" + taskname
		logs.Info(writeEs("crontask", beego_taskname, string(rets)))
	}
}

func AddTask(project, taskname, spec, tasklist string) {
	beego_taskname := project + "-" + taskname
	if beego_taskname != "monitor" {
		//f := func() error { fmt.Println(name + " task " + time.Now().Format("2006-01-02 15:04:05")); return nil }
		f := func() error { doFunc(project, taskname, tasklist); return nil }
		tk := toolbox.NewTask(beego_taskname, spec, f)
		toolbox.AddTask(beego_taskname, tk)
		tk.SetNext(time.Now())
		//每个任务的执行task列表存在TaskList
		BeegoTaskNameToTaskList = make(map[string]string)
		BeegoTaskNameToTaskList[beego_taskname] = tasklist
		//fmt.Println(BeegoTaskNameToTaskList)
		BeegoTaskNameToTaskLists = append(BeegoTaskNameToTaskLists, BeegoTaskNameToTaskList)
		//fmt.Println(BeegoTaskNameToTaskLists)
		logs.Info("add task:Project " + project + " TaskName " + taskname + " TaskList " + tasklist)
	}
}

func DelTask(taskname string) {
	if taskname != "monitor" {
		toolbox.DeleteTask(taskname)
		logs.Info("del task " + taskname)
	}
}

func DelayTask(taskname string) {
	admintasklist := toolbox.AdminTaskList
	for beego_taskname, tasker := range admintasklist {
		if beego_taskname == taskname {
			next := tasker.GetNext()
			fmt.Println("before", tasker.GetNext())
			mm, _ := time.ParseDuration("10m")
			mm1 := next.Add(mm)
			tasker.SetNext(mm1)
			fmt.Println("after", tasker.GetNext())
		}
	}
}

//增加监控，每10s执行一次
func addTask4Monitor() {
	f := func() error { monitor(); return nil }
	tk := toolbox.NewTask("monitor", "*/10 * * * * *", f)
	toolbox.AddTask("monitor", tk)
}

//监控所有task列表
//获得的任务列表会写入es
func monitor() {
	//var tasklist map[string]toolbox.Tasker
	admintasklist := toolbox.AdminTaskList
	var croninfos []*CronInfo
	for beego_taskname, tasker := range admintasklist {
		if beego_taskname != "monitor" {
			var ci *CronInfo
			project := strings.Split(beego_taskname, "-")[0]
			taskname := strings.Split(beego_taskname, "-")[1]
			for _, beegoTaskNameToTaskList := range BeegoTaskNameToTaskLists {
				//fmt.Println(beegoTaskNameToTaskList)
				tasklist := beegoTaskNameToTaskList[beego_taskname]
				if tasklist != "" {
					ci = &CronInfo{project, taskname, tasker.GetSpec(), tasklist}
				}
			}
			//fmt.Println(ci)
			croninfos = append(croninfos, ci)
		}
	}
	//切片序列化为json
	if cronconfig, err := json.Marshal(&croninfos); err != nil {
		panic(err)
	} else {
		logs.Info("croninfos: " + string(cronconfig))
		writeEs(EsCronConfig, EsCronConfig, string(cronconfig))
	}
}

func init() {
	//loadCronFromConfig()
	loadCronFromEs()
	addTask4Monitor()
}

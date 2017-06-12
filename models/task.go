package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func doFunc(name, tasklist string) {
	tk := strings.Split(tasklist, ";")
	isexit := false
	for _, taskname := range tk {
		//fmt.Println("taskname", taskname)
		switch {
		case taskname == "checkout":
			result := GetCheckOutResult(name)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			writeEs("checkout", result, "test")
		case taskname == "codecheck":
			result := GetCodeCheckResult(name)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			writeEs("codecheck", result, "test")
		case taskname == "compile":
			result := GetCompileResult(name)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			writeEs("compile", result, "test")
		case taskname == "pack":
			result := GetPackResult(name, "1.0", "N")["result"]
			if strings.Contains(result, "exit status") {
				writeEs("pack", result, "test")
			}
			writeEs("pack", result, "test")
		}
		if isexit {
			break
		}
	}
}

func AddTask(name, spec, tasklist string) {
	if name != "monitor" {
		//f := func() error { fmt.Println(name + " task " + time.Now().Format("2006-01-02 15:04:05")); return nil }
		f := func() error { doFunc(name, tasklist); return nil }
		tk := toolbox.NewTask(name, spec, f)
		toolbox.AddTask(name, tk)
		tk.SetNext(time.Now())
		//每个任务的执行task列表存在TaskList
		TaskList = make(map[string]string)
		TaskList[name] = tasklist
		logs.Info("add task " + name + " " + tasklist)
	}
}

func DelTask(name string) {
	if name != "monitor" {
		toolbox.DeleteTask(name)
		logs.Info("del task " + name)
	}
}

func DelayTask(name string) {
	tasklist := toolbox.AdminTaskList
	for taskname, tasker := range tasklist {
		if taskname == name {
			next := tasker.GetNext()
			fmt.Println("before", tasker.GetNext())
			mm, _ := time.ParseDuration("10m")
			mm1 := next.Add(mm)
			tasker.SetNext(mm1)
			fmt.Println("after", tasker.GetNext())
		}
	}
}

//加载配置文件task
func addTask(ci CronInfo) {
	f := func() error { doFunc(ci.Project, ci.TaskList); return nil }
	tk := toolbox.NewTask(ci.Project, ci.Spec, f)
	toolbox.AddTask(ci.Project, tk)
	//每个任务的执行task列表存在TaskList
	TaskList = make(map[string]string)
	TaskList[ci.Project] = ci.TaskList
}

//增加监控，每10s执行一次
func addTask4Monitor() {
	f := func() error { monitor(); return nil }
	tk := toolbox.NewTask("monitor", "*/10 * * * * *", f)
	toolbox.AddTask("monitor", tk)
}

//监控所有task列表
func monitor() {
	//var tasklist map[string]toolbox.Tasker
	tasklist := toolbox.AdminTaskList
	var croninfos []*CronInfo
	for taskname, tasker := range tasklist {
		ci := &CronInfo{taskname, tasker.GetSpec(), TaskList[taskname]}
		croninfos = append(croninfos, ci)
	}
	//切片序列化为json
	if cronconfig, err := json.Marshal(&croninfos); err != nil {
		panic(err)
	} else {
		logs.Info(string(cronconfig))
	}
}

var TaskList map[string]string

type CronInfo struct {
	Project  string `json:"project"`
	Spec     string `json:"spec"`
	TaskList string `json:"tasklist"`
}

func loadCron4Config() {
	cronconig := beego.AppConfig.String("cron")
	var croninfos []CronInfo
	if cronconig != "" {
		//json反序列化可以存储struct的切片中
		if err := json.Unmarshal([]byte(cronconig), &croninfos); err != nil {
			logs.Error("load cron from config has err" + err.Error())
		}
		for _, croninfo := range croninfos {
			addTask(croninfo)
			logs.Info("load cron from config")
			logs.Info(croninfo)
		}
	}
}

func init() {
	loadCron4Config()
	addTask4Monitor()
}

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
	EsCronInfo   = "croninfo"
	FileCronInfo = "./conf/croninfo"
)

var BeegoTaskNameToTaskLists []map[string]string

//save [project+"|"+taskname]tasklist
var BeegoTaskNameToTaskList map[string]string

func loadCronFromConfig() {
	var croninfos []CronInfo

	configfile_cron := beego.AppConfig.String("cron")
	if configfile_cron != "" {
		//json反序列化可以存储struct的切片中
		if err := json.Unmarshal([]byte(configfile_cron), &croninfos); err != nil {
			logs.Error("load cron from configfile has err" + err.Error())
		}
		for _, croninfo := range croninfos {
			logs.Info("load cron from configfile")
			AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList)
		}
	}
}
func loadCronFromEs() {
	var croninfos []CronInfo

	// es /devclouds_logs/croninfo/croninfo
	es_ret, err := search("EsCronInfo", "EsCronInfo")
	if err != nil {
		logs.Error("load task from es has err " + err.Error())
	} else {
		es_cron := es_ret.Source["msg"]
		if es_cron != nil {
			//json反序列化可以存储struct的切片中
			if err := json.Unmarshal([]byte(es_cron.(string)), &croninfos); err != nil {
				logs.Error("load cron from es has err" + err.Error())
			}
			for _, croninfo := range croninfos {
				logs.Info("load cron from es")
				AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList)
			}
		}
	}
}
func loadCronFromFile() {
	var croninfos []CronInfo
	file_ret, err := Read(FileCronInfo)
	if err != nil {
		logs.Error("load task from file has err " + err.Error())
	} else {
		if err := json.Unmarshal([]byte(file_ret), &croninfos); err != nil {
			logs.Error("load cron from file has err" + err.Error())
		}
		for _, croninfo := range croninfos {
			logs.Info("load cron from file")
			AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList)
		}
	}
}

type CiResult struct {
	CHECKOUT  string
	CODECHECK string
	COMPILE   string
	PACK      string
	Error     string
}

func doFunc(project, taskname, tasklist string) {
	var ret CiResult
	tks := strings.Split(tasklist, ";")
	isexit := false
	for _, tk := range tks {
		tk_name := strings.Split(tk, "|")[0]
		switch {
		case tk_name == "checkout":
			result := GetCheckOutResult(project)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.CHECKOUT = result
		case tk_name == "codecheck":
			result := GetCodeCheckResult(project)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.CODECHECK = result
		case tk_name == "compile":
			jdk_version := strings.Split(tk, "|")[1]
			result := GetCompileResult(project, jdk_version)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.COMPILE = result
		case tk_name == "pack":
			version := strings.Split(tk, "|")[1]
			result := GetPackResult(project, version, "N")["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.PACK = result
		default:
			ret.Error = "Project: " + project + " TaskName: " + taskname + " TaskList: " + tasklist + " has error"
			isexit = true
		}
		if isexit {
			break
			//fmt.Println(isexit)
		}
	}
	if rets, err := json.Marshal(&ret); err != nil {
		panic(err)
	} else {
		beego_taskname := project + "-" + taskname
		logs.Info("TaskResult: " + beego_taskname + " " + string(rets))
		// /devclouds_logs/crontask/$project-$taskname
		writeEs("crontask", beego_taskname, string(rets))
	}
}

//重启系统，所有定时任务的下一次时间重新定义
func AddTask(project, taskname, spec, tasklist string) {
	beego_taskname := project + "-" + taskname
	if beego_taskname != "monitor" {
		//f := func() error { fmt.Println(name + " task " + time.Now().Format("2006-01-02 15:04:05")); return nil }
		f := func() error { doFunc(project, taskname, tasklist); return nil }
		tk := toolbox.NewTask(beego_taskname, spec, f)
		toolbox.AddTask(beego_taskname, tk)
		tk.SetNext(time.Now())
		//每个任务的执行task列表存在BeegoTaskNameToTaskList
		BeegoTaskNameToTaskList = make(map[string]string)
		BeegoTaskNameToTaskList[beego_taskname] = tasklist
		BeegoTaskNameToTaskLists = append(BeegoTaskNameToTaskLists, BeegoTaskNameToTaskList)
		logs.Info("AddTask:Project " + project + " TaskName " + taskname + " Spec " + spec + " TaskList " + tasklist)
	}
}

func DelTask(taskname string) {
	if taskname != "monitor" {
		toolbox.DeleteTask(taskname)
		logs.Info("Del Task " + taskname)
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

type CronInfo struct {
	Project     string
	TaskName    string
	Spec        string
	TaskList    string
	PreRunTime  string
	NextRunTime string
}

//获取定时任务完整的信息，包括上一步执行时间和下一步执行时间
//但是在重启的时候，没有加载执行时间
func GetAllTask() []*CronInfo {
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
				tasklist, isexist := beegoTaskNameToTaskList[beego_taskname]
				if isexist {
					ci = &CronInfo{project, taskname, tasker.GetSpec(), tasklist, tasker.GetPrev().Format("2006-01-02 15:04:05"), tasker.GetNext().Format("2006-01-02 15:04:05")}
				}
			}
			//fmt.Println(ci)
			croninfos = append(croninfos, ci)
		}
	}
	return croninfos
}

//增加监控，每10s执行一次
func addTask4Monitor() {
	f := func() error { monitor(); return nil }
	tk := toolbox.NewTask("monitor", "*/10 * * * * *", f)
	toolbox.AddTask("monitor", tk)
}

//定时将任务信息写入到文件中
func monitor() {
	croninfo := GetAllTask()
	//切片序列化为json
	if croninfos, err := json.Marshal(&croninfo); err != nil {
		panic(err)
	} else {
		if string(croninfos) != "null" {
			logs.Info("CronInfo: " + string(croninfos))
		}
		//写入本地文件
		Write(FileCronInfo, string(croninfos))
		//写入es
		//writeEs(EsCronInfo, EsCronInfo, string(croninfos))
	}
}

func init() {
	beego.SetLogFuncCall(false)
	beego.BeeLogger.DelLogger("console")
	//logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":false}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/beego.log","level":7}`)
	//loadCronFromConfig()
	//loadCronFromEs()
	loadCronFromFile()
	addTask4Monitor()
}

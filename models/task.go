package models

import (
	"encoding/json"
	"fmt"
	"strconv"
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

/*
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
			AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList, time.Now())
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
				AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList, time.Now())

			}
		}
	}
}
*/
//重启系统时，如果配置文件里面有定时任务，不管怎么设置时间，都会以time.Now()重新开始
//为了让状态为N的任务继续保持为N，所以先添加在停止
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
			//只加载装态为Y的任务
			if croninfo.TaskStatus == "Y" {
				AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList, time.Now())
			}
			if croninfo.TaskStatus == "N" {
				//状态为N的，先添加执行时间是01/01/0001，然后在停止
				next, _ := time.Parse("01/02/2006", "01/01/0001")
				AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList, next)
				StopTask(croninfo.TaskName)
			}
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
			result := GetCheckOutResult(project, taskname)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.CHECKOUT = result
		case tk_name == "codecheck":
			result := GetCodeCheckResult(project, taskname)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.CODECHECK = result
		case tk_name == "compile":
			jdk_version := strings.Split(tk, "|")[1]
			result := GetCompileResult(project, jdk_version, taskname)["result"]
			if strings.Contains(result, "exit status") {
				isexit = true
			}
			ret.COMPILE = result
		case tk_name == "pack":
			version := strings.Split(tk, "|")[1]
			jdkversion := strings.Split(tk, "|")[2]
			result := GetPackResult(project, version, "N", jdkversion, taskname)["result"]
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
		logs.Info("TaskResult: " + project + " " + taskname + " " + string(rets))
		// /devclouds_logs/crontask/$taskname
		writeEs(Es_index, "crontask", taskname, map[string]interface{}{"msg": string(rets), "intime": time.Now().Format("2006-01-02 15:04:05")})
	}
}

//根据项目名/taskname查询是否存在
func IsInTaskList(condition string) (bool, string) {
	for _, croninfo := range CronInfos {
		if condition == croninfo.Project {
			return true, croninfo.TaskStatus
		}
	}

	for _, croninfo := range CronInfos {
		if condition == croninfo.TaskName {
			return true, croninfo.TaskStatus
		}
	}
	return false, ""
}

//校验定时任务spec的格式
func CheckSpec(spec string) (int, string) {
	specs := strings.Split(spec, " ")
	if len(specs) != 6 {
		return 10021, "spec has error,spec must be like (秒 分钟 小时 天 月 星期)"
	}
	if strings.Count(spec, "*") == 6 {
		return 10021, "spec has error,spec can not set all *"
	}
	for n, s := range specs {
		if n == 1 {
			m, err := strconv.Atoi(s)
			if err != nil {
				return 10021, "spec has error,minutes must be number"
			} else {
				if m < 0 || m > 59 {
					return 10021, "spec has error,0 <= minutes <= 59"
				}
			}
		}
		if n == 2 {
			h, err := strconv.Atoi(s)
			if err != nil {
				return 10021, "spec has error,hour must be number"
			} else {
				if h < 0 || h >= 24 {
					return 10021, "spec has error,0 <= hour < 24"
				}
			}
		}
		if n == 4 {
			if s != "*" {
				month, err := strconv.Atoi(s)
				if err != nil {
					return 10021, "spec has error,month must be number or *"
				} else {
					if month < 1 || month > 12 {
						return 10021, "spec has error,0 < mouth < 13"
					}
				}
			}
		}
	}
	return 0, ""

}

//重启系统，所有定时任务的下一次时间重新定义
//所有的task都会保存到CronInfos
func AddTask(project, taskname, spec, tasklist string, next time.Time) {
	if taskname != "monitor" {
		//f := func() error { fmt.Println(name + " task " + time.Now().Format("2006-01-02 15:04:05")); return nil }
		f := func() error { doFunc(project, taskname, tasklist); return nil }
		tk := toolbox.NewTask(taskname, spec, f)
		toolbox.AddTask(taskname, tk)
		//tk.SetNext(time.Now())
		tk.SetNext(next)
		isNew := true
		//如果添加的任务是之前停止的，则只是修改状态
		//如果是新增的任务，继续添加到CronInfos
		for _, croninfo := range CronInfos {
			if taskname == croninfo.TaskName {
				if taskname != "monitor" {
					croninfo.TaskStatus = "Y"
					croninfo.PreRunTime = toolbox.AdminTaskList[taskname].GetPrev().Format("2006-01-02 15:04:05")
					croninfo.NextRunTime = toolbox.AdminTaskList[taskname].GetNext().Format("2006-01-02 15:04:05")
				}
				isNew = false
			}
		}
		if isNew {
			preRunTime := toolbox.AdminTaskList[taskname].GetPrev().Format("2006-01-02 15:04:05")
			nextRunTime := toolbox.AdminTaskList[taskname].GetNext().Format("2006-01-02 15:04:05")
			ci := &CronInfo{project, taskname, spec, tasklist, preRunTime, nextRunTime, "Y"}
			//添加任务信息至CronInfos
			CronInfos = append(CronInfos, ci)
			if croninfos, err := json.Marshal(&ci); err != nil {
				panic(err)
			} else {
				logs.Info("AddTask: " + string(croninfos))
			}
		}
	}
}

//将task从taskadminlist中删除
//并从CronInfos中删除
func DelTask(task_name string) {
	var croninfos []*CronInfo

	logs.Info("Del Task " + task_name)
	StopTask(task_name)
	//利用一个新的切片，将删除的剔除
	for _, croninfo := range CronInfos {
		if task_name != croninfo.TaskName {
			croninfos = append(croninfos, croninfo)
		}
	}
	CronInfos = croninfos

}

//将task重新添加
//同时修改CronInfos中该task的状态
func StartTask(task_name string) {
	for _, croninfo := range CronInfos {
		if task_name == croninfo.TaskName {
			logs.Info("StartTask: " + croninfo.Project + " " + task_name)
			AddTask(croninfo.Project, croninfo.TaskName, croninfo.Spec, croninfo.TaskList, time.Now())
		}
	}
}

//停止任务首先删除保存在toolbox里面的任务，然后将全局变量CronInfos中该条任务的状态设为N
func StopTask(task_name string) {
	for _, croninfo := range CronInfos {
		if task_name == croninfo.TaskName {
			logs.Info("StopTask: " + croninfo.Project + " " + task_name)
			toolbox.DeleteTask(task_name)
			croninfo.TaskStatus = "N"
		}
	}
}

func DelayTask(task_name string) {
	admintasklist := toolbox.AdminTaskList
	for beego_taskname, tasker := range admintasklist {
		if beego_taskname == task_name {
			next := tasker.GetNext()
			fmt.Println("before", tasker.GetNext())
			mm, _ := time.ParseDuration("10m")
			mm1 := next.Add(mm)
			tasker.SetNext(mm1)
			fmt.Println("after", tasker.GetNext())
		}
	}
}

var CronInfos []*CronInfo

type CronInfo struct {
	Project     string
	TaskName    string
	Spec        string
	TaskList    string
	PreRunTime  string
	NextRunTime string
	TaskStatus  string
}

//根据beego_taskname来获取最新的上一次执行时间和下一次执行时间
func GetAllTask() []*CronInfo {
	admintasklist := toolbox.AdminTaskList
	for beego_taskname, tasker := range admintasklist {
		if beego_taskname != "monitor" {
			for _, croninfo := range CronInfos {
				//fmt.Println(reflect.TypeOf(croninfo.Project))
				if beego_taskname == croninfo.TaskName {
					croninfo.PreRunTime = tasker.GetPrev().Format("2006-01-02 15:04:05")
					croninfo.NextRunTime = tasker.GetNext().Format("2006-01-02 15:04:05")
				}
			}
		}
	}
	return CronInfos
}

//增加监控，每10s执行一次
func addTask4Monitor() {
	f := func() error { monitor(); return nil }
	tk := toolbox.NewTask("monitor", "*/10 * * * * *", f)
	toolbox.AddTask("monitor", tk)
}

//定时将任务信息写入到文件中,包括状态为Y和N
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
		//writeEs(Es_index, EsCronInfo, EsCronInfo, map[string]interface{}{"msg": string(croninfos), "intime": time.Now().Format("2006-01-02 15:04:05")})
	}
	//logs.Info(toolbox.AdminTaskList)
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

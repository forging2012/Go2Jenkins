package models

import (
	"encoding/json"
	"fmt"
	"time"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func AddTask(name,spec string) {
	f:= func() error { fmt.Println(name,time.Now()); return nil }
	tk := toolbox.NewTask(name, spec, f)
	toolbox.AddTask(name, tk)
	//time.Sleep(time.Second * 10)
	//toolbox.StopTask()
	//toolbox.StartTask()
}

func DelTask(name string) {
	toolbox.DeleteTask(name)
}

func DelayTask(name string) {
	tasklist := toolbox.AdminTaskList
	for taskname,tasker := range tasklist{
		if (taskname == "test1") {
			next := tasker.GetNext()
			fmt.Println("before",tasker.GetNext())
			mm, _ := time.ParseDuration("10m")  
			mm1 := next.Add(mm)
			tasker.SetNext(mm1)
			fmt.Println("after",tasker.GetNext())
		}
	}
}
//加载配置文件task
func addTask(ci CronInfo) {
	f:= func() error { fmt.Println(ci.Project,time.Now()); return nil }
	tk := toolbox.NewTask(ci.Project, ci.Spec, f)
	toolbox.AddTask(ci.Project, tk)
}
//监控所有task状态
func addTask4Monitor(){
	f:= func() error { monitor(); return nil }
	tk := toolbox.NewTask("monitor", "*/10 * * * * *", f)
	toolbox.AddTask("monitor", tk)
}

func monitor(){
	//var tasklist map[string]toolbox.Tasker
	tasklist := toolbox.AdminTaskList
	var croninfos []*CronInfo
	for taskname,tasker := range tasklist{
		ci := &CronInfo{taskname, tasker.GetSpec()}
		croninfos = append(croninfos,ci)
	}
	//切片序列化为json
	if cronconfig, err := json.Marshal(&croninfos); err != nil {
		panic(err)
	} else {
		fmt.Println(string(cronconfig))
	}
}

type CronInfo struct {
	Project string `json:"project"`
	Spec string `json:"spec"`
}

func loadCron4Config(){
	cronconig := beego.AppConfig.String("cron")
	var croninfos []CronInfo
	//json反序列化可以存储struct的切片中
	if err := json.Unmarshal([]byte(cronconig), &croninfos); err != nil {
		fmt.Println(err)
	}
	for _, croninfo := range croninfos {
		addTask(croninfo)
	}
}

func init(){
	loadCron4Config()
	addTask4Monitor()
}

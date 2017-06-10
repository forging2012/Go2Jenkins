package models

import (
	"encoding/json"
	"fmt"
	"time"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

func AddTask(name,spec string) {
	f:= func() error { fmt.Println(name,time.Now()); return nil }
    tk := toolbox.NewTask(name, spec, f)
    toolbox.AddTask(name, tk)
	tk.Run()
	toolbox.StopTask()
	toolbox.StartTask()
	logs.Debug(tk.GetStatus())
}

func addTask(name,spec string) {
	f:= func() error { fmt.Println(name,time.Now()); return nil }
    tk := toolbox.NewTask(name, spec, f)
    toolbox.AddTask(name, tk)
}

type CronInfo struct {
	Project string `json:"project"`
	Spec string `json:"spec"`
}

func LoadCron(){
	cronconig := beego.AppConfig.String("cron")
	var data []CronInfo
	if err := json.Unmarshal([]byte(cronconig), &data); err != nil {
            fmt.Println(err)
    }
	for _, b := range data {
            fmt.Println(b.Project, b.Spec)  //显示2组数据
            addTask(b.Project,b.Spec)
    }
}

func init(){
	LoadCron()
}

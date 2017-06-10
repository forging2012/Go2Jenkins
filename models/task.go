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
	toolbox.StopTask()
	toolbox.StartTask()
}

func DelTask(name string) {
    	toolbox.DeleteTask(name)
}

func addTask(ci CronInfo) {
	f:= func() error { fmt.Println(ci.Project,time.Now()); return nil }
    	tk := toolbox.NewTask(ci.Project, ci.Spec, f)
    	toolbox.AddTask(ci.Project, tk)
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
			addTask(b)
    	}
}

func init(){
	LoadCron()
}

package main

import (
	_ "devcloud/routers"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
)

var FilterIp = func(ctx *context.Context) {
	isIn := false
	isacl := beego.AppConfig.String("isacl")
	if isacl == "Y" {
		clientip := ctx.Input.IP()
		allow_ip := beego.AppConfig.String("allow_ip")
		s_allow_ip := strings.Split(allow_ip, ";")
		for _, ip := range s_allow_ip {
			if clientip == ip {
				isIn = true
			}
		}
		if !isIn {
			ctx.Redirect(403, "")
		}
	}
}

func main() {
	beego.SetLogFuncCall(false)
	beego.BeeLogger.DelLogger("console")
	//logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":false}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"./logs/beego.log","level":7}`)
	beego.InsertFilter("/v1/dc/*", beego.BeforeExec, FilterIp)
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	if beego.BConfig.RunMode == "prod" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/Snapshot"] = beego.AppConfig.String("filepath")
	}
	toolbox.StartTask()
	beego.Run()
}

package models

import (
	"os/exec"
	"strings"
	"regexp"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	
)

func GetCreateResult(project_name,svn_url string) string {
        var ret string
        out,err := exec.Command("/bin/bash", beego.AppConfig.String("create"),project_name,svn_url).Output()
        if err != nil {
                ret = string(out)
        } else {
		ret = "Create sucess "+project_name
        }
        return ret
}

func GetCheckOutResult(project_name string) string {
        var ret string
        out, err := exec.Command("/bin/bash", beego.AppConfig.String("checkout"),project_name).Output()
        if err != nil {
                ret = string(out)
        } else {
		ret = "Checkout sucess "+project_name
        }
        return ret
}

func GetCompileResult(project_name string) string {
        var ret string
        out, err := exec.Command("/bin/bash", beego.AppConfig.String("compile"),project_name).Output()
        if err != nil {
                ret = string(out)
        } else {
                ret = "Compile sucess "+project_name
        }
        return ret
}
	
func GetPackResult(project_name,version string) string {
	var ret string
        out, err := exec.Command("/bin/bash", beego.AppConfig.String("pack"),project_name,version).Output()
	logs.Debug(string(out))
        if err != nil {
                ret = string(out)
        } else {
		isok := strings.Contains(string(out),"ERROR")
		if(!isok){
			reg := regexp.MustCompile("PACKAGE_URL:.*[0-9]")
			PACKAGE_URL := strings.Replace(reg.FindAllString(string(out), -1)[0],"PACKAGE_URL:","",-1)
			ret = "Pack sucess "+PACKAGE_URL
		} else {
			ret = string(out)
		}
        }
        return ret
}

package models

import (
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type Result struct {
	ID      string
	_index  string
	_type   string
	result  string
	current string
	//next        string
	PACKAGE_URL string
	SONAR_URL   string
	es_result   string
}

var resp map[string]string

const (
	create    = "create.sh"
	checkout  = "checkout.sh"
	codecheck = "codecheck.sh"
	compile   = "compile.sh"
	pack      = "pack.sh"
)

var Es_index = beego.AppConfig.String("Index")

func GetCreateResult(project_name, svn_url string) (resp map[string]interface{}) {
	resp = make(map[string]interface{})
	var ret string
	out, err := exec.Command("/bin/bash", create, project_name, svn_url).Output()
	if err != nil {
		ret = string(out) + "|" + err.Error()
		resp["current"] = "N"
		//resp["next"] = "N"
	} else {
		ret = project_name + " Create sucess"
		resp["current"] = "Y"
		//resp["next"] = "Y"
	}
	//write es
	Index, Type, Id, err := writeEs(Es_index, "create", "", map[string]interface{}{"msg": ret, "intime": time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		//resp["result"] = ret
		//es报错
		resp["es_result"] = err.Error()
		//return
	}
	resp["PACKAGE_URL"] = ""
	resp["SONAR_URL"] = ""
	resp["_index"] = Index
	resp["_type"] = Type
	resp["ID"] = Id
	resp["result"] = ret

	return
}

func GetCheckOutResult(project_name string) (resp map[string]string) {
	resp = make(map[string]string)
	var ret string
	out, err := exec.Command("/bin/bash", checkout, project_name).Output()
	if err != nil {
		ret = string(out) + "|" + err.Error()
		resp["current"] = "N"
		//resp["next"] = "N"
	} else {
		ret = string(out)
		resp["current"] = "Y"
		//resp["next"] = "Y"
	}

	//write es
	Index, Type, Id, err := writeEs(Es_index, "checkout", "", map[string]interface{}{"msg": ret, "intime": time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		//resp["result"] = ret
		resp["es_result"] = err.Error()
		//return
	}
	resp["PACKAGE_URL"] = ""
	resp["SONAR_URL"] = ""
	resp["_index"] = Index
	resp["_type"] = Type
	resp["ID"] = Id
	resp["result"] = ret
	return
}

func GetCodeCheckResult(project_name string) (resp map[string]string) {
	resp = make(map[string]string)
	var ret string
	out, err := exec.Command("/bin/bash", codecheck, project_name).Output()
	if err != nil {
		ret = string(out) + "|" + err.Error()
		resp["current"] = "N"
		resp["SONAR_URL"] = ""
		//resp["next"] = "N"
	} else {
		reg := regexp.MustCompile("SONAR_URL:.*[a-z]")
		SONAR_URL := strings.Replace(reg.FindAllString(string(out), -1)[0], "SONAR_URL:", "", -1)
		resp["SONAR_URL"] = SONAR_URL
		ret = string(out)
		resp["current"] = "Y"
		//resp["next"] = "Y"
	}

	//write es
	Index, Type, Id, err := writeEs(Es_index, "codecheck", "", map[string]interface{}{"msg": ret, "intime": time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		//resp["result"] = ret
		resp["es_result"] = err.Error()
		//return
	}
	resp["PACKAGE_URL"] = ""
	resp["_index"] = Index
	resp["_type"] = Type
	resp["ID"] = Id
	resp["result"] = ret
	return
}

func GetCompileResult(project_name, jdk_version string) (resp map[string]string) {
	resp = make(map[string]string)
	var ret string
	out, err := exec.Command("/bin/bash", compile, project_name, jdk_version).Output()
	if err != nil {
		ret = string(out) + "|" + err.Error()
		resp["current"] = "N"
		//resp["next"] = "N"
	} else {
		//ret = "Compile sucess "+project_name
		ret = string(out)
		resp["current"] = "Y"
		//resp["next"] = "Y"
	}
	//write es
	Index, Type, Id, err := writeEs(Es_index, "compile", "", map[string]interface{}{"msg": ret, "intime": time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		//resp["result"] = ret
		resp["es_result"] = err.Error()
		//return
	}
	resp["PACKAGE_URL"] = ""
	resp["SONAR_URL"] = ""
	resp["_index"] = Index
	resp["_type"] = Type
	resp["ID"] = Id
	resp["result"] = ret
	return
}

func GetPackResult(project_name, version, isE string) (resp map[string]string) {
	resp = make(map[string]string)
	var ret string
	out, err := exec.Command("/bin/bash", pack, project_name, version, isE).Output()
	if err != nil {
		ret = string(out) + "|" + err.Error()
		resp["current"] = "N"
		resp["PACKAGE_URL"] = ""
		//resp["next"] = "N"
	} else {
		isok := strings.Contains(string(out), "ERROR")
		if !isok {
			ret = string(out)
			reg := regexp.MustCompile("PACKAGE_URL:.*[0-9]")
			PACKAGE_URL := strings.Replace(reg.FindAllString(string(out), -1)[0], "PACKAGE_URL:", "", -1)
			resp["PACKAGE_URL"] = PACKAGE_URL
			resp["current"] = "Y"
			//resp["next"] = "Y"
		} else {
			ret = string(out)
			resp["current"] = "N"
			resp["PACKAGE_URL"] = ""
			//resp["next"] = "N"
		}
	}

	//write es
	Index, Type, Id, err := writeEs(Es_index, "pack", "", map[string]interface{}{"msg": ret, "intime": time.Now().Format("2006-01-02 15:04:05")})
	if err != nil {
		//resp["result"] = ret
		resp["es_result"] = err.Error()
		//return
	}
	resp["SONAR_URL"] = ""
	resp["_index"] = Index
	resp["_type"] = Type
	resp["ID"] = Id
	resp["result"] = ret
	return
}

package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/url"
	"time"

	"github.com/OwnLocal/goes"
	"github.com/astaxie/beego"
)

func getClient() (conn *goes.Client) {
	conn = goes.NewClient(beego.AppConfig.String("ES_HOST"), beego.AppConfig.String("ES_PORT"))
	return
}

func writeEs(es_type, es_id string, ret interface{}) (string, string, string, error) {
	conn := getClient()
	var esid interface{}
	if es_id != "" {
		esid = es_id
	}
	d := goes.Document{
		Index: beego.AppConfig.String("Index"),
		Type:  es_type,
		ID:    esid,
		Fields: map[string]interface{}{
			"msg":    ret,
			"intime": time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	extraArgs := make(url.Values, 0)
	response, err := conn.Index(d, extraArgs)
	if err != nil {
		return "", "", "", err
	}
	return response.Index, response.Type, response.ID, nil
}

func search(es_type, es_id string) (ret *goes.Response, err error) {
	conn := getClient()
	es_index := beego.AppConfig.String("Index")
	extraArgs := make(url.Values, 0)
	ret, err = conn.Get(es_index, es_type, es_id, extraArgs)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func Write(filename, text string) bool {
	data := []byte(text)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		//panic(err)
		return false
	}
	return true
}

func Read(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

//生成一个大于30000小于65535的数
func randNum() int {
	min := 30000
	//max := 65535

SUIJI:
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(65534)
	if num < min {
		goto SUIJI
	}
	return num
}

func GetJsonFromMss(text map[string]string) (string, error) {
	if info, err := json.Marshal(&text); err != nil {
		return "", err
	} else {
		return string(info), nil
	}
}

func GetJsonFromMsi(text map[string]interface{}) (string, error) {
	if info, err := json.Marshal(&text); err != nil {
		return "", err
	} else {
		return string(info), nil
	}
}

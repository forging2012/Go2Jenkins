// @APIVersion 1.0.0
// @Title DevCloud API
// @Contact yaoyf@asiainfo.com
// @TermsOfServiceUrl http://10.11.20.102:6979/DevClouds
package routers

import (
	"devcloud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/dc",
			beego.NSInclude(
				&controllers.DevCloudController{},
			),
		),
		beego.NSNamespace("/index",
			beego.NSInclude(
				&controllers.IndexController{},
			),
		),
		beego.NSNamespace("/task",
			beego.NSInclude(
				&controllers.TaskController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

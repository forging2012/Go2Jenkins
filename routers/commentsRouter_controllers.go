package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"] = append(beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"],
		beego.ControllerComments{
			Method: "Create",
			Router: `/create`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"] = append(beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"],
		beego.ControllerComments{
			Method: "CheckOut",
			Router: `/checkout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"] = append(beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"],
		beego.ControllerComments{
			Method: "CodeCheck",
			Router: `/codecheck`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"] = append(beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"],
		beego.ControllerComments{
			Method: "Compile",
			Router: `/compile`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"] = append(beego.GlobalControllerRouter["devcloud/controllers:DevCloudController"],
		beego.ControllerComments{
			Method: "Pack",
			Router: `/pack`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:IndexController"] = append(beego.GlobalControllerRouter["devcloud/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:SshController"] = append(beego.GlobalControllerRouter["devcloud/controllers:SshController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:TaskController"] = append(beego.GlobalControllerRouter["devcloud/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:TaskController"] = append(beego.GlobalControllerRouter["devcloud/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Del",
			Router: `/del`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:TaskController"] = append(beego.GlobalControllerRouter["devcloud/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetALL",
			Router: `/getall`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:TaskController"] = append(beego.GlobalControllerRouter["devcloud/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Stop",
			Router: `/stop`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:TaskController"] = append(beego.GlobalControllerRouter["devcloud/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Start",
			Router: `/start`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}

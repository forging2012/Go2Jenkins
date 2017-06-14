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

	beego.GlobalControllerRouter["devcloud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["devcloud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["devcloud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["devcloud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["devcloud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:ObjectController"] = append(beego.GlobalControllerRouter["devcloud/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
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

	beego.GlobalControllerRouter["devcloud/controllers:UserController"] = append(beego.GlobalControllerRouter["devcloud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:UserController"] = append(beego.GlobalControllerRouter["devcloud/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:UserController"] = append(beego.GlobalControllerRouter["devcloud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:UserController"] = append(beego.GlobalControllerRouter["devcloud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:UserController"] = append(beego.GlobalControllerRouter["devcloud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:UserController"] = append(beego.GlobalControllerRouter["devcloud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["devcloud/controllers:UserController"] = append(beego.GlobalControllerRouter["devcloud/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}

package routers

import (
	"github.com/oikomi/FishChatServer/monitor/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.SetStaticPath("/doc", "doc")
	beego.AddNamespace(ns)
}

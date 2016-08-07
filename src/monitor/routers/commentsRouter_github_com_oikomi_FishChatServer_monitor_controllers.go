package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"],
		beego.ControllerComments{
			"Get",
			`/:objectId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"],
		beego.ControllerComments{
			"Put",
			`/:objectId`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["github.com/oikomi/FishChatServer/monitor/controllers:ObjectController"],
		beego.ControllerComments{
			"Delete",
			`/:objectId`,
			[]string{"delete"},
			nil})

}

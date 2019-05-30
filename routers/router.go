// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ss-door/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/door",
			beego.NSRouter("/add",&controllers.DoorController{},"get:Add"),
			beego.NSRouter("/remove", &controllers.DoorController{}, "get:Remove"),
			beego.NSRouter("/list", &controllers.DoorController{}, "get:List"),

		),
	)
	beego.AddNamespace(ns)
}

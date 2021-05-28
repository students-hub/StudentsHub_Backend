// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"StudentsHub_Backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns1 := beego.NewNamespace("/course",
		beego.NSInclude(
			&controllers.CourseController{},
		),
	)
	ns2 := beego.NewNamespace("/user",
		beego.NSInclude(
			&controllers.UserController{},
		),
	)
	beego.AddNamespace(ns1, ns2)
}

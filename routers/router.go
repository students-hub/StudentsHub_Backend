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

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	beego.Router("/User/Login", &controllers.UserController{}, "get:Login")
	beego.Router("/User/Logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/User/Signup", &controllers.UserController{}, "get:AddUser")
}

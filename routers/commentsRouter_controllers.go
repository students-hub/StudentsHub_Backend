package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:CourseController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:CourseController"],
        beego.ControllerComments{
            Method: "AddCourse",
            Router: "/add",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:CourseController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:CourseController"],
        beego.ControllerComments{
            Method: "DeleteCourse",
            Router: "/delete",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/delete",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "AddUser",
            Router: "/sign-up",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUsername",
            Router: "/updatename",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdatePassword",
            Router: "/updatepswd",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"] = append(beego.GlobalControllerRouter["StudentsHub_Backend/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateRole",
            Router: "/updaterole",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}

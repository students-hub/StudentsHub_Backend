package main

import (
	_ "StudentsHub_Backend/routers"

	"github.com/astaxie/beego"
)

func main() {
	//fmt.Println(beego.AppPath)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

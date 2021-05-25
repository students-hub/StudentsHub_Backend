package main

import (
	_ "StudentsHub_Backend/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//fmt.Println(beego.AppPath)

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

package main

import (
	"github.com/astaxie/beego"
	_ "ss-door/routers"
	"ss-door/service"
	_ "ss-door/service"
	)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	service.Init()
	go func() {
		service.State()
	}()
	defer service.Release()
	beego.Run()
}

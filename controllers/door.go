package controllers

import (
	"github.com/astaxie/beego"
	"ss-door/service"
)

type DoorController struct{
	beego.Controller
}



func (door *DoorController) Add(){

	service.Add(8000,"12345")
	door.ServeJSONP()
}


func (door *DoorController) Remove(){
	service.Remove(8384)
	door.ServeJSONP()
}


func (door *DoorController) List(){
	service.List()
	door.ServeJSONP()
}

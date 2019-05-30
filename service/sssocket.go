package service

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"net"
)

var Client *net.UnixConn
var bufferSize = 1024
var addr *net.UnixAddr
func Init() {

	var err error
	socketPath := beego.AppConfig.String("socket")
	addr, err = net.ResolveUnixAddr("unixgram", socketPath)

	if err != nil {
		beego.Error(" cannot relove unix addr:", addr)
		panic(err)
	}

	Client, err = net.DialUnix("unixgram", nil, addr)

	if err != nil {
		beego.Error(" build client error")
		panic(err)
	}

	beego.Info(Client)

}

type PortInfo struct {
	ServerPort int32  `json:"server_port"`
	Password   string `json:"password"`
}

func Add(port int32, password string) {

	portInfo := PortInfo{
		ServerPort: port,
		Password:   password,
	}

	bytes, err := json.Marshal(portInfo)

	if err != nil {
		beego.Error("reslove data fail", portInfo)
		return
	}
	body := string(bytes)
	command := "add:" + body
	Client.Write([]byte(command))

}

func Remove(port int32) {
	target := make(map[string]int32)
	target["server_port"] = port
	bytes, err := json.Marshal(target)
	if err != nil {
		beego.Error("reslove data fail", target)
		return
	}

	body := string(bytes)
	command := "remove:" + body
	Client.Write([]byte(command))
}
func List() {

	body := "ping"

	Client.Write([]byte(body))

}

func State() {
	buffer := make([]byte, bufferSize)
	for {
		beego.Info("Wait Output")
		nr,remote, _ := Client.ReadFromUnix(buffer)
		beego.Info("state:", string(buffer[0:nr]), remote)
	}

}

func Release(){

	if Client!=nil{
		Client.Close()
	}

}

package main

import (
	//_ "api-gateway/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}

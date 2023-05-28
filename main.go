package main

import (
	_ "Fashion-Freaks/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}


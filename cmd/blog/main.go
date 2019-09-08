package main

import (
	"github.com/astaxie/beego"
	_ "go_mod-project/cmd/blog/routers"
	"go_mod-project/cmd/blog/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

package main

import (
	"rbacAdmin/core"
	"rbacAdmin/flags"
	"rbacAdmin/global"
)

func main() {
	core.InitLogger("logs")
	global.Config = core.ReadConfig()
	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()
	flags.Run()
}

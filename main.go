package main

import (
	"rbacAdmin/core"
	"rbacAdmin/flags"
	"rbacAdmin/global"
	"rbacAdmin/utils/captcha"
	route "rbacAdmin/routes"
)

func main() {
	core.InitLogger("logs")           // 初始化日志
	global.Config = core.ReadConfig() // 读取配置
	global.DB = core.InitGorm()       // 初始化数据库
	global.Casbin = core.InitCasbin() // 初始化casbin
	global.Redis = core.InitRedis()   // 初始化redis

	// 启动邮件验证码清理定时器
	captcha.EmailStore.StartCleanupTimer()

	flags.Run() // 运行应用
	route.Run() // 运行路由
}

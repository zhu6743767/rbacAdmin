package routes

import (
	"rbacAdmin/global"

	gin "github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
)

func Run() {
	s := global.Config.System

	gin.SetMode(s.Mode)

	// 调试模式下开启 gin 调试模式
	if s.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	// 非调试模式下关闭 gin 调试模式
	if s.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	// router
	r := gin.Default()

	// 设置 /uploads 目录为静态文件目录
	r.Static("/uploads", "./uploads")

	// api 路由组
	g := r.Group("api")

	// 中间件
	//g.Use()

	// 用户组路由
	UserRouter(g)

	// 验证码路由
	CaptchaRouter(g)

	// 运行路由
	logrus.Infof("后端服务运行在 %s", s.Addr())
	r.Run(s.Addr())
}

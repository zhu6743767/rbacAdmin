package route

import (
	gin "github.com/gin-gonic/gin"

	"rbacAdmin/global"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	if global.Config.System.Mode == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.Default()
	r.Run(global.Config.System.Addr())
}

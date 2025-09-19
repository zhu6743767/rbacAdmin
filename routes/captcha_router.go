package routes

import (
	"rbacAdmin/api/captcha_api"

	gin "github.com/gin-gonic/gin"
)

func CaptchaRouter(r *gin.RouterGroup) {
	g := r.Group("")
	app := captcha_api.CaptchaApi{}
	g.GET("captcha", app.GenerateCaptchaView)
}

package routes

import (
	"rbacAdmin/api/email_api"
	"rbacAdmin/middleware"

	gin "github.com/gin-gonic/gin"
)

func EmailRouter(r *gin.RouterGroup) {
	g := r.Group("")
	app := email_api.EmailApi{}
	g.POST("email/send_email", middleware.BindJson[email_api.SendEmailRequest], app.SendEmailView)
}

package routes

import (
	"rbacAdmin/api"
	"rbacAdmin/api/user_api"
	"rbacAdmin/middleware"

	gin "github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	g := r.Group("")
	app := api.App.UserApi
	g.POST("login", middleware.BindJson[user_api.LogingRequest], (app.LoginView))
}

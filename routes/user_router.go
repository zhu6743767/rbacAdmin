package routes

import (
	"rbacAdmin/api"

	gin "github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	g := r.Group("")
	app := api.App.UserApi
	g.GET("login", app.LoginView)
}

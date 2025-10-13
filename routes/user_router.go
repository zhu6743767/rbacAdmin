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
	g.POST("register", middleware.BindJson[user_api.RegisterRequest], (app.RegisterView))
	g.PUT("user/password", middleware.AuthMiddleware, middleware.BindJson[user_api.UpdatePasswordRequest], app.UpdatePasswordView)
	g.PUT("users", middleware.AuthMiddleware, middleware.BindJson[user_api.UpdateUserInfoRequest], app.UpdateUserInfoView)
	g.GET("users/info", middleware.AuthMiddleware, app.UserInfoView)
	g.GET("users/list", middleware.AuthMiddleware, middleware.BindQuery[user_api.UserListRequest], app.UserListView)
}

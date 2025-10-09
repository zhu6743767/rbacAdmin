package routes

import (
	"rbacAdmin/api"
	"rbacAdmin/middleware"

	gin "github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	g := r.Group("image").Use(middleware.AuthMiddleware)
	app := api.App.ImageApi
	g.POST("upload", app.UploadView)
}

package user_api

import "github.com/gin-gonic/gin"

func (u *UserApi) LoginView(c *gin.Context) {
	c.String(200, "login")
}

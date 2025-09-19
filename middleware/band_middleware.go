package middleware

import (
	"rbacAdmin/common/resp"

	"github.com/gin-gonic/gin"
)

func BindJson[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		//c.JSON(200, gin.H{"code": 1001, "msg": err.Error(), "data": nil})
		resp.FailWithBindingError(err, c)
		c.Abort() // 直接拦截并返回
		return
	}

	c.Set("request", cr)
	return
}

func BindQuery[T any](c *gin.Context) {
	var cr T
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		//c.JSON(200, gin.H{"code": 1001, "msg": err.Error(), "data": nil})
		resp.FailWithBindingError(err, c)
		c.Abort() // 直接拦截并返回
		return
	}

	c.Set("request", cr)
	return
}

func GetBind[T any](c *gin.Context) (data T) {
	return c.MustGet("request").(T)
}

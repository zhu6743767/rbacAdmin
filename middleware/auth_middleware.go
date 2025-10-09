package middleware

import (
	"rbacAdmin/common/resp"
	"rbacAdmin/service/redis_service/token_black"
	"rbacAdmin/utils/jwts"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	claims, err := jwts.ParseToken(token)
	if err != nil {
		resp.FailWithMsg("请先登录", c)
		c.Abort()
		return
	}
	// 判断这个token是否在黑名单中
	if token_black.HaveToken(token) {
		resp.FailWithMsg("该用户已注销", c)
		return
	}
	c.Set("claims", claims)
	return

}

func GetAuth(c *gin.Context) jwts.ClaimsUserInfo {
	claims := c.MustGet("claims").(*jwts.Claims)
	return claims.ClaimsUserInfo
}

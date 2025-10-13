package user_api

import (
	"rbacAdmin/common/resp"
	"rbacAdmin/global"
	"rbacAdmin/middleware"
	"rbacAdmin/models"

	"github.com/gin-gonic/gin"
)

type UpdateUserInfoRequest struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

func (u *UserApi) UpdateUserInfoView(c *gin.Context) {

	cr := middleware.GetBind[UpdateUserInfoRequest](c)
	claims := middleware.GetAuth(c)

	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		resp.FailWithMsg("用户不存在", c)
		c.Abort()
		return
	}

	err = global.DB.Model(&user).Updates(models.UserModel{
		Nickname: cr.Nickname,
		Avatar:   cr.Avatar,
	}).Error
	if err != nil {
		resp.FailWithMsg("修改用户信息失败", c)
		c.Abort()
		return
	}

	resp.OkWithMsg("修改用户信息成功", c)
}

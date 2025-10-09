package user_api

import (
	"rbacAdmin/common/resp"
	"rbacAdmin/global"
	"rbacAdmin/middleware"
	"rbacAdmin/models"

	"github.com/gin-gonic/gin"
)

type UserInfoResponse struct {
	UserID   uint   `json:"user_id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	RoleList []uint `json:"roleList"`
}

func (UserApi) UserInfoView(c *gin.Context) {
	claims := middleware.GetAuth(c)

	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		resp.FailWithMsg("用户不存在", c)
		c.Abort()
		return
	}

	data := UserInfoResponse{
		UserID:   user.ID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		RoleList: user.GetRoleList(),
	}

	resp.OkWithData(data, c)
}

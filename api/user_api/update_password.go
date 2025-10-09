package user_api

import (
	"rbacAdmin/common/resp"
	"rbacAdmin/global"
	"rbacAdmin/middleware"
	"rbacAdmin/models"
	"rbacAdmin/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required"`
	Pwd    string `json:"pwd" binding:"required,min=6,max=64"`
	RePwd  string `json:"re_pwd" binding:"required,min=6,max=64"`
}

func (u *UserApi) UpdatePasswordView(c *gin.Context) {

	cr := middleware.GetBind[UpdatePasswordRequest](c)

	claims := middleware.GetAuth(c)

	var user models.UserModel
	err := global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		resp.FailWithMsg("用户不存在", c)
		c.Abort()
		return
	}
	if !pwd.ComparePassword(user.Password, cr.OldPwd) {
		resp.FailWithMsg("原密码错误", c)
		c.Abort()
		return
	}
	if cr.Pwd != cr.RePwd {
		resp.FailWithMsg("两次密码不一致", c)
		c.Abort()
		return
	}

	hashPwd := pwd.HashedPassword(cr.Pwd)
	global.DB.Model(&user).Update("password", hashPwd)

	resp.OkWithMsg("修改密码成功", c)
}

package user_api

import (
	"rbacAdmin/common/resp"
	"rbacAdmin/global"
	"rbacAdmin/middleware"
	"rbacAdmin/models"
	"rbacAdmin/utils/captcha"
	"rbacAdmin/utils/pwd"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RegisterRequest struct {
	Email      string `json:"email" binding:"required,email"`
	EmailID    string `json:"emailID" binding:"required"`
	EmailCode  string `json:"emailCode" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required"`
}

func (u *UserApi) RegisterView(c *gin.Context) {
	cr := middleware.GetBind[RegisterRequest](c)
	// 验证邮件验证码
	if !captcha.CaptchaStore.Verify(cr.EmailID, cr.EmailCode, false) {
		resp.FailWithMsg("邮件验证码错误", c)
		return
	}
	// 校验密码是否一致
	if cr.Password != cr.RePassword {
		resp.FailWithMsg("两次密码不一致", c)
		return
	}
	// 校验邮箱是否存在
	var user models.UserModel
	err := global.DB.Take(&user, "email = ?", cr.Email).Error
	if err == nil {
		resp.FailWithMsg("该邮箱已注册", c)
		return
	}
	// 密码加密
	pwdHash := pwd.HashedPassword(cr.RePassword)

	// 注册用户
	err = global.DB.Create(&models.UserModel{
		Username: cr.Email,
		Email:    cr.Email,
		Password: pwdHash,
	}).Error
	if err != nil {
		resp.FailWithMsg("注册失败", c)
		return
	}
	logrus.Info("注册用户: %s", cr.Email)
	resp.OkWithData(cr.Email, c)

}

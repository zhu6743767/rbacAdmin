package email_api

import (
	"fmt"
	"rbacAdmin/common/resp"
	"rbacAdmin/global"
	"rbacAdmin/middleware"
	"rbacAdmin/utils/captcha"
	"rbacAdmin/utils/email"
	"rbacAdmin/utils/random"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type EmailApi struct {
}

type SendEmailRequest struct {
	Email       string `json:"email" binding:"required,email"`
	CaptchaID   string `json:"captchaID"`
	CaptchaCode string `json:"captchaCode"`
}

type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (e *EmailApi) SendEmailView(c *gin.Context) {
	cr := middleware.GetBind[SendEmailRequest](c)
	if !global.Config.Email.Verify() {
		resp.FailWithMsg("管理员尚未配置邮箱，暂时无法发送邮箱验证码", c)
		return
	}
	if global.Config.Captcha.Enable {
		if cr.CaptchaID == "" || cr.CaptchaCode == "" {
			resp.FailWithMsg("请输入验证码", c)
			return
		}
		// 验证码校验
		if !captcha.CaptchaStore.Verify(cr.CaptchaID, cr.CaptchaCode, true) {
			resp.FailWithMsg("验证码错误", c)
			return
		}
	}

	emailID := uuid.New().String()
	code := random.RandStrByCode("1234567890", 6)

	email.Set(emailID, cr.Email, code)

	content := fmt.Sprintf("您正在完成用户注册， 您的验证码为 %s , 请在5分钟内容使用，过时无效！", code)
	err := email.SendEmail(cr.Email, "用户注册", content)
	if err != nil {
		logrus.Error("发送验证码邮件失败: ", err)
		resp.FailWithMsg("发送验证码邮件失败", c)
		return
	}

	// 记录日志
	logrus.Info("验证码邮件发送成功: ", cr.Email)
	resp.OkWithData(SendEmailResponse{
		EmailID: emailID, // 使用邮箱作为ID，用于后续验证
	}, c)
}

package captcha_api

import (
	"rbacAdmin/common/resp"
	"rbacAdmin/utils/captcha"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type CaptchaApi struct {
}

type GenerateCaptchaResponse struct {
	CaptchaID string `json:"captchaID"`
	Captcha   string `json:"captcha"`
}

func (CaptchaApi) GenerateCaptchaView(c *gin.Context) {
	var driver = base64Captcha.DriverString{
		Height:          200,
		Width:           60,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "1234567890",
	}
	cp := base64Captcha.NewCaptcha(&driver, captcha.CaptchaStore)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		logrus.Error("生成验证码失败: ", err)
		resp.FailWithMsg("验证码生成失败", c)
		return
	}
	resp.OkWithData(GenerateCaptchaResponse{
		CaptchaID: id,
		Captcha:   b64s,
	}, c)
}

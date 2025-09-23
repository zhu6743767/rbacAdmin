package captcha_api

import (
	"fmt"
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
		Width:           600,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "1234567890",
	}
	cp := base64Captcha.NewCaptcha(&driver, captcha.CaptchaStore)
	id, b64s, amswer, err := cp.Generate()
	if err != nil {
		logrus.Error("生成验证码失败: ", err)
		resp.FailWithMsg("验证码生成失败", c)
		return
	}
	resp.OkWithData(GenerateCaptchaResponse{
		CaptchaID: id,
		Captcha:   b64s,
	}, c)
	fmt.Println("验证码ID:", id)
	fmt.Println("验证码：", amswer)
	//fmt.Println("验证码base64:", b64s)
	//fmt.Println("验证码图片:", ConvertB64sToImage(b64s, "./uploads/1.jpg"))
}

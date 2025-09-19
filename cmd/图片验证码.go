package main

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

var CaptchaStore = base64Captcha.DefaultMemStore

func main() {
	var driver = base64Captcha.DriverString{
		Height:          200,
		Width:           60,
		NoiseCount:      2,
		ShowLineOptions: 4,
		Length:          6,
		Source:          "1234567890",
	}
	cp := base64Captcha.NewCaptcha(&driver, CaptchaStore)
	id, b64s, answer, err := cp.Generate()
	if err != nil {
		logrus.Error("生成验证码失败: ", err)
		return
	}
	fmt.Println(id)
	fmt.Println(b64s)
	fmt.Println(answer)
}

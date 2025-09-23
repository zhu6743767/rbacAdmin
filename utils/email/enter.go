package email

import (
	"crypto/tls"
	"rbacAdmin/global"

	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

func SendEmail(user string, subject string, content string) error {
	e := global.Config.Email
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", e.User)     // 发送者
	mailer.SetHeader("To", user)         // 接收者
	mailer.SetHeader("Subject", subject) // 邮件主题
	mailer.SetBody("text/html", content) // 邮件内容
	// 构建SMTP客户端
	dialer := gomail.NewDialer(e.Host, e.Port, e.User, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 忽略证书校验
	if err := dialer.DialAndSend(mailer); err != nil {
		logrus.Errorf("发送邮件失败: %v", err)
		return err
	}
	return nil

}

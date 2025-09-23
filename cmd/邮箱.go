package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"rbacAdmin/core"
	"rbacAdmin/global"

	gomail "gopkg.in/gomail.v2"
)

func main() {
	// 初始化配置
	global.Config = core.ReadConfig()
	
	// 检查邮箱配置
	if global.Config.Email.User == "" {
		log.Fatal("邮箱配置未设置，请检查 settings_dev.yaml 中的 email 配置")
	}
	
	// 设置SMTP服务器
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", global.Config.Email.User)
	mailer.SetHeader("To", "315420498@qq.com")
	mailer.SetHeader("Subject", "测试邮件")
	mailer.SetBody("text/html", "这是一封测试邮件")

	// QQ邮箱的SMTP服务器配置
	// 主机: smtp.qq.com
	// 端口: 465
	// 用户名: 你的QQ邮箱地址
	// 密码: 你在QQ邮箱设置中生成的授权码

	// smtpHost := "smtp.qq.com"
	// smtpPort := 465
	// smtpUser := global.Config.Email.User
	// smtpPassword := global.Config.Email.Password

	e := global.Config.Email

	// 构建SMTP客户端
	dialer := gomail.NewDialer(e.Host, e.Port, e.User, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 忽略证书校验

	// 发送邮件
	if err := dialer.DialAndSend(mailer); err != nil {
		log.Fatalf("发送邮件失败: %v", err)
		return
	}
	fmt.Println("邮件发送成功")
	log.Println("邮件发送成功")

}

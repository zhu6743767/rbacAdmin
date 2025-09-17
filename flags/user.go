package flags

import (
	"fmt"
	"os"
	"rbacAdmin/global"
	"rbacAdmin/models"
	"rbacAdmin/utils/pwd"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

type User struct {
}

func (User) Create() {
	// 创建用户
	fmt.Println("请输入用户名")
	var username string
	fmt.Scanln(&username)
	var user models.UserModel
	err := global.DB.Take(&user, "username = ?", username).Error
	if err == nil {
		fmt.Println("用户已存在")
		logrus.Error("用户已存在")
		return
	}
	fmt.Println("请输入密码")
	password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码失败")
		logrus.Error("读取密码失败", err)
		return
	}
	fmt.Println("请再次输入密码")
	rePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("读取密码失败")
		logrus.Error("读取密码失败", err)
		return
	}
	if string(password) != string(rePassword) {
		fmt.Println("两次密码不一致")
		logrus.Error("两次密码不一致")
		return
	}
	// 密码加密
	hashPwd := pwd.HashedPassword(string(password))
	err = global.DB.Create(&models.UserModel{
		Username: username,
		Password: hashPwd,
		IsAdmin:  true,
	}).Error
	if err != nil {
		fmt.Println("创建用户失败")
		logrus.Error("创建用户失败", err)
		return
	}
	fmt.Println("创建用户成功")
	logrus.Info("创建用户成功")

}

package main

// 密码测试
import (
	"fmt"

	"rbacAdmin/utils/pwd"
)

func main() {
	hashPwd := pwd.HashedPassword("123456")
	fmt.Println("密码哈希:", hashPwd)
	// 密码校验
	ok := pwd.ComparePassword(hashPwd, "123456")
	fmt.Println("密码验证结果:", ok)
}

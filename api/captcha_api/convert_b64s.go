package captcha_api

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"strings"
)

// ConvertB64sToImage 将base64编码的验证码图片字符串转换为图片文件并保存
func ConvertB64sToImage(b64s string, outputPath string) error {
	// base64图片字符串通常格式为: data:image/png;base64,xxxxxx
	// 我们需要先提取实际的base64编码数据
	const prefix = "data:image/" 
	const base64Prefix = ";base64,"
	
	// 查找base64前缀的位置
	base64Start := strings.Index(b64s, base64Prefix)
	if base64Start == -1 {
		// 如果没有找到标准前缀，可能是纯base64数据
		// 直接尝试解码
	} else {
		// 提取base64数据部分
		dataPart := b64s[base64Start+len(base64Prefix):]
		b64s = dataPart
	}
	
	// 解码base64数据
	imageData, err := base64.StdEncoding.DecodeString(b64s)
	if err != nil {
		return fmt.Errorf("base64解码失败: %v", err)
	}
	
	// 保存为图片文件
	if err := ioutil.WriteFile(outputPath, imageData, 0644); err != nil {
		return fmt.Errorf("图片保存失败: %v", err)
	}
	
	return nil
}

// FormatB64sForDisplay 将base64字符串格式化为更易读的形式（适用于调试）
func FormatB64sForDisplay(b64s string) string {
	// 显示前100个字符和后100个字符，中间用省略号
	if len(b64s) > 200 {
		return b64s[:100] + "...[中间省略]..." + b64s[len(b64s)-100:]
	}
	return b64s
}
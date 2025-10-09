package image_api

import (
	"fmt"
	"os"
	"path"
	"rbacAdmin/common/resp"
	"rbacAdmin/global"
	"rbacAdmin/middleware"
	"rbacAdmin/utils/md5"
	"rbacAdmin/utils/random"
	"strings"

	"github.com/gin-gonic/gin"
)

var WhiteMap = map[string]bool{
	".jpg":  true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

func (u *ImageApi) UploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		resp.FailWithMsg("获取文件失败", c)
		c.Abort()
		return
	}
	// 验证文件大小
	if fileHeader.Size > 1024*1024*global.Config.Upload.Size {
		resp.FailWithMsg(fmt.Sprintf("文件大小不能超过 %d MB", global.Config.Upload.Size), c)
		c.Abort()
		return
	}
	// 验证文件类型，以防木马类文件上传
	_, ok := WhiteMap[strings.ToLower(path.Ext(fileHeader.Filename))]
	if !ok {
		resp.FailWithMsg("文件类型不支持", c)
		c.Abort()
		return
	}
	// 文件名重复，但是内容不同，不覆盖

	auth := middleware.GetAuth(c)

	dst := path.Join("uploads", global.Config.Upload.Dir, auth.Username, fileHeader.Filename)

	for {
		_, err = os.Stat(dst)
		// 文件不在就break
		if err != nil {
			break
		}
		file, _ := fileHeader.Open()
		fileHash := md5.FileToMD5(file)
		oldFile, _ := os.Open(dst)
		oldfileHash := md5.FileToMD5(oldFile)
		if fileHash == oldfileHash {

			break
		}
		// 文件内容不一致，需要改名称
		//fmt.Println("文件内容不一致", oldfileHash, fileHash)

		// 文件名重复，但是内容不同，需要改名称
		name := strings.TrimRight(fileHeader.Filename, path.Ext(fileHeader.Filename))
		ReName := random.RandStr(3)
		ext := path.Ext(fileHeader.Filename)
		dst = path.Join("uploads", global.Config.Upload.Dir, auth.Username, fmt.Sprintf("%s_%s%s", name, ReName, ext))
	}

	// 确保目录存在
	if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
		resp.FailWithMsg(fmt.Sprintf("保存文件失败: %v", err), c)
		c.Abort()
		return
	}
	fmt.Printf("文件 %s 上传成功\n", dst)
	resp.Ok("图片上传成功", "/"+dst, c)
}

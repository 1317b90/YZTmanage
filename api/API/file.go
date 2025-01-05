package API

import (
	"YZT/Config"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// 上传文件 统一
func Up_file(c *gin.Context) {
	file, err := c.FormFile("file")
	// 文件名称
	fileName := c.DefaultQuery("fileName", file.Filename)

	if err != nil {
		c.JSON(500, gin.H{"message": "文件上传失败"})
		return
	}

	// 保存文件到 "file/" 路径下
	dst := filepath.Join("file", fileName)
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(500, gin.H{"message": "保存文件失败"})
		return
	}

	// 将Windows路径分隔符转换为URL路径分隔符
	urlPath := strings.ReplaceAll(dst, "\\", "/")
	c.JSON(200, gin.H{"message": "ok", "url": fmt.Sprintf("%s/%s", Config.Local_Addr, urlPath)})
}

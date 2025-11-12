package user

import (
	"compus-second-hand/global"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

//注册之后上传头像

func PutImage() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, ok := c.Get("id")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			return
		}
		img, err := c.FormFile("img")
		if err != nil {
			global.Logger.Error("api:上传头像时获取头像发生错误")
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    401,
				"message": "获取头像失败",
			})
			return
		}
		file, err := img.Open()
		if err != nil {
			global.Logger.Error("api:上传头像时打开头像发生错误")
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    402,
				"message": "服务器内部错误",
			})
			return
		}
		defer file.Close()
		imgData, err := io.ReadAll(file)
		if err != nil {
			global.Logger.Error("api:上传头像时读取头像发生错误")
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    402,
				"message": "服务器内部错误",
			})
			return
		}
	}

}

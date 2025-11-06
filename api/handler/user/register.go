package user

import (
	"compus-second-hand/api/utils"
	"compus-second-hand/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    401,
			"message": "上传头像失败",
		})
		return
	}
	name := c.PostForm("username")
	password := c.PostForm("password")
	md5Password, err := utils.Md5(password)
	if err != nil {
		global.Logger.Error("加密密码失败/err:" + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    402,
			"message": "服务器内部发生错误",
		})
		return
	}
	gender := c.PostForm("gender")
	email := c.PostForm("email")
	campusID := c.PostForm("campus")
	if file == nil || name == "" || password == "" || md5Password == "" || gender == "" || email == "" || campusID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "参数错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

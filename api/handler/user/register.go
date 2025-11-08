package user

import (
	"compus-second-hand/api/utils"
	"compus-second-hand/global"
	pb "compus-second-hand/rpc/user/pb"
	"context"
	"net/http"
	"strconv"

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
	//检测参数是否为空
	if file == nil || name == "" || password == "" || md5Password == "" || gender == "" || email == "" || campusID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    403,
			"message": "参数不全或格式错误",
		})
		return
	}
	//将参数转换为int64
	campusInt, _ := strconv.ParseInt(campusID, 10, 64)
	genderInt, _ := strconv.ParseInt(gender, 10, 64)
	//创建请求的结构体
	req := pb.RegisterRequest{
		Username: name,
		Password: md5Password,
		Email:    email,
		Campus:   campusInt,
		Gender:   genderInt,
	}
	//调用rpc服务
	resp, err := UserMicroClient.Register(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    402,
			"message": "服务器内部发生错误",
		})
		return
	}
	if resp.Code != 0 {
		switch resp.Code {
		case 1:
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    404,
				"message": "用户名或邮箱已存在",
			})
			return
		case 2:
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    405,
				"message": "所选校区不存在",
			})
			return
		case 3:
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    406,
				"message": "必填项不能为空",
			})
			return
		case 4:
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    402,
				"message": "服务器内部发生错误",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

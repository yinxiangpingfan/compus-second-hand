package user

import (
	"compus-second-hand/api/utils"
	"compus-second-hand/global"
	pb "compus-second-hand/rpc/user/pb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		email := c.PostForm("email")
		if username == "" && email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    411,
				"message": "用户名或邮箱必须填写一个",
				"id":      "",
				"token":   "",
			})
			return
		}
		if password == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    412,
				"message": "密码不能为空",
				"id":      "",
				"token":   "",
			})
			return
		}
		md5Password, err := utils.Md5(password)
		if err != nil {
			global.Logger.Error("登陆：加密密码失败/err:" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    413,
				"message": "服务器内部发生错误",
				"id":      "",
				"token":   "",
			})
			return
		}
		req := pb.LoginRequest{
			Username: username,
			Password: md5Password,
			Email:    email,
		}
		res, err := UserMicroClient.Login(context.Background(), &req)
		if err != nil {
			global.Logger.Error("登陆：rpc出现错误/err:" + err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    413,
				"message": "服务器内部发生错误",
				"id":      "",
				"token":   "",
			})
			return
		}
		if res.Code != 0 {
			switch res.Code {
			case 1:
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    414,
					"message": "密码错误",
					"id":      "",
					"token":   "",
				})
				return
			case 2:
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    415,
					"message": "用户不存在",
					"id":      "",
					"token":   "",
				})
				return
			case 0:
				var token string
				token, err := utils.GenerateToken(int(res.Id))
				if err != nil {
					global.Logger.Error("登陆：生成token失败/err:" + err.Error())
					c.JSON(http.StatusBadRequest, gin.H{
						"code":    413,
						"message": "服务器内部发生错误",
						"id":      "",
						"token":   "",
					})
					return
				}
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "登陆成功",
					"id":      res.Id,
					"token":   token,
				})
				return
			default:
				c.JSON(http.StatusBadRequest, gin.H{
					"code":    413,
					"message": "服务器内部发生错误",
					"id":      "",
					"token":   "",
				})
				return
			}
		}
	}
}

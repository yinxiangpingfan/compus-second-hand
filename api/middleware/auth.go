package middleware

import (
	"compus-second-hand/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//用于鉴权

func MidJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) < 7 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			c.Abort()
			return
		}
		tokenString := c.GetHeader("Authorization")
		//去除tokenString前面的Bearer
		tokenString = tokenString[7:]
		id, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权"})
			c.Abort()
			return
		}
		c.Set("id", id)
		c.Next()
	}
}

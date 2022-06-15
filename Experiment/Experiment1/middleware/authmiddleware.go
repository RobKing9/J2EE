package middleware

import (
	"Experiment1/common"
	"Experiment1/dao"
	"Experiment1/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		fmt.Print("请求token", tokenString)

		//validate token formate
		//if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
		//	c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足"})
		//	c.Abort()
		//	return
		//}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		//token通过验证, 获取claims中的UserID
		name := claims.Name
		DB := dao.GetDB()
		var u models.User
		DB.First(&u, name)

		// 验证用户是否存在
		if u.Name == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "权限不足"})
			c.Abort()
			return
		}

		//用户存在 将user信息写入上下文
		c.Set("user", u)

		c.Next()
	}
}

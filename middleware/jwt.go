package middleware

import (
	"github.com/catbugdemo/project_order/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	TOKEN_FIELD = "access-token"
)

//
func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(TOKEN_FIELD)

		// 解析 token
		if err := utils.ParseJWT(jwtSecret, token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error_no": 1, "error_msg": err.Error()})
		}

		// TODO 认证用户以及接下来的操作

		c.Next()
	}
}

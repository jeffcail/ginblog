package middleware

import (
	"gin-blog-api/handler"
	"gin-blog-api/types"
	"gin-blog-api/util"
	"github.com/gin-gonic/gin"
	"strings"
)

//AuthMiddleware
func AuthMiddleware() func(c *gin.Context)  {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			util.Error(c, int(types.ApiCode.NOAUTH), types.ApiCode.GetMessage(types.ApiCode.NOAUTH))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			util.Error(c, int(types.ApiCode.AUTHFORMATERROR), types.ApiCode.GetMessage(types.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}

		mc, err := handler.ParseToken(parts[1])
		if err != nil {
			util.Error(c, int(types.ApiCode.INVALIDTOKEN), types.ApiCode.GetMessage(types.ApiCode.INVALIDTOKEN))
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}

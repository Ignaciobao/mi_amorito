package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DeviceAuth 设备认证中间件
func DeviceAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceID := c.GetHeader("X-Device-ID")
		if deviceID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "missing device ID",
				"code":  "DEVICE_ID_REQUIRED",
			})
			c.Abort()
			return
		}

		// 将设备ID存储到上下文中
		c.Set("device_id", deviceID)
		c.Next()
	}
}
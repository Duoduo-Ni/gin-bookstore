package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			//c.Header("Access-Control-Allow-Origin", "http://192.168.36.61:8081") // 可将将 * 替换为指定的域名  //全开的跨域不能带cookie，所以这边之给一个ip开了跨域
			//c.Header("Access-Control-Allow-Origin", "http://192.168.31.162:8081")
			c.Header("Access-Control-Allow-Origin", "http://127.0.0.1:8081")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

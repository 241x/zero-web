package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/241x/zero-web/config"

	"github.com/gin-gonic/gin"
)

// CORS 跨域中间件。
func CORS(cfg config.CorsConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(cfg.AllowOrigins) > 0 {
			c.Header("Access-Control-Allow-Origin", strings.Join(cfg.AllowOrigins, ", "))
		}
		c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(cfg.AllowCredentials))
		c.Header("Access-Control-Allow-Headers", strings.Join(cfg.AllowHeaders, ", "))
		c.Header("Access-Control-Allow-Methods", strings.Join(cfg.AllowMethods, ", "))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

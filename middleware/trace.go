package middleware

import (
	"time"

	"github.com/241x/zero-web/ctxkeys"

	"github.com/241x/zero-kit/gormutil"
	"github.com/241x/zero-kit/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Trace 请求链路追踪中间件，注入 traceID、beginTime、logger 到上下文。
func Trace(log logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		ctx := c.Request.Context()
		ctx = ctxkeys.WithTraceID(ctx, traceID)
		ctx = gormutil.WithTraceID(ctx, traceID)
		ctx = ctxkeys.WithBeginTime(ctx, time.Now())

		l := log.With("traceId", traceID)
		ctx = l.WithContext(ctx)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

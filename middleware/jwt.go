// Package middleware 提供可复用的 gin 中间件。
package middleware

import (
	"strings"

	"github.com/241x/zero-kit/apperror"
	"github.com/241x/zero-web/errcode"
	"github.com/241x/zero-web/response"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const jwtClaimsKey = "jwt:claims"

// JWTGuard JWT 验证中间件，解析 Bearer Token 并将 claims 注入 gin.Context。
// 验证成功后可通过 GetJWTClaims 获取完整 claims，或通过 GetJWTUserID 直接获取用户 ID。
func JWTGuard(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")

		if !strings.HasPrefix(tokenString, "Bearer ") {
			unauthorized(c)
			return
		}

		token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (any, error) {
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

		if err != nil {
			unauthorized(c)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			unauthorized(c)
			return
		}

		c.Set(jwtClaimsKey, claims)
		c.Next()
	}
}

// GetJWTClaims 从 gin.Context 中读取 JWTGuard 解析出的完整 claims。
func GetJWTClaims(c *gin.Context) jwt.MapClaims {
	v, _ := c.Get(jwtClaimsKey)
	claims, _ := v.(jwt.MapClaims)
	return claims
}

// GetJWTUserID 从 claims 中便捷提取 user_id。
func GetJWTUserID(c *gin.Context) uint32 {
	claims := GetJWTClaims(c)
	if claims == nil {
		return 0
	}
	userId, _ := claims["user_id"].(float64)
	return uint32(userId)
}

func unauthorized(c *gin.Context) {
	response.Error(c, apperror.New(errcode.Unauthorized))
	c.Abort()
}

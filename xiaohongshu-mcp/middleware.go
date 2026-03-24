package main

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// corsMiddleware CORS 中间件，仅允许 localhost origin
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		if isLocalhostOrigin(origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// isLocalhostOrigin 检查 origin 是否为 localhost
func isLocalhostOrigin(origin string) bool {
	if origin == "" {
		return false
	}
	lower := strings.ToLower(origin)
	// 匹配 http(s)://localhost[:port]、http(s)://127.0.0.1[:port]、http(s)://[::1][:port]
	for _, prefix := range []string{
		"http://localhost", "https://localhost",
		"http://127.0.0.1", "https://127.0.0.1",
		"http://[::1]", "https://[::1]",
	} {
		if lower == prefix || strings.HasPrefix(lower, prefix+":") {
			return true
		}
	}
	return false
}

// authMiddleware Bearer Token 认证中间件，使用常量时间比较防止 timing attack
func authMiddleware(token string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token == "" {
			c.Next()
			return
		}

		auth := c.GetHeader("Authorization")
		expected := "Bearer " + token
		if subtle.ConstantTimeCompare([]byte(auth), []byte(expected)) != 1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			return
		}

		c.Next()
	}
}

// errorHandlingMiddleware 错误处理中间件
func errorHandlingMiddleware() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		logrus.Errorf("服务器内部错误: %v, path: %s", recovered, c.Request.URL.Path)

		respondError(c, http.StatusInternalServerError, "INTERNAL_ERROR",
			"服务器内部错误", recovered)
	})
}

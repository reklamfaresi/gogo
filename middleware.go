package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization başlığını al
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Yetkisiz erişim, token eksik"})
			c.Abort()
			return
		}

		// Bearer token kısmını ayır
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Yetkisiz erişim, geçersiz token formatı"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Token'ı doğrula
		claims := &jwt.RegisteredClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Yetkisiz erişim, geçersiz veya süresi dolmuş token"})
			c.Abort()
			return
		}

		// Kullanıcıyı isteğe ekle
		c.Set("username", claims.Subject)

		// Devam et
		c.Next()
	}
}

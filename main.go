package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Veritabanı bağlantısını başlat
	ConnectDatabase()

	// Gin router oluştur
	r := gin.Default()

	// Login endpoint
	r.POST("/login", LoginHandler)

	// Authenticated group
	auth := r.Group("/")
	auth.Use(AuthMiddleware())
	{
		// Ayarlar endpoint'i (örnek endpoint)
		auth.GET("/settings", func(c *gin.Context) {
			// Veritabanından ayarları getireceğiz (örnek veri)
			c.JSON(http.StatusOK, gin.H{
				"site_title":  "GoGo Admin Paneli",
				"description": "Bu, admin panelinin genel ayarlarıdır.",
			})
		})
	}

	// Sunucuyu başlat
	r.Run(":8080")
}

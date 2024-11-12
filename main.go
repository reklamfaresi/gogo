package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	ConnectDatabase()
	r := gin.Default()

	// Login endpoint
	r.POST("/login", func(c *gin.Context) {
		// Burada login işlemleri yapılacak
		c.JSON(http.StatusOK, gin.H{
			"message": "Login endpoint'i burada!",
		})
	})

	// Sunucuyu başlat
	r.Run(":8080") // 8080 portunda çalıştır
}

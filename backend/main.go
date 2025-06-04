package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"derma-detection/api"
	"derma-detection/services"
)

func main() {
	// Python servisini başlat
	pythonService := services.NewPythonService("http://localhost:5001")

	// API işleyicisini oluştur
	handler := api.NewHandler(pythonService, "./uploads")

	// Gin router'ı oluştur
	router := gin.Default()

	// CORS ayarlarını yapılandır
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	router.Use(cors.New(config))

	// Rotaları tanımla
	router.POST("/api/upload/image", handler.UploadImage)
	router.POST("/api/upload/pdf", handler.UploadPDF)

	// Sunucuyu başlat
	log.Println("Sunucu başlatılıyor...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Sunucu başlatılamadı: %v", err)
	}
}

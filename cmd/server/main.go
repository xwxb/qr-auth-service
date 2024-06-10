package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xwxb/qr-auth-service/internal/config"
	"github.com/xwxb/qr-auth-service/internal/handler"
	"github.com/xwxb/qr-auth-service/internal/service"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	service.Init(cfg)

	r := gin.Default()
	initRouter(r)
	portStr := fmt.Sprintf(":%d", cfg.Server.Port)
	r.Run(portStr)
}

func initRouter(r *gin.Engine) {
	authRoute := r.Group("/auth")
	{
		qrcodeRoute := authRoute.Group("/qrcode")
		qrcodeRoute.POST("/verify", handler.VerifyUsername)
		qrcodeRoute.POST("/uuid", handler.GenUUID)
		qrcodeRoute.GET("/status", handler.VerifyQRCode)
	}
}

package service

import (
	"github.com/xwxb/qr-auth-service/internal/config"
	"log"
)

func Init(cfg *config.Config) {
	var err error
	redisServiceIns, err = NewRedisService(cfg.Redis.Address, cfg.Redis.Password, cfg.Redis.Tls, cfg.Redis.DB)
	if err != nil {
		log.Fatalf("failed to initialize Redis service: %v", err)
	}
}

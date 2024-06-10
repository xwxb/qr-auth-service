package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	client *redis.Client
}

var redisServiceIns *RedisService

func NewRedisService(addr, password string, tls bool, db int) (*RedisService, error) {
	var uri string
	if tls {
		uri = fmt.Sprintf("rediss://default:%s@%s", password, addr)
	} else {
		uri = fmt.Sprintf("redis://%s:%s", addr, password)
	}

	opt, err := redis.ParseURL(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Redis URI: %v", err)
	}
	opt.DB = db // 设置数据库索引

	client := redis.NewClient(opt)

	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	return &RedisService{
		client: client,
	}, nil
}

func (s *RedisService) SetVerification(uuid string, expiration time.Duration) error {
	err := s.client.Set(context.Background(), uuid, "verified", expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set verification in Redis: %v", err)
	}
	return nil
}

func (s *RedisService) IsVerificationSuccessful(uuid string) (bool, error) {
	val, err := s.client.Get(context.Background(), uuid).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, fmt.Errorf("failed to get verification from Redis: %v", err)
	}
	return val == "verified", nil
}

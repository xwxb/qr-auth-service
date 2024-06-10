package service

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const (
	redisKeyPrefix = "auth:"
)

func SaveAuthSessionToRedis(uuid string, expiration time.Duration) error {
	key := fmt.Sprintf("%s%s", redisKeyPrefix, uuid)
	err := redisServiceIns.client.Set(context.Background(), key, "1", expiration).Err()
	if err != nil {
		return fmt.Errorf("failed to set verification in Redis: %v", err)
	}
	return nil
}

func CheckAuthSessionFromRedis(uuid string) (bool, error) {
	key := fmt.Sprintf("%s%s", redisKeyPrefix, uuid)
	val, err := redisServiceIns.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, fmt.Errorf("failed to get verification from Redis: %v", err)
	}
	return val == "verified", nil
}

func CheckUserNameFormRedis(userName string) (bool, error) {
	key := fmt.Sprintf("%s%s", redisKeyPrefix, userName)
	val, err := redisServiceIns.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, fmt.Errorf("failed to get verification from Redis: %v", err)
	}
	return val == "verified", nil
}

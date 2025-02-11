package token

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func SaveTokenToRedis(userID, token string, exp time.Duration, redisClient *redis.Client) error {
	key := "session:" + userID
	return redisClient.Set(ctx, key, token, exp).Err()
}

func InvalidateSession(userID string, redisClient *redis.Client) error {
	key := "session" + userID
	return redisClient.Del(ctx, key).Err()
}

func IsTokenValid(userID, token string, redisClient *redis.Client) bool {
	storedToken, err := redisClient.Get(ctx, "session:"+userID).Result()
	if err != nil {
		return false
	}
	return storedToken == token
}

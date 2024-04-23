package ctl

import (
	"context"
	"encoding/json"
	"time"
	"todo_list/app/user/repository/cache"

	"github.com/redis/go-redis/v9"
)

var userKey string = "user_key"

type UserInfo struct {
	Id uint `json:"id"`
}

func NewContext(ctx context.Context, u *UserInfo) context.Context {
	return context.WithValue(ctx, userKey, u)
}

func FromContext(ctx context.Context) (*UserInfo, bool) {
	u, ok := ctx.Value(userKey).(*UserInfo)
	return u, ok
}

func InitUserinfo(ctx context.Context, user *UserInfo) error {
	cachedUser, err := GetUserInfo(ctx)
	if err != nil {
		return err
	}

	if cachedUser != nil {
		return nil
	}

	if err := SetUserInCache(ctx, user, cache.RedisClient); err != nil {
		return err
	}

	return nil
}

func GetUserInfo(ctx context.Context) (*UserInfo, error) {
	data, err := cache.RedisClient.Get(ctx, userKey).Bytes()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	var user UserInfo
	if err = json.Unmarshal(data, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func SetUserInCache(ctx context.Context, user *UserInfo, redisClient *redis.Client) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := redisClient.Set(ctx, userKey, data, time.Second*30).Err(); err != nil {
		return err
	}
	return nil
}

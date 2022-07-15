package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
)

type (
	WrapClient struct {
		*redis.Client
		ctx context.Context
	}

	Config struct {
		Addr     string
		Password string
		DB       int
	}
)

func NewRedis(conf *Config) (c *WrapClient) {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           conf.DB,
		MaxConnAge:   time.Minute * 30, // 连接池连接有效时间
		MinIdleConns: 4,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	})

	ctx := context.Background()
	c = &WrapClient{}
	c.ctx = ctx
	c.Client = client
	ping := c.Ping(ctx)
	if ping.Err() != nil {
		panic(errors.Errorf("连接redis失败:%s", ping.Err()))
	}
	return
}

// GetString 字符串
func (r *WrapClient) GetString(key string) (string, error) {
	return r.Client.Get(r.ctx, key).Result()
}

// SetSimple 通用set
func (r *WrapClient) SetSimple(key string, value interface{}, t ...time.Duration) (string, error) {
	var t2 time.Duration
	if len(t) > 0 {
		t2 = t[0]
	}
	return r.Client.Set(r.ctx, key, value, t2).Result()
}

//GetJson json序列化
func (r *WrapClient) GetJson(key string) (interface{}, error) {
	res := r.Client.Get(r.ctx, key)
	if res.Err() != nil {
		return nil, res.Err()
	}

	b, err := res.Bytes()
	if err != nil {
		return nil, errors.Errorf("get key:%s 反序列化json失败(-1)", key)
	}
	var result interface{}
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, errors.Errorf("get key:%s 反序列化json失败(-2)", key)
	}
	return result, nil
}

//SetJson json序列化set
func (r *WrapClient) SetJson(key string, value interface{}, t ...time.Duration) (string, error) {
	var t2 time.Duration
	if len(t) > 0 {
		t2 = t[0]
	}
	v, err := json.Marshal(value)
	if err != nil {
		return "", fmt.Errorf("set key:%s 序列化json失败", key)
	}
	return r.Client.Set(r.ctx, key, v, t2).Result()
}

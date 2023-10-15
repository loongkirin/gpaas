package util

import (
	"context"
	"fmt"
	"time"

	app "github.com/loongkirin/gpaas/app"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
)

func NewDefaultCaptchaRedisStore() *CaptchaRedisStore {
	return &CaptchaRedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA_",
	}
}

type CaptchaRedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *CaptchaRedisStore) UseWithContext(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *CaptchaRedisStore) Set(id string, value string) error {
	err := app.AppContext.APP_REDIS.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (rs *CaptchaRedisStore) Get(key string, clear bool) string {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// exists, err := app.AppContext.APP_REDIS.Exists(rs.Context, key).Result()

	// if err != nil {
	// 	fmt.Println(err)
	// 	return ""
	// }
	// if exists <= 0 {
	// 	fmt.Println("key not exists")
	// 	return ""
	// }

	val, err := app.AppContext.APP_REDIS.Get(rs.Context, key).Result()
	if err == redis.Nil {
		fmt.Println(key, " does not exist")
	} else if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := app.AppContext.APP_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

func (rs *CaptchaRedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}

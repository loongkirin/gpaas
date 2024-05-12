package app

import (
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/loongkirin/gdk/database/gorm"
	cfg "github.com/loongkirin/gpaas/pkg/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
)

type appContext struct {
	APP_CONFIG                 cfg.AppConfig
	APP_VP                     *viper.Viper
	APP_REDIS                  *redis.Client
	APP_DbContext              gorm.DbContext
	APP_Concurrency_Controller *singleflight.Group
}

var AppContext appContext

func InitAppContext() {
	AppContext = appContext{
		APP_Concurrency_Controller: &singleflight.Group{},
	}
	AppContext.initViper()
	AppContext.initRedis()
	AppContext.initDbContext()
}

func (ctx *appContext) initViper() {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	vp.AddConfigPath("./")
	err := vp.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error when read config file: %s", err))
	}

	vp.WatchConfig()

	vp.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := vp.Unmarshal(&ctx.APP_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err := vp.Unmarshal(&ctx.APP_CONFIG); err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("fatal error when unmarshal config file: %s", err))
	}

	ctx.APP_VP = vp
}

func (ctx *appContext) initRedis() {
	redisCfg := ctx.APP_CONFIG.RedisConfig
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	} else {
		ctx.APP_REDIS = client
	}
}

func (ctx *appContext) initDbContext() {
	ctx.APP_DbContext = gorm.CreateDbContext(ctx.APP_CONFIG.DbConfig)
}

package app

import (
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
	core "github.com/loongkirin/gpaas/core"
	postgres "github.com/loongkirin/gpaas/domain/infrastructure/postgres"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type appContext struct {
	APP_CONFIG    core.AppConfig
	APP_VP        *viper.Viper
	APP_REDIS     *redis.Client
	APP_DbContext core.DbContext
}

var AppContext appContext

func InitAppContext() {
	AppContext = appContext{}
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
		panic(fmt.Errorf("Fatal error when read config file: %s", err))
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
		panic(fmt.Errorf("Fatal error when unmarshal config file: %s", err))
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
	ctx.APP_DbContext = createDbContext(ctx.APP_CONFIG.DbConfig)
}

func createDbContext(cfg core.DbConfig) core.DbContext {
	var dbcontext core.DbContext
	if cfg.DbType == "postgres" {
		postgresDbCtx := postgres.NewDbContext(&cfg)
		dbcontext = &postgresDbCtx
	}

	return dbcontext
}

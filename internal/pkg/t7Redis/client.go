package t7Redis

import (
	"cli-tool/internal/pkg/config"
	"github.com/Template7/common/logger"
	"github.com/go-redis/redis"
	"sync"
)

const (
	VerifyCodePrefix = "verifyCode"
)

var (
	once     sync.Once
	instance *redis.Client
	log = logger.GetLogger()
)

func New() *redis.Client {
	once.Do(func() {
		instance = redis.NewClient(&redis.Options{
			Addr:     config.New().Redis.Host,
			Password: config.New().Redis.Password,
			//PoolSize: config.New().Redis.PollSize,
			//ReadTimeout: time.Duration(config.Redis.ReadTimeout >> 9), // nano second
		})
		if err := instance.Ping().Err(); err != nil {
			log.Fatal(err)
		}
		log.Debug("redis client initialized")
	})
	return instance
}

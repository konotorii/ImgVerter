package connectors

import (
	"context"
	"imgverter/util"

	"github.com/redis/go-redis/v9"
)

var (
	RedixCTX = context.TODO()
	RedisDB  *redis.Client
)

func RedisDatabaseInit() (*redis.Client, error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     util.Config.RedisHost,
		Password: util.Config.RedisPass,
		DB:       util.Config.RedisDb,
	})
	if err := RedisDB.Ping(RedixCTX).Err(); err != nil {
		return nil, err
	}
	return RedisDB, nil
}

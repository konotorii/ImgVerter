package connectors

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"imgverter/util"
)

var (
	RsErrNil       = errors.New("no matching record found in redis database")
	Ctx            = context.TODO()
	RedisDB *redis.Client
)

func RedisDatabaseInit() (*redis.Client, error) {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     util.Config.RedisHost,
		Password: util.Config.RedisPass,
		DB:       util.Config.RedisDb,
	})
	if err := RedisDB.Ping(Ctx).Err(); err != nil {
		return nil, err
	}
	return RedisDB, nil
}
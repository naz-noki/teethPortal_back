package cacheDB

import (
	"fmt"

	"github.com/go-redis/redis"
)

type cacheDB struct {
	client *redis.Client
}

var (
	DB *cacheDB
)

func (t *cacheDB) Close() error {
	return t.client.Close()
}

func New(
	host string, port int,
) error {
	DB = new(cacheDB)

	DB.client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
	})

	_, errPing := DB.client.Ping().Result()

	if errPing != nil {
		return errPing
	}
	return nil
}

package redisclient

import (
	"sync"
)

type RedisClientManager struct {
	redisClientMap *sync.Map
}

var redisClientManager *RedisClientManager = nil

func RedisClientManagerInit(configs []RedisClientConfig) error {
	redisClientManager = &RedisClientManager{
		redisClientMap: &sync.Map{},
	}

	for _, config := range configs {
		client, err := NewRedisClient(config)
		if err != nil {
			Log.Error(err)
			return err
		}

		redisClientManager.redisClientMap.Store(config.Name, client)
	}
	return nil
}

func GetRedisClientManager() *RedisClientManager {
	return redisClientManager
}

func (r *RedisClientManager) GetRedisClient(name string) *RedisClient {
	clientobj, ok := r.redisClientMap.Load(name)
	if !ok {
		Log.Warnf("can't find client in map, name=%v", name)
		return nil
	}

	client := clientobj.(*RedisClient)

	return client
}

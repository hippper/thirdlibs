package redisclient

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	. "github.com/luckyweiwei/base/logger"
)

// Redis client config
type RedisClientConfig struct {
	Name         string
	Addr         string
	Active       int // pool
	Idle         int // pool
	DialTimeout  int
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	DBNum        string // db数据库序号
	Password     string
}

type RedisClient struct {
	env  string
	pool *redis.Pool
}

func NewRedisClient(c RedisClientConfig) (*RedisClient, error) {

	env := fmt.Sprintf("[%s]tcp@%s", c.Name, c.Addr)
	cnop := redis.DialConnectTimeout(time.Duration(c.DialTimeout) * time.Second)
	rdop := redis.DialReadTimeout(time.Duration(c.ReadTimeout) * time.Second)
	wrop := redis.DialWriteTimeout(time.Duration(c.WriteTimeout) * time.Second)

	dialFunc := func() (redis.Conn, error) {
		conn, err := redis.Dial("tcp", c.Addr, cnop, rdop, wrop)
		if err != nil {
			Log.Errorf("Redis connect %s error: %s", env, err)
			return nil, err
		}

		if c.Password != "" {
			_, err = conn.Do("AUTH", c.Password)
			if err != nil {
				Log.Errorf("Redis %s AUTH(password: %s) error: %s", env, c.Password, err)
				conn.Close()
				return nil, err
			}

		}

		_, err = conn.Do("SELECT", c.DBNum)
		if err != nil {
			Log.Errorf("Redis %s SELECT %s error: %s", env, c.DBNum, err)
			conn.Close()
			return nil, err
		}
		return conn, nil
	}

	redisClient := &RedisClient{
		env: env,
		pool: &redis.Pool{
			MaxActive:   c.Active,
			MaxIdle:     c.Idle,
			IdleTimeout: time.Duration(c.IdleTimeout),
			Wait:        true,
			Dial:        dialFunc,
		},
	}

	return redisClient, nil
}

func (r *RedisClient) Get() redis.Conn {
	return r.pool.Get()
}

func (r *RedisClient) Close() error {
	return r.pool.Close()
}

func (r *RedisClient) Err() error {
	conn := r.Get()
	return conn.Err()
}

func (r *RedisClient) CloseConn() error {
	conn := r.Get()
	return conn.Close()
}

func (r *RedisClient) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	conn := r.Get()
	defer conn.Close()

	return conn.Do(commandName, args...)
}

// Pipelining
// Send
// Flush
// Receive

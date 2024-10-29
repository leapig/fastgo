package helper

import (
	"github.com/dubbogo/gost/log/logger"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/redigo"
	"github.com/gomodule/redigo/redis"
	"os"
	"time"
)

var luaScriptBatchDel = `
		local i=0
		local count=0
		local successful=0
		while (i<#ARGV)
			do
				i=i+1
				redis.call('DEL',ARGV[i])
				successful=successful+1
			end
		return successful
	`
var RS *Redis

type Redis struct {
	Url         string
	MaxIdle     int
	IdleTimeout time.Duration
	MaxActive   int
	Pool        *redis.Pool
	Wait        bool
	Password    string
}

func RegisterRedis() *Redis {
	r := Redis{
		Url:         os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		MaxActive:   10000,
		Wait:        false,
		Password:    os.Getenv("REDIS_PASSWORD"),
	}
	r.Pool = &redis.Pool{
		MaxIdle:     r.MaxIdle,
		IdleTimeout: r.IdleTimeout,
		MaxActive:   r.MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", r.Url)
			if r.Password != "" {
				if _, err := c.Do("AUTH", r.Password); err != nil {
					_ = c.Close()
					logger.Errorf("%v", err)
					return nil, err
				}
			}
			return c, err
		},
		Wait: r.Wait,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	RS = &r
	return &r
}

func (r *Redis) CloseRedis() {
	err := r.Pool.Close()
	if err != nil {
		panic(err)
	}
}

func (r *Redis) Set(key, value string) error {
	conn := r.Pool.Get()
	_, err := conn.Do("Set", key, value)
	defer r.CloseRedisConn(conn)
	return err
}

func (r *Redis) SetEx(key string, seconds int, value string) error {
	conn := r.Pool.Get()
	_, err := conn.Do("SetEX", key, seconds, value)
	defer r.CloseRedisConn(conn)
	return err
}

func (r *Redis) Get(key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("Get", key)
}

func (r *Redis) GetString(key string) (string, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return redis.String(conn.Do("Get", key))
}

func (r *Redis) Del(key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("Del", key)
}

func (r *Redis) HSet(name, key string, value interface{}) error {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	_, err := conn.Do("HSet", name, key, value)
	return err
}

func (r *Redis) HMSet(key string, value interface{}) error {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	_, err := conn.Do("HMSet", key, value)
	return err
}

func (r *Redis) HGetAll(name string) (map[string]string, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return redis.StringMap(conn.Do("HGetAll", name))
}

func (r *Redis) HGet(name, key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("HGet", name, key)
}

func (r *Redis) HIncrby(name, key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("HINCRBY", name, key)
}

func (r *Redis) HIncrbyIncrement(name, key string, increment int) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("HINCRBY", name, key, increment)
}
func (r *Redis) HDel(name, key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("HDel", name, key)
}

func (r *Redis) LPush(name, key string) error {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	_, err := conn.Do("LPush", name, key)
	return err
}

func (r *Redis) RPop(key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("RPop", key)
}

func (r *Redis) GetListLen(key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return conn.Do("LLen", key)
}

func (r *Redis) LRange(key string, start, end int) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return redis.Values(conn.Do("LRange", key, start, end))
}

func (r *Redis) LREM(key string, value string, count int) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	return redis.Values(conn.Do("LREM", key, value, count))
}

func (r *Redis) Expire(key string, seconds int) error {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	_, err := conn.Do("EXPIRE", key, seconds)
	return err
}

func (r *Redis) MGet(key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("MGET", key)
	return res, err
}

func (r *Redis) ZAdd(key string, score int64, value string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("ZADD", key, score, value)
	return res, err
}
func (r *Redis) SAdd(key string, value string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("SADD", key, value)
	return res, err
}
func (r *Redis) SIsmembers(key string, value string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("SISMEMBER", key, value)
	return res, err
}
func (r *Redis) SMembers(key string) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("SMEMBERS", key)
	return res, err
}
func (r *Redis) ZRang(key string, min int64, max int64) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("ZRange", key, min, max)
	return res, err
}

func (r *Redis) ZRangWithScore(key string, min int64, max int64) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("ZRANGEBYSCORE", key, min, max)
	return res, err
}

func (r *Redis) ZMinWithScore(key string, min int64) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("ZRANGEBYSCORE", key, "-inf", min)
	return res, err
}
func (r *Redis) ZMaxWithScore(key string, max int64) (interface{}, error) {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)
	res, err := conn.Do("ZRANGEBYSCORE", key, max, "+inf")
	return res, err
}

// DeleteBatchByKey 批量删除key
func (r *Redis) DeleteBatchByKey(delKeys []string) error {
	conn := r.Pool.Get()
	defer r.CloseRedisConn(conn)

	lua := redis.NewScript(0, luaScriptBatchDel)
	_, err := redis.Int(lua.Do(conn, delKeys))
	return err
}

func (r *Redis) CloseRedisConn(c redis.Conn) {
	if err := c.Close(); err != nil {
		logger.Error("close redis error:", err)
	}
}

// GetRedSyncMutex  配置超时参数"redsync.WithTries(1000), redsync.WithExpiry(time.Second*20)"
func GetRedSyncMutex(key string, pool *redis.Pool, options ...redsync.Option) (*redsync.Mutex, error) {
	rs := redsync.New(redigo.NewPool(pool))
	mutex := rs.NewMutex(key, options...)
	if err := mutex.Lock(); err != nil {
		logger.Error("获取redis分布式锁异常", err)
		return nil, err
	}
	return mutex, nil
}

func UnlockRedSyncMutex(in *redsync.Mutex) {
	if ok, err := in.Unlock(); !ok || err != nil {
		logger.Error("释放redis分布式锁异常", err)
	}
}

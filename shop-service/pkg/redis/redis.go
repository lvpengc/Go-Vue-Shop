// Package redis 建立redis连接池
package redis

import (
	"log"
	"time"

	etcdStruct "github.com/hellozouyou/easy-mall/pkg/structure/etcd"

	"github.com/gomodule/redigo/redis"
)

// sprint from https://github.com/hopehook/golang-db

// ConnPool is RDS struct
type ConnPool struct {
	Pool redis.Pool
}

// Client redis 的本地连接
var Client *ConnPool

// Init init redis
func Init(config *etcdStruct.Redis) {
	Client = InitRedisPool(config.Host, config.Password, config.Database, config.MaxOpenConns, config.MaxIdleConns)
}

// InitRedisPool func init RDS fd
func InitRedisPool(host, password string, database, maxOpenConns, maxIdleConns int) *ConnPool {
	r := &ConnPool{}
	r.Pool = redis.Pool{
		MaxActive:   maxOpenConns, // max number of connections
		MaxIdle:     maxIdleConns,
		IdleTimeout: 120 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("select", database); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	if _, err := r.Do("PING"); err != nil {
		log.Panicln("Init redis pool failed.", err.Error())
	}
	return r
}

// GetConn pool
func (p *ConnPool) GetConn() redis.Conn {
	conn := p.Pool.Get()
	if conn.Err() != nil {
		log.Printf("GetConn redis fail, %v", conn.Err())
		time.Sleep(1)
		for {
			// 直到获取成功
			conn = p.Pool.Get()
			if conn.Err() == nil {
				break
			}
		}
	}

	return conn
}

// Close pool
func (p *ConnPool) Close() error {
	err := p.Pool.Close()
	return err
}

// Do commands
func (p *ConnPool) Do(command string, args ...interface{}) (interface{}, error) {
	conn := p.GetConn()
	defer conn.Close()
	return conn.Do(command, args...)
}

// Set for value
func (p *ConnPool) Set(key string, value interface{}) (interface{}, error) {
	conn := p.GetConn()
	defer conn.Close()
	return conn.Do("SET", key, value)
}

// SetEx 设置带有效期的value
func (p *ConnPool) SetEx(key string, seconds int64, value interface{}) (interface{}, error) {
	conn := p.GetConn()
	defer conn.Close()
	return conn.Do("SETEX", key, seconds, value)
}

// GetString for string
func (p *ConnPool) GetString(key string) (string, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.String(conn.Do("GET", key))
}

// GetBytes for bytes
func (p *ConnPool) GetBytes(key string) ([]byte, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

// GetInt for int
func (p *ConnPool) GetInt(key string) (int, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("GET", key))
}

// GetInt64 for int64
func (p *ConnPool) GetInt64(key string) (int64, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.Int64(conn.Do("GET", key))
}

// GetUint64 for uint64
func (p *ConnPool) GetUint64(key string) (uint64, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.Uint64(conn.Do("GET", key))
}

// Del for key
func (p *ConnPool) Del(key string) (interface{}, error) {
	conn := p.GetConn()
	defer conn.Close()
	return conn.Do("DEL", key)
}

// Expire for key
func (p *ConnPool) Expire(key string, seconds int64) (interface{}, error) {
	conn := p.GetConn()
	defer conn.Close()
	return conn.Do("EXPIRE", key, seconds)
}

// Keys for key
func (p *ConnPool) Keys(pattern string) ([]string, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.Strings(conn.Do("KEYS", pattern))
}

// KeysByteSlices for key
func (p *ConnPool) KeysByteSlices(pattern string) ([][]byte, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.ByteSlices(conn.Do("KEYS", pattern))
}

// SetHashMap for hash map
func (p *ConnPool) SetHashMap(key string, fieldValue map[string]interface{}) (interface{}, error) {
	conn := p.GetConn()
	defer conn.Close()
	return conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
}

// GetHashMapString for hash map
func (p *ConnPool) GetHashMapString(key string) (map[string]string, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.StringMap(conn.Do("HGETALL", key))
}

// GetHashMapInt for hash map
func (p *ConnPool) GetHashMapInt(key string) (map[string]int, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.IntMap(conn.Do("HGETALL", key))
}

// GetHashMapInt64 for hash map
func (p *ConnPool) GetHashMapInt64(key string) (map[string]int64, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.Int64Map(conn.Do("HGETALL", key))
}

// INCR 将 key 中储存的数字值增一，如果 key 不存在，那么 key 的值会先被初始化为 0 ，然后再执行 INCR 操作。
func (p *ConnPool) INCR(key string) (int, error) {
	conn := p.GetConn()
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

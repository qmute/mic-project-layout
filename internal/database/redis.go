package database

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/qmute/mic-project-layout/pkg/ut"
	"github.com/quexer/utee"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

// RedisOption Option redis连接选项
type RedisOption struct {
	Address  string `validate:"required"`
	Username string `validate:"-"`
	Password string `validate:"-"`
	Db       int    `validate:"gte=0"` // db number
	Pool     int    `validate:"gte=0"`
	Name     string `validate:"required"`
}

func NewRdb(opt RedisOption) (*Rdb, error) {
	utee.Chk(ut.ValidStruct(opt))

	poolSize := 20
	if opt.Pool > 0 {
		poolSize = opt.Pool
	}
	options := &redis.Options{
		Addr:     opt.Address,
		Password: opt.Password,
		DB:       opt.Db,
		PoolSize: poolSize,
	}

	if opt.Username != "" {
		options.Username = opt.Username
	}

	clt := redis.NewClient(options)

	if err := clt.Ping(context.Background()).Err(); err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info("redis connect success")

	return &Rdb{
		Clt: clt,
	}, nil
}

type Rdb struct {
	Clt redis.UniversalClient
}

// SaveInt 保存数值 expire(秒)
func (p *Rdb) SaveInt(ctx context.Context, key string, val int, expire int) error {
	err := p.Clt.Set(ctx, key, val, time.Duration(expire)*time.Second).Err()
	return errors.WithStack(err)
}

// GetInt 按键获取数值
func (p *Rdb) GetInt(ctx context.Context, key string) (int, error) {
	i, err := p.Clt.Get(ctx, key).Int()
	if err != nil && err == redis.Nil {
		return 0, nil
	}

	return i, errors.WithStack(err)
}

// MustGetInt 按键一定会获取数值，否则返回错误
func (p *Rdb) MustGetInt(ctx context.Context, key string) (int, error) {
	i, err := p.Clt.Get(ctx, key).Int()
	return i, errors.WithStack(err)
}

// SaveString 保存字符串
func (p *Rdb) SaveString(ctx context.Context, key string, val string, expire int) error {
	err := p.Clt.Set(ctx, key, val, time.Duration(expire)*time.Second).Err()
	return errors.WithStack(err)
}

// GetString 按键获取到数据
func (p *Rdb) GetString(ctx context.Context, key string) (string, error) {
	s, err := p.Clt.Get(ctx, key).Result()
	if err != nil && err == redis.Nil {
		return "", nil
	}
	return s, errors.WithStack(err)
}

// MustGetString 按键一定会获取到字符串数据，否则返回错误
func (p *Rdb) MustGetString(ctx context.Context, key string) (string, error) {
	s, err := p.Clt.Get(ctx, key).Result()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return s, nil
}

// SaveByte 保存字节数组
func (p *Rdb) SaveByte(ctx context.Context, key string, val []byte, expire int) error {
	err := p.Clt.Set(ctx, key, val, time.Duration(expire)*time.Second).Err()
	return errors.WithStack(err)
}

// GetByte 按键获取到数据
func (p *Rdb) GetByte(ctx context.Context, key string) ([]byte, error) {
	b, err := p.Clt.Get(ctx, key).Bytes()
	if err != nil && err == redis.Nil {
		return nil, nil
	}
	return b, errors.WithStack(err)
}

// MustGetByte 按键一定会获取到字节数组，否则返回错误
func (p *Rdb) MustGetByte(ctx context.Context, key string) ([]byte, error) {
	b, err := p.Clt.Get(ctx, key).Bytes()
	if err != nil {
		return b, errors.WithStack(err)
	}
	return b, nil
}

// TTL 检查过期时间， 返回秒
func (p *Rdb) TTL(ctx context.Context, key string) (int, error) {
	d, err := p.Clt.TTL(ctx, key).Result()
	if err == redis.Nil {
		// expire
		return 0, nil
	}
	return int(d / time.Second), errors.WithStack(err)
}

// DelKey 按键删除
func (p *Rdb) DelKey(ctx context.Context, key ...string) error {
	err := p.Clt.Del(ctx, key...).Err()
	return errors.WithStack(err)
}

func (p *Rdb) AutoCacheString(ctx context.Context, key string, useCache bool, f func() (string, int, error)) (string, error) {
	var result string

	if useCache {
		s, err := p.GetString(ctx, key)
		if err != nil {
			return "", err
		}
		result = s
	}

	if result == "" {
		s, expire, err := f()
		if err != nil {
			return "", err
		}
		err = p.SaveString(ctx, key, s, expire)
		if err != nil {
			return "", err
		}
		result = s
	}
	return result, nil
}

// IncrBy 为键 key 储存的数字值加上增量 increment
func (p *Rdb) IncrBy(ctx context.Context, key string, val int) error {
	err := p.Clt.IncrBy(ctx, key, int64(val)).Err()
	return errors.WithStack(err)
}

// Incr 为键 key 储存的数字值加1
func (p *Rdb) Incr(ctx context.Context, key string) (int, error) {
	result, err := p.Clt.Incr(ctx, key).Result()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(result), nil
}

// Expire 为给定 key 设置生存时间
func (p *Rdb) Expire(ctx context.Context, key string, second int) error {
	err := p.Clt.Expire(ctx, key, time.Duration(second)*time.Second).Err()
	return errors.WithStack(err)
}

func (p *Rdb) Keys(ctx context.Context, key string) ([]string, error) {
	result, err := p.Clt.Keys(ctx, key).Result()
	if err != nil && err == redis.Nil {
		return nil, nil
	}
	return result, errors.WithStack(err)
}

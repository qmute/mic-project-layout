package database

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/quexer/utee"
)

type MockRedisOpt struct {
	Size int // pool size 默认10
	Db   int // 默认0
}

func WithDockerRedisPool(size int) func(opt *MockRedisOpt) {
	return func(opt *MockRedisOpt) {
		opt.Size = size
	}
}

func WithDockerRedisDb(db int) func(opt *MockRedisOpt) {
	return func(opt *MockRedisOpt) {
		opt.Db = db
	}
}

func MockRdb() (*Rdb, func()) {

	opt, cleanup := mockRedisOption()

	rdb, _ := NewRdb(opt)

	return rdb, cleanup
}

func mockRedisOption() (RedisOption, func()) {
	s, err := miniredis.Run()
	utee.Chk(err)

	opt := RedisOption{
		Address:  s.Addr(),
		Password: "",
		Db:       0,
		Pool:     20,
		Name:     "mic-project-layout-test",
	}

	return opt, func() {
		s.Close()
	}
}

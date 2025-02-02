package database

import (
	"fmt"
	"strings"
	"time"

	"github.com/qmute/mic-project-layout/pkg/ut"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"

	"github.com/pkg/errors"
	"github.com/qmute/dbc/gdb"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Option 数据库连接选项
type Option struct {
	User     string `validate:"required"`
	Password string `validate:"required"`
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	Database string `validate:"required"`
}

func (p Option) Conn() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable connect_timeout=5",
		p.Host, p.Port, p.User, p.Database, p.Password)
}

// NewGdb 返回一个gorm db
func NewGdb(opt Option) (*gorm.DB, error) {
	if err := ut.ValidStruct(opt); err != nil {
		return nil, err
	}

	gOpts := getGromOptions()
	gormDB, err := gdb.ConnectToPG(opt.Conn(), GetGormConfig(), gOpts...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return gormDB, nil

}

func getGromOptions() []gdb.Option {
	opts := []gdb.Option{
		gdb.WithConnMaxLifetime(3 * time.Minute),
		gdb.WithConnMaxIdleTime(3 * time.Minute),
		gdb.WithMaxIdleConns(10),
		gdb.WithMaxOpenConns(100),
	}
	return opts
}

func GetGormConfig() *gorm.Config {
	cfg := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.New(logrus.StandardLogger(), logger.Config{
			SlowThreshold:             2 * time.Second,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		}),

		CreateBatchSize: 100,
	}

	return cfg
}

// NotFound 数据没有找到
func NotFound(err error) bool {
	if err == nil {
		return false
	}
	return gdb.NotFound(err) || strings.Contains(err.Error(), gorm.ErrRecordNotFound.Error())
}

// Dup 数据重复
func Dup(err error) bool {
	return gdb.Dup(err)
}

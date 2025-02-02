package config

import (
	"encoding/json"
	"os"

	"easyslip.cc/mic-project-layout/internal/database"
	"easyslip.cc/mic-project-layout/pkg/ut"
	"github.com/pkg/errors"
	"github.com/tidwall/jsonc"
)

type Path string

func (p Path) Val() string {
	return string(p)
}

func New(path Path) (Config, error) {
	if path == "" {
		panic("config path is empty")
	}

	var cfg Config

	data, err := os.ReadFile(path.Val())
	if err != nil {
		return cfg, errors.WithStack(err)
	}

	err = json.Unmarshal(jsonc.ToJSON(data), &cfg)
	if err != nil {
		return cfg, errors.WithStack(err)
	}

	if err = cfg.Valid(); err != nil {
		return cfg, err
	}

	return cfg, nil
}

type Config struct {
	Pg    Pg    `json:"pg" validate:"required"`
	Redis Redis `json:"redis" validate:"required"`
}

func (p *Config) Valid() error {
	return ut.ValidStruct(p)
}

type Pg struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	Host     string `json:"host" validate:"required"`
	Port     string `json:"port" validate:"required"`
	Database string `json:"database" validate:"required"`
}

type Redis struct {
	Addr     string `json:"addr" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Db       int    `json:"db" validate:"gte=0"`
	Pool     int    `json:"pool" validate:"gte=0"`
	Name     string `json:"name" validate:"required"` // 保存数据库名称，便于使用共享redis实例时识别用途
}

func (p Redis) ToOpt() database.RedisOption {
	return database.RedisOption{
		Address:  p.Addr,
		Username: p.Username,
		Password: p.Password,
		Db:       p.Db,
		Pool:     p.Pool,
		Name:     p.Name,
	}
}

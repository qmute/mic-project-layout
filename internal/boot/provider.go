package boot

import (
	"easyslip.cc/mic-project-layout/internal/config"
	"easyslip.cc/mic-project-layout/internal/database"
	micConifg "go-micro.dev/v4/config"
)

func GetConfigPath(micCfg micConifg.Config) config.Path {
	return config.Path(micCfg.Get("config").String("./resources/dev.json5"))
}

func DatabaseOpt(cfg config.Config) database.Option {
	return database.Option{
		User:     cfg.Pg.User,
		Password: cfg.Pg.Password,
		Host:     cfg.Pg.Host,
		Port:     cfg.Pg.Port,
		Database: cfg.Pg.Database,
	}
}

func RedisOpt(cfg config.Config) database.RedisOption {
	return cfg.Redis.ToOpt()
}

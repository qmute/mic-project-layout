package boot

import (
	"easyslip.cc/mic-project-layout/internal/config"
	"easyslip.cc/mic-project-layout/internal/database"
	repoImpl "easyslip.cc/mic-project-layout/internal/repo/impl"
	srvImpl "easyslip.cc/mic-project-layout/internal/service/impl"
	"github.com/google/wire"
	"github.com/qmute/mic/v4"
)

var BaseSet = wire.NewSet(
	mic.DefaultConfig,
	mic.DefaultService,
	mic.DefaultWeb,
	mic.NewSync,

	GetConfigPath,
	config.New,

	wire.Struct(new(Bootloader), "*"),

	dbSet,

	repoImpl.Set,
	srvImpl.Set,
)

var dbSet = wire.NewSet(
	DatabaseOpt,
	database.NewGdb,

	RedisOpt,
	database.NewRdb,
)

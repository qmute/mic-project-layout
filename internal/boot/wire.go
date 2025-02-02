package boot

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/config"
	"github.com/qmute/mic-project-layout/internal/database"
	repoImpl "github.com/qmute/mic-project-layout/internal/repo/impl"
	srvImpl "github.com/qmute/mic-project-layout/internal/service/impl"
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

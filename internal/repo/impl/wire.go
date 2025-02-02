package impl

import (
	"easyslip.cc/mic-project-layout/internal/repo"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(repo.DbInitializer), "*"),
)

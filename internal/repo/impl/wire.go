package impl

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/repo"
)

var Set = wire.NewSet(
	wire.Struct(new(repo.DbInitializer), "*"),
)

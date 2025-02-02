package hdl

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/app/serve/admin/hdl/mid"
)

var Set = wire.NewSet(
	wire.Struct(new(Hdl), "*"),
	wire.Struct(new(Base), "*"),
	wire.Struct(new(mid.Mid), "*"),

	wire.Struct(new(PubHdl), "*"),
)

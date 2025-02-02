package hdl

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/app/serve/front/hdl/mid"
)

var Set = wire.NewSet(
	wire.Struct(new(Hdl), "*"),
	wire.Struct(new(Base), "*"),
	wire.Struct(new(mid.Mid), "*"),
	wire.Struct(new(mid.ClientMid), "*"),

	wire.Struct(new(PubHdl), "*"),
)

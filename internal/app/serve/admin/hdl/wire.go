package hdl

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/admin/hdl/mid"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(Hdl), "*"),
	wire.Struct(new(Base), "*"),
	wire.Struct(new(mid.Mid), "*"),

	wire.Struct(new(PubHdl), "*"),
)

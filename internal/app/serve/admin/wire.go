package admin

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/admin/hdl"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(App), "*"),

	hdl.Set,
)

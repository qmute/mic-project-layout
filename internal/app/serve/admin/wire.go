package admin

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/app/serve/admin/hdl"
)

var Set = wire.NewSet(
	wire.Struct(new(App), "*"),

	hdl.Set,
)

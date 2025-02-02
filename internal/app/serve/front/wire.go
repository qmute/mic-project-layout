package front

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/front/hdl"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Struct(new(App), "*"),

	hdl.Set,
)

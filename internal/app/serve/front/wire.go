package front

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/app/serve/front/hdl"
)

var Set = wire.NewSet(
	wire.Struct(new(App), "*"),

	hdl.Set,
)

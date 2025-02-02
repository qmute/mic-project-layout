//go:build wireinject
// +build wireinject

package serve

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/app/serve/admin"
	"github.com/qmute/mic-project-layout/internal/app/serve/front"
	"github.com/qmute/mic-project-layout/internal/boot"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		wire.Struct(new(App), "*"),
		wire.Struct(new(Bootloader), "*"),

		boot.BaseSet,

		defaultMicOpt,
		defaultMicWebOpt,

		wire.Struct(new(Web), "*"),
		front.Set,
		admin.Set,
	))
}

//go:build wireinject
// +build wireinject

package serve

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/admin"
	"easyslip.cc/mic-project-layout/internal/app/serve/front"
	"easyslip.cc/mic-project-layout/internal/boot"
	"github.com/google/wire"
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

//go:build wireinject
// +build wireinject

package daemon

import (
	"easyslip.cc/mic-project-layout/internal/app/daemon/internal/daemon"
	"easyslip.cc/mic-project-layout/internal/boot"
	"github.com/google/wire"
)

func New() (*App, func(), error) {
	panic(wire.Build(
		wire.Struct(new(App), "*"),
		wire.Struct(new(Bootloader), "*"),

		boot.BaseSet,

		daemon.Set,

		defaultMicDaemonOpt,
		defaultMicWebDaemonOpt,
	))
}

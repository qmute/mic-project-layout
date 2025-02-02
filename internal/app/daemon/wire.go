//go:build wireinject
// +build wireinject

package daemon

import (
	"github.com/google/wire"
	"github.com/qmute/mic-project-layout/internal/app/daemon/internal/daemon"
	"github.com/qmute/mic-project-layout/internal/boot"
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

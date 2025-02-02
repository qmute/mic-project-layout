package daemon

import (
	"context"

	"github.com/qmute/mic-project-layout/internal/app/daemon/internal/daemon"
	"github.com/qmute/mic-project-layout/internal/boot"
	"github.com/qmute/mic/v4"
)

// Bootloader 本应用的启动引导。 每个应用都把自己关心的初始化器加到这里
type Bootloader struct {
	boot.Bootloader
	Daemon *daemon.Initializer
}

func (p *Bootloader) getInitializer() []mic.Initializer {
	return []mic.Initializer{
		p.Daemon,
	}
}

func (p *Bootloader) Boot(ctx context.Context) error {
	if err := p.Bootloader.Boot(ctx); err != nil {
		return err
	}

	return mic.InitAll(ctx, p.getInitializer()...)
}

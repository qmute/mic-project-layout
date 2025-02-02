package daemon

import (
	"context"

	"github.com/qmute/mic/v4"
)

type Initializer struct {
}

func (p Initializer) Name() string {
	return "daemon_initializer"
}

func (p Initializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

func (p Initializer) Initialize(ctx context.Context) error {
	return mic.InitDaemon(p.getDaemon()...)
}

func (p Initializer) getDaemon() []mic.Daemon {
	return []mic.Daemon{}
}

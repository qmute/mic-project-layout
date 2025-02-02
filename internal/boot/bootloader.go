package boot

import (
	"context"

	"github.com/qmute/mic/v4"
)

// Bootloader 最基本启动引导
type Bootloader struct {
	// DbInitializer *repo.DbInitializer
}

func (p *Bootloader) Boot(ctx context.Context, it ...mic.Initializer) error {
	l := p.getInitializer(it...)
	return mic.InitAll(ctx, l...)
}

func (p *Bootloader) getInitializer(it ...mic.Initializer) []mic.Initializer {
	l := []mic.Initializer{
		// p.DbInitializer,
	}

	l = append(l, it...)
	return l
}

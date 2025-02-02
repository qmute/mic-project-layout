package hdl

import (
	"github.com/qmute/gi"
	"github.com/qmute/mic-project-layout/internal/app/serve/admin/hdl/mid"
)

type Base struct {
	gi.BaseHdl `wire:"-"`
	Mid        *mid.Mid
}

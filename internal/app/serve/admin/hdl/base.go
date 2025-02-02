package hdl

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/admin/hdl/mid"
	"github.com/qmute/gi"
)

type Base struct {
	gi.BaseHdl `wire:"-"`
	Mid        *mid.Mid
}

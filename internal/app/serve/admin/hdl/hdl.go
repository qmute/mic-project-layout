package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/qmute/mic-project-layout/internal/app/serve/admin/hdl/mid"
)

type Hdl struct {
	Mid *mid.Mid
	Pub *PubHdl
}

func (p *Hdl) Mount(rg gin.IRouter) {
	pubGroup := rg.Group("/pub")
	p.Pub.Mount(pubGroup)
}

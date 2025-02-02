package hdl

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/admin/hdl/mid"
	"github.com/gin-gonic/gin"
)

type Hdl struct {
	Mid *mid.Mid
	Pub *PubHdl
}

func (p *Hdl) Mount(rg gin.IRouter) {
	pubGroup := rg.Group("/pub")
	p.Pub.Mount(pubGroup)
}

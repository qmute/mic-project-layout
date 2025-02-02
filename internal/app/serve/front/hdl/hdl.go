package hdl

import (
	"github.com/gin-gonic/gin"
	"github.com/qmute/mic-project-layout/internal/app/serve/front/hdl/mid"
)

type Hdl struct {
	Mid    *mid.Mid
	PubHdl *PubHdl
}

func (p *Hdl) Mount(rg gin.IRouter) {

	clientTypeG := rg.Group("", p.Mid.ClientMid.ExistClientInfo())
	{
		p.PubHdl.Mount(clientTypeG.Group("/pub"))
	}

}

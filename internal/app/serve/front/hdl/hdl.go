package hdl

import (
	"easyslip.cc/mic-project-layout/internal/app/serve/front/hdl/mid"
	"github.com/gin-gonic/gin"
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

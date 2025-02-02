package hdl

import (
	"github.com/gin-gonic/gin"
)

type PubHdl struct {
	Base
}

func (p *PubHdl) Mount(rg gin.IRouter) {
	rg.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello front",
		})
	})
}

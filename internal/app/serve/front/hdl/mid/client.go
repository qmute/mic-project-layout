package mid

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/qmute/mic-project-layout/internal/domain"
	log "github.com/sirupsen/logrus"
)

const (
	XClientType    = "X-Client-Type"
	XClientVersion = "X-Client-Version"
	XUserType      = "X-User-Type"
	ClientInfoKey  = "clientInfoKey"
)

type ClientMid struct {
}

func (p ClientMid) ExistClientInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		cltTp := domain.ClientType(strings.TrimSpace(c.Request.Header.Get(XClientType)))
		cltVer := strings.TrimSpace(c.Request.Header.Get(XClientVersion))
		info := &domain.ClientInfo{
			ClientType:    cltTp,
			ClientVersion: cltVer,
			IP:            c.ClientIP(),
		}
		if err := info.Valid(); err != nil {
			log.WithError(err).Errorln("valid clientInfo error")
			c.String(http.StatusForbidden, "客户端信息错误")
			c.Abort()
			return
		}

		c.Set(ClientInfoKey, info)

	}
}

func (p ClientMid) MustGetClientInfo(c *gin.Context) *domain.ClientInfo {
	return c.MustGet(ClientInfoKey).(*domain.ClientInfo)
}

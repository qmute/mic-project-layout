package front

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qmute/gi"
	"github.com/qmute/mic-project-layout/internal/app/serve/front/hdl"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	sessionName = "web-session-front"
	sessionSalt = "!FrOn)t.XDyki&*DP]Lh6KqY9s8"
)

// App 消费者端API
// @title  mic-project-layout用户端API
// @version 1.0
// @description mic-project-layout用户端API，必须有的Header: X-Client-Type: ios,android,harmony
// @host http://mic-project-layout.com
// @BasePath /api/front
// @schemes https
type App struct {
	Hdl *hdl.Hdl
}

func (p *App) Run(rg gin.IRouter) {
	rg.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), func(c *ginSwagger.Config) {
		c.InstanceName = "front"
	}))

	rg.Use(gi.MidCookieSession(sessionName, sessionSalt, p.sessionOpt()))
	p.Hdl.Mount(rg)
}

func (p *App) sessionOpt() sessions.Options {
	return sessions.Options{
		Path:     "/api/front",
		MaxAge:   86400 * 360 * 50,
		Secure:   true,
		HttpOnly: true,
	}
}

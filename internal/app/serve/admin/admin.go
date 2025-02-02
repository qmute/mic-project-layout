package admin

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/qmute/gi"
	"github.com/qmute/mic-project-layout/internal/app/serve/admin/hdl"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	sessionName = "web-session-admin"
	sessionSalt = "!adMiN.MZeIhEN*DP]LhKqY6s8"
)

// App 管理后台api
// @title  mic-project-layout管理后台API
// @version 1.0
// @description mic-project-layout管理后台API
// @host https://mic-project-layout.com
// @BasePath /api/admin
// @schemes https
type App struct {
	Hdl *hdl.Hdl
}

func (p *App) Run(rg gin.IRouter) {
	rg.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), func(c *ginSwagger.Config) {
		c.InstanceName = "admin"
	}))

	rg.Use(gi.MidCookieSession(sessionName, sessionSalt, p.sessionOpt()))
	p.Hdl.Mount(rg)
}

func (p *App) sessionOpt() sessions.Options {
	return sessions.Options{
		Path:     "/api/admin",
		MaxAge:   86400 * 360 * 50,
		Secure:   true,
		HttpOnly: true,
	}
}

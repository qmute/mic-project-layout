package serve

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/qmute/mic/v4"
	"github.com/quexer/utee"
	log "github.com/sirupsen/logrus"
	"go-micro.dev/v4"
	"go-micro.dev/v4/web"
)

type App struct {
	Bootloader      *Bootloader
	MicroWebService web.Service
	Service         micro.Service
	Web             *Web
}

func (p *App) Run() error {
	return p.MicroWebService.Run()
}

func (p *App) Init() *App {
	mic.InitLogger()

	err := p.MicroWebService.Init(
		web.AfterStart(p.startLog),
		web.BeforeStart(p.beforeStart),
	)
	utee.Chk(err)

	p.MicroWebService.Handle("/", p.Web.Mount(p.getVersion()))

	return p
}

func (p *App) Name() string {
	return "mic-project-layout-serve"
}

func (p *App) getVersion() string {
	return p.MicroWebService.Options().Version
}

func (p *App) startLog() error {
	log.WithField("app", p.Name()).
		WithField("version", p.getVersion()).
		WithField("ginMode", gin.Mode()).
		Infoln("startup")
	return nil
}

func (p *App) beforeStart() error {
	return p.Bootloader.Boot(context.Background())
}

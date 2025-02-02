package daemon

import (
	"context"
	"net/http"

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
}

func (p *App) Init() *App {
	mic.InitLogger()
	p.Service.Init(
		micro.BeforeStart(p.beforeStart),
		micro.AfterStart(p.startLog),
	)
	return p
}

func (p *App) Run() error {
	go func() {
		utee.Chk(p.MicroWebService.Run())
	}()
	return p.Service.Run()
}

func (p *App) name() string {
	return "daemon"
}

func (p *App) getVersion() string {
	return p.MicroWebService.Options().Version
}

func (p *App) startLog() error {
	log.WithField("app", p.name()).
		WithField("version", p.getVersion()).
		WithField("ginMode", gin.Mode()).
		Infoln("startup", p.name(), p.getVersion(), gin.Mode())
	return nil
}

func (p *App) beforeStart() error {
	p.MicroWebService.Handle("/debug/", http.DefaultServeMux)
	p.MicroWebService.HandleFunc("/v", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(p.MicroWebService.Options().Version))
	})

	return p.Bootloader.Boot(context.Background())
}

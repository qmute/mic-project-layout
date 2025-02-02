package daemon

import (
	"time"

	"github.com/qmute/mic/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/config"
)

func defaultMicDaemonOpt(conf config.Config) mic.Opt {
	return mic.Opt{
		Name:           "mic-project-layout.srv.daemon",
		Version:        mic.BuildVersion("1.0.0"),
		HystrixTimeout: 15 * time.Second,
	}
}

func defaultMicWebDaemonOpt(conf config.Config, service micro.Service) mic.WebOpt {
	return mic.WebOpt{
		Name:    "mic-project-layout.web.daemon",
		Addr:    conf.Get("addr").String(":3002"),
		Service: service,
	}
}

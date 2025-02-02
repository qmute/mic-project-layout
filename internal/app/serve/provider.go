package serve

import (
	"time"

	"github.com/qmute/mic/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/config"
)

func defaultMicOpt(conf config.Config) mic.Opt {
	return mic.Opt{
		Version:        mic.BuildVersion("1.0.0"),
		Name:           "mic-project-layout.srv.serve",
		HystrixTimeout: 15 * time.Second,
	}
}

func defaultMicWebOpt(conf config.Config, service micro.Service) mic.WebOpt {
	return mic.WebOpt{
		Name:    "mic-project-layout.web.serve",
		Addr:    conf.Get("addr").String(":3000"),
		Service: service,
	}
}

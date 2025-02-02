package main

import (
	_ "expvar"
	_ "net/http/pprof"

	"easyslip.cc/mic-project-layout/cmd"
	"github.com/qmute/mic/v4"
)

func main() {
	mic.InitLogger()
	cmd.Execute()
}

package main

import (
	_ "expvar"
	_ "net/http/pprof"

	"github.com/qmute/mic-project-layout/cmd"
	"github.com/qmute/mic/v4"
)

func main() {
	mic.InitLogger()
	cmd.Execute()
}
